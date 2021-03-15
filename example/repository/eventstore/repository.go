package eventstore

import (
    es "github.com/fanie42/sansa/eventsourcing"
    "github.com/fanie42/sansa/example"
)

type repository struct {
    db es.EventStore
}

// New TODO
func New(
    database es.EventStore,
) example.Service {
    return &repository{
        db: database,
    }
}

// Load TODO
func (repo *repository) Load(
    id example.DeviceID,
) (*example.Device, error) {
    events, err := repo.db.GetEvents(id)
    if err != nil {
        return err
    }

    device := &Device{}

    for _, event := range events {
        device.Apply(event)
    }

    return device, nil
}

// Save TODO
func (repo *repository) Save(
    device *example.Device,
) error {
    events, err := repo.db.GetEvents(id)
    if err != nil {
        return err
    }

    device := &Device{}

    for _, event := range events {
        device.Apply(event)
    }

    return device, nil
}
