package inmem

import "github.com/fanie42/sansa/dvras"

// Sensor TODO
type Sensor struct {
    id     dvras.SensorID
    device admin.DeviceID
    state  dvras.State
}
