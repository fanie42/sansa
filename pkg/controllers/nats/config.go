package nats

// Config TODO
type Config struct {
    URL       string `env:"URL"`
    ClusterID string `env:"CLUSTERID"`
    ClientID  string `env:"CLIENTID"`
}
