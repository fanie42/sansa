package serial

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/fanie42/sansa/pkg/lemi011b"
    "github.com/tarm/serial"
)

// So technically, this is a handler, because it only handles the
// specific case for lemi011b. A controller would be something that
// controlled the device and did not care which instrument it was
// controlling (more disconnected infrastructure vibe.)

// Controller TODO
// This should USE the presenter as well as the request, response and
// service. also define a presenter interface that it can use.
type Controller struct {
    config     Config
    service    lemi011b.Service
    serialPort *serial.Port
}

// New TODO
func New(
    config Config,
    service lemi011b.Service,
) *Controller {
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

    return &Controller{
        config:     config,
        service:    service,
        serialPort: sp,
    }
}

// Run TODO
func (ctrl *Controller) Run() {
    scanner := bufio.NewScanner(ctrl.serialPort)
    for scanner.Scan() {
        raw := &raw{
            timestamp: time.Now().UTC(),
            bytes:     scanner.Bytes(),
        }
        d, err := raw.toLemi011bData()

        if err != nil {
            fmt.Printf(
                "could not parse bytes to lemi011b data with error: %v",
                err,
            )
            continue
        }

        ctrl.service.AddNewData(d)
    }
}

type raw struct {
    timestamp time.Time
    bytes     []byte
}

func (r *raw) toLemi011bData() (*lemi011b.Request, error) {
    s := strings.Split(string(r.bytes), ", ")
    l := len(s)
    if l != 4 {
        return nil, fmt.Errorf("expected 4 data fields, but got: %d", l)
    }

    x, err := strconv.ParseInt(s[0], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse x value: %v to int with error: %v",
            s[0],
            err,
        )
    }

    y, err := strconv.ParseInt(s[1], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse y value: %v to int with error: %v",
            s[1],
            err,
        )
    }

    z, err := strconv.ParseInt(s[2], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse z value: %v to int with error: %v",
            s[2],
            err,
        )
    }

    temperature, err := strconv.ParseInt(s[3], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse temperature value: %v to int with error: %v",
            s[3],
            err,
        )
    }

    return &lemi011b.Request{
        Timestamp:   r.timestamp,
        X:           int32(x),
        Y:           int32(y),
        Z:           int32(z),
        Temperature: int32(temperature),
    }, nil
}
