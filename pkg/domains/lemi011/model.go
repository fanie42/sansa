package lemi011

import (
    "fmt"
    "time"
)

type data struct {
    id          string    `json:"id" db:"id"`
    timestamp   time.Time `json:"timestamp" db:"id"`
    x           int64     `json:"x" db:"x"`
    y           int64     `json:"y" db:"y"`
    z           int64     `json:"z" db:"z"`
    temperature int64     `json:"temperature" db:"temperature"`
}

// ID TODO
func (d *data) ID() string {
    return d.id
}

// Timestamp TODO
func (d *data) Timestamp() time.Time {
    return d.timestamp
}

// String TODO - This is a presenter
func (d *data) String() string {
    return fmt.Sprintf(
        "%s, %d, %d, %d, %d",
        d.timestamp.Format("2006-01-02 15:04:05.000000"),
        d.x,
        d.y,
        d.z,
        d.temperature,
    )
}
