package filesystem

// Config TODO
type Config struct {
    Base BaseConfig `env:"BASE"`
    Path PathConfig `env:"PATH"`
}

// BaseConfig TODO
type BaseConfig struct {
    Prefix    string `env:"PREFIX"`
    Pattern   string `env:"PATTERN"`
    Postfix   string `env:"POSTFIX"`
    Extension string `env:"EXTENSION"`
}

// PathConfig TODO
type PathConfig struct {
    Location string `env:"LOCATION"`
    Pattern  string `env:"PATTERN"`
}
