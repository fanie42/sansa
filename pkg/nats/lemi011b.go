package nats

import (
    "context"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/fanie42/sansa/pkg/lemi011b"
    "github.com/nats-io/nats.go"
    "github.com/nats-io/stan.go"
)

// Implements the presenter that the serial.Controller uses!!! In other words
// every time there is a new datapoint, the controller will call the service
// with the request, receive the response and relay it to the presenter.

// Presenter TODO - should implement the lemi011b Presenter. Exactly the same
// as
type Presenter struct {
    config Config
    sc     stan.Conn
}

// NewPresenter TODO
func NewPresenter(
    config Config,
) *Presenter {
    natsConn, err := nats.Connect(
        config.URL,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    stanConn, err := stan.Connect(
        config.ClusterID,
        config.ClientID,
        stan.NatsConn(natsConn),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    return &Presenter{
        config: config,
        sc:     stanConn,
    }
}

// Send TODO
func (pres *Presenter) Send(r *lemi011b.Response) error {
    s := fmt.Sprintf(
        "%s, %d, %d, %d, %d",
        r.Timestamp.Format("2006-01-02 15:04:05.000000"),
        r.X,
        r.Y,
        r.Z,
        r.Temperature,
    )
    pres.sc.Publish("marl111", []byte(s))

    return nil
}

// Controller TODO - should implement the lemi011b Presenter. Exactly the same
// as
type Controller struct {
    config  Config
    service lemi011b.Service
    sc      stan.Conn
}

// NewController TODO
func NewController(
    config Config,
    service lemi011b.Service,
) *Controller {
    natsConn, err := nats.Connect(
        config.URL,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    stanConn, err := stan.Connect(
        config.ClusterID,
        config.ClientID,
        stan.NatsConn(natsConn),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    return &Controller{
        config:  config,
        service: service,
        sc:      stanConn,
    }
}

// Run TODO
func (ctrl *Controller) Run(ctx context.Context) error {
    sub, err := ctrl.sc.Subscribe(
        "marl111",
        func(m *stan.Msg) {
            d, err := toLemi011bData(m.Data)
            if err != nil {
                fmt.Println(err)
            }
            ctrl.service.AddNewData(d)
        },
    )
    defer sub.Close()

    if err != nil {
        return err
    }

    <-ctx.Done()

    return ctx.Err()
}

func toLemi011bData(b []byte) (*lemi011b.Request, error) {
    s := strings.Split(string(b), ", ")
    l := len(s)
    if l != 5 {
        return nil, fmt.Errorf("expected 5 data fields, but got: %d", l)
    }

    timestamp, err := time.Parse(
        "2006-01-02 15:04:05.000000",
        string(s[0]),
    )
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse timestamp value: %v to time.Time with error: %v",
            s[0],
            err,
        )
    }

    x, err := strconv.ParseInt(s[1], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse x value: %v to int with error: %v",
            s[1],
            err,
        )
    }

    y, err := strconv.ParseInt(s[2], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse y value: %v to int with error: %v",
            s[2],
            err,
        )
    }

    z, err := strconv.ParseInt(s[3], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse z value: %v to int with error: %v",
            s[3],
            err,
        )
    }

    temperature, err := strconv.ParseInt(s[4], 10, 32)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse temperature value: %v to int with error: %v",
            s[4],
            err,
        )
    }

    return &lemi011b.Request{
        Timestamp:   timestamp,
        X:           int32(x),
        Y:           int32(y),
        Z:           int32(z),
        Temperature: int32(temperature),
    }, nil
}
