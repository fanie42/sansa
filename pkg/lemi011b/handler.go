package lemi011b

import "os"

// Lemi011b TODO
type Lemi011b struct {
    gateway    *Gateway
    controller *Controller
    presenter  *Presenter
}

// New TODO
func New(options ...Option) *Lemi011b {
    // Default options
    l := &Lemi011b{}

    for _, option := range options {
        option(l)
    }

    svc := newService(l.gateway, l.presenter)
    go controller.Execute(svc)

    return l
}

// Option TODO
type Option func(*Lemi011b)

// WithLocal TODO
func WithLocal(h *local.Handler) Option {
    return func(l *Lemi011b) {
        gateway := newLocalRepo(h)
        l.gateways = append(l.gateways, gateway)
    }
}

// This needs to go in local
type localRepo struct {
    f *os.File
}

func newLocalRepo(h *local.Handler) *localRepo {
    return &localRepo{}
}
