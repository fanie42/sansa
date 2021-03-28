package sansa

// AggregateID TODO
type AggregateID interface{}

// AggregateType TODO
type AggregateType string

// Aggregate TODO
type Aggregate struct {
    ID      AggregateID   `json:"id"`
    Type    AggregateType `json:"type"`
    Changes []Event       `json:"changes"`
}
