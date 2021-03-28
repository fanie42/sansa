package eventstore

import (
    "context"
    "encoding/json"
    "reflect"

    "github.com/fanie42/sansa/dvras"
)

type repository struct{}

func (r *repository) Load(
    ctx context.Context,
    id dvras.DeviceID,
) (*dvras.Device, error) {
    // load up all events
    records, err := r.store.Load(ctx, id.String(), 0, 0)
    if err != nil {
        return nil, err
    }

    events := []patient.Event{}
    linq.From(records).
        SelectT(func(record eventsource.Record) patient.Event {
            if err != nil {
                return nil
            }

            var typed struct {
                Type string `json:"event_type"`
            }
            err = json.Unmarshal(record.Data, &typed)
            if err != nil {
                return nil
            }

            var e patient.Event
            switch typed.Type {
            case eventName(&patient.PatientAdmitted{}):
                e = &patient.PatientAdmitted{}
            case eventName(&patient.PatientTransferred{}):
                e = &patient.PatientTransferred{}
            case eventName(&patient.PatientDischarged{}):
                e = &patient.PatientDischarged{}
            }

            err = json.Unmarshal(record.Data, e)
            if err != nil {
                return nil
            }

            return e
        }).
        ToSlice(&events)
    if err != nil {
        return nil, err
    }

    return patient.NewFromEvents(events), nil
}

func (r *repository) Save(
    ctx context.Context,
    p *dvras.Device,
) error {
    records := make([]eventsource.Record, len(p.Events()))

    var err error
    linq.From(p.Events()).
        SelectT(func(event patient.Event) eventsource.Record {
            var data []byte
            switch e := event.(type) {
            case *patient.PatientAdmitted:
                data, err = json.Marshal(admitted{
                    Type:            eventName(e),
                    PatientAdmitted: e,
                })
            case *patient.PatientDischarged:
                data, err = json.Marshal(discharged{
                    Type:              eventName(e),
                    PatientDischarged: e,
                })
            case *patient.PatientTransferred:
                data, err = json.Marshal(transferred{
                    Type:               eventName(e),
                    PatientTransferred: e,
                })
            }
            return eventsource.Record{
                Data: data,
            }
        }).
        ToSlice(&records)
    if err != nil {
        return err
    }

    for i := range records {
        expectedVersion := i + p.Version()
        records[i].Version = expectedVersion
    }

    return r.store.Save(ctx, p.ID().String(), records...)
}

type admitted struct {
    Type string `json:"event_type"`
    *patient.PatientAdmited
}

type transferred struct {
    Type string `json:"event_type"`
    *patient.PatientTransferred
}

type discharged struct {
    Type string `json:"event_type"`
    *patient.PatientDischarged
}

func eventName(event patient.Event) string {
    t := reflect.TypeOf(event)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }

    return t.Name()
}
