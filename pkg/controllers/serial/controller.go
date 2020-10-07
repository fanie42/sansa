package serial

import (
    "bufio"
    "context"
    "fmt"
    "time"

    "github.com/fanie42/sansa/pkg/controllers"
    "github.com/fanie42/sansa/pkg/domains"
    term "github.com/tarm/serial"
)

type controller struct {
    service domains.Service
    config  Config
    port    *term.Port
}

// NewController TODO
func NewController(
    cfg Config,
    svc domains.Service,
) controllers.Controller {
    ctrl := &controller{
        service: svc,
        config:  cfg,
    }
    return ctrl
}

// Listen TODO
func (ctrl *controller) Listen(
    ctx context.Context,
) error {
    sp, err := term.OpenPort(
        &term.Config{
            Name: ctrl.config.Name,
            Baud: ctrl.config.Baud,
        },
    )
    if err != nil {
        return err
    }

    ctrl.port = sp

    time.Sleep(time.Second)
    err = ctrl.port.Flush()
    if err != nil {
        return err
    }

    return ctrl.loop()
}

// Serve TODO
func (ctrl *controller) Serve(
    ctx context.Context,
) error {
    return nil
}

// ListenAndServe TODO
func (ctrl *controller) ListenAndServe(
    ctx context.Context,
) error {
    return nil
}

// still need to implement cancel logic here
func (ctrl *controller) loop() error {
    scanner := bufio.NewScanner(ctrl.port)
    for scanner.Scan() {
        s := time.Now().UTC().Format(
            "2006-01-02 15:04:05.000000",
        ) + ", " + scanner.Text()

        m, err := ctrl.service.Build(s)
        if err != nil {
            fmt.Printf(
                "could not build model from %s with error: %v",
                s,
                err,
            )
        }

        err = ctrl.service.Create(m)
        if err != nil {
            fmt.Printf(
                "could not create model from %v with error: %v",
                m,
                err,
            )
        }
    }
    return fmt.Errorf("scanner exited serial read loop: %v", scanner.Err())
}
