package nats

import (
    "github.com/fanie42/sansa"
    natsio "github.com/nats-io/nats.go"
)

type dispatcher struct {
    nc *natsio.Conn
}

// NewDispatcher TODO
func NewDispatcher(
    conn *natsio.Conn,
) sansa.Dispatcher {
    return &dispatcher{
        nc: conn,
    }
}

// Handle TODO
func (dispatch *dispatcher) Dispatch(
    record Record,
) error {
    data, err := record.Event.MarshalBinary()
    if err != nil {
        return err
    }

    err := dispatch.nc.Publish(
        record.AggregateType+"."+record.AggregateID+"."+record.EventType,
        data,
    )

    if err != nil {
        return err
    }
}
