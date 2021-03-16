package inmem

import (
    "time"

    "github.com/fanie42/sansa/dvras"
)

// Data TODO
type Data struct {
    sensor    dvras.SensorID
    timestamp time.Time
    x         int
    y         int
    z         int
    t         int
}
