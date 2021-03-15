package eventsourcing

import "fmt"

// AggregateRoot TODO
type AggregateRoot struct {
    ID      AggregateID
    Events  []*Event
    Version int
}

// NewAggregateRoot TODO
func NewAggregateRoot() (*AggregateRoot, error) {
    aggID, err := GenerateID()
    if err != nil {
        return nil, err
    }
    return &AggregateRoot{
        AggregateID: aggID,
    }, nil
}

// Store TODO
func (ar *AggregateRoot) Store() {
    fmt.Println("hello")
}

// Apply TODO
func (ar *AggregateRoot) Apply(e *Event) {
    ar.Events = append(ar.Events, e)
}

//GenerateID generates a unique ID using UUID v4.
func GenerateID() (string, error) {
    u, err := uuid.GenerateUuidV4()
    if err != nil {
        return "", err
    }
    return u, nil
}
