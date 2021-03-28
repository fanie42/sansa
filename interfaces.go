package sansa

// EventStore TODO
type EventStore interface {
    Save(Event) error
    Load(ID) (Event, error)
}

// Projector TODO
type Projector interface {
    Project(AggregateType)
}

// Dispatcher TODO
type Dispatcher interface {
    Publish(Event)
}
