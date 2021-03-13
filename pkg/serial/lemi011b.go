package serial

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/fanie42/sansa/pkg/lemi011b"
)

// Lemi011b TODO
func (h *Handler) Lemi011b(svc lemi011b.Service) error {
    scanner := bufio.NewScanner(ctrl.serialPort)
    for scanner.Scan() {
        t := time.Now().UTC()
        b := scanner.Bytes()
        d, err := toLemi011bData(t, b)

        if err != nil {
            fmt.Printf(
                "could not parse bytes to lemi011b data with error: %v",
                err,
            )
            continue
        }

        svc.Create(d)
    }

    return scanner.Err()
}

func toLemi011bRequest(t time.Time, b []byte) (*lemi011b.Request, error) {
    s := strings.Split(string(b), ", ")
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
