package nats

import (
    "github.com/fanie42/sansa"
    natsio "github.com/nats-io/nats.go"
)

type projector struct {
    nc *natsio.Conn
    h  EventHandler
}

// NewProjector TODO
func NewProjector(
    connection *natsio.Conn,
    handler EventHandler,
) sansa.Projector {
    return &projector{
        nc: connection,
        h:  handler,
    }
}

// Handle TODO
func (proj *projector) Project(
    aggregateType string, // devices
    eventType string, // created
) error {
    sub, err := proj.nc.Subscribe(
        aggregateType+".*."+eventType,
        callback,
    )

    if err != nil {
        return err
    }

    return nil
}

func (proj *projector) callback(msg *natsio.Msg) {
    err := event.UnmarshalBinary(msg.Data)
    proj.h.Handle(event)
}
