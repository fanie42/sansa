package sansa

import "fmt"

// AggregateOutdatedError TODO
type AggregateOutdatedError struct {
    Latest    uint64
    Attempted uint64
}

// Error TODO
func (err AggregateOutdatedError) Error() string {
    return fmt.Sprintf(
        "attempted to save aggregate on sequence: %d, "+
            "but latest sequence is at: %d",
        err.Attempted,
        err.Latest,
    )
}
