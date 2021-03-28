package sansa

// Event TODO
type Event struct {
    ID        AggregateID `json:"id"`
    Aggregate string      `json:"aggregate"`
    Type      string      `json:"type"`
    Sequence  uint64      `json:"sequence"`
    Data      Event       `json:"data"`
}
