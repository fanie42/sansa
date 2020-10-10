package postgres

// Config TODO
type Config struct {
    Host           string `env:"HOST"`
    Port           uint16 `env:"PORT"`
    Database       string `env:"DATABASE"`
    User           string `env:"USER"`
    Password       string `env:"PASSWORD"`
    ConnectTimeout string `env:"CONNECT_TIMEOUT"`
}
