package serial

import (
    "fmt"
    "time"

    "github.com/tarm/serial"
)

// So technically, this is a handler, because it only handles the
// specific case for lemi011b. A controller would be something that
// controlled the device and did not care which instrument it was
// controlling (more disconnected infrastructure vibe.)

// Handler TODO
// This should USE the presenter as well as the request, response and
// service. also define a presenter interface that it can use.
type Handler struct {
    config     Config
    serialPort *serial.Port
}

// New TODO
func New(
    config Config,
) *Handler {
    sp, err := serial.OpenPort(
        &serial.Config{
            Name: config.Name,
            Baud: config.Baud,
        },
    )
    if err != nil {
        fmt.Println("could not open serial port")
    }

    time.Sleep(time.Second)
    if err := sp.Flush(); err != nil {
        fmt.Println("could not flush serial port")
    }

    return &Handler{
        config:     config,
        serialPort: sp,
    }
}
