package nats

import (
    "context"
    "fmt"
    "os"

    "github.com/fanie42/sansa/pkg/controllers"
    "github.com/fanie42/sansa/pkg/domains"
    "github.com/nats-io/nats.go"
    "github.com/nats-io/stan.go"
)

type controller struct {
    config  Config
    service domains.Service
    sc      stan.Conn
}

// NewController TODO
func NewController(
    cfg Config,
    svc domains.Service,
) controllers.Controller {
    natsConn, err := nats.Connect(
        cfg.URL,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    stanConn, err := stan.Connect(
        cfg.ClusterID,
        cfg.ClientID,
        stan.NatsConn(natsConn),
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
        return nil
    }
    ctrl := &controller{
        config:  cfg,
        service: svc,
        sc:      stanConn,
    }

    return ctrl
}

func (ctrl *controller) Listen(
    ctx context.Context,
) error {
    _, err := ctrl.sc.Subscribe(
        "lemi011",
        func(msg *stan.Msg) {
            m, err := ctrl.service.Build(string(msg.Data))
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println(m.String())
            ctrl.service.Create(m)
        },
        // stan.DurableName("lemi0111"),
    )
    return err
}

func (ctrl *controller) Serve(
    ctx context.Context,
) error {
    ch := ctrl.service.Stream()
    for {
        select {
        case m := <-ch:
            fmt.Println(m.String())
            ctrl.sc.Publish(
                "lemi011",
                []byte(m.String()),
            )
        case <-ctx.Done():
            return ctx.Err()
        }
    }
    return nil
}

func (ctrl *controller) ListenAndServe(
    ctx context.Context,
) error {
    return nil
}
