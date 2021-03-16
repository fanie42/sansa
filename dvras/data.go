package dvras

import (
    "time"
)

// Data TODO - Aggregate Root
type Data struct {
    id        DataID
    deviceID  device.ID
    timestamp time.Time
    ch1       []int16
    ch2       []int16
    pps       []int16

    events  []Event
    version int
}

// New TODO
func New() (*Data, error) {
    return &Data{
        id: id,
    }
}

// Acquire TODO
func (d *Data) Acquire(
    deviceID device.ID,
    timestamp time.Time,
    ch1 []int16,
    ch2 []int16,
    pps []int16,
) error {
    d.raise(&AcquiredEvent{
        AggregateID: d.id,
        DeviceID:    deviceID,
        Timestamp:   timestamp,
        Ch1:         ch1,
        Ch2:         ch2,
        PPS:         pps,
    })

    return nil
}

func (d *Device) raise(event Event) {
    d.events = append(d.events, event)
    event.Apply(d)
}
