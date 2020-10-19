package serial

// Config TODO
type Config struct {
    Name string `env:"NAME"`
    Baud int    `env:"BAUD"`
}
