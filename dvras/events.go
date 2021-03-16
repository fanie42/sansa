package dvras

import "time"

// Event TODO
type Event interface {
    Apply(*Device)
}

// AcquiredEvent TODO
type AcquiredEvent struct {
    AggregateID DataFrameID `json:"id"`
    DeviceID    device.ID   `json:"device_id"`
    Timestamp   time.Time   `json:"timestamp"`
    Ch1         []int16     `json:"ch1"`
    Ch2         []int16     `json:"ch2"`
    PPS         []int16     `json:"pps"`
}

// Apply TODO
func (event *AcquiredEvent) Apply(
    data *Data,
) {
    data.id = event.AggregateID
    data.timestamp = event.Timestamp
    data.ch1 = event.Ch1
    data.ch2 = event.Ch2
    data.pps = event.PPS
}
