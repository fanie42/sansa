package eventsourced

import (
    "time"

    "github.com/fanie42/sansa/dvras"
)

// AcquireDataCommand TODO
type AcquireDataCommand struct {
    ID        dvras.DataID
    DeviceID  admin.DeviceID
    Timestamp time.Time
    Ch1       []int16
    Ch2       []int16
    PPS       []int16
}
