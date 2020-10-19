package local

// Config TODO
type Config struct {
    Path PathConfig `env:"PATH"`
    Base BaseConfig `env:"BASE"`
}

// PathConfig TODO
type PathConfig struct {
    Location string `env:"LOCATION"`
    Format   string `env:"FORMAT"`
}

// BaseConfig TODO
type BaseConfig struct {
    Prefix    string `env:"PREFIX"`
    Format    string `env:"FORMAT"`
    Postfix   string `env:"POSTFIX"`
    Extension string `env":"EXTENSION"`
}
