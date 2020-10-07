package serial

// Config TODO
type Config struct {
    Name string `env:"NAME"`
    Baud int    `env:"BAUD"`
}

// Option TODO
type Option func(ctrl *controller)

// WithName TODO
func WithName(name string) Option {
    return func(ctrl *controller) {
        ctrl.config.Name = name
    }
}

// WithBaud TODO
func WithBaud(baud int) Option {
    return func(ctrl *controller) {
        ctrl.config.Baud = baud
    }
}
