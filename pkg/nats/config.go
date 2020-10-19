package nats

// Config TODO
type Config struct {
    URL       string `env:"URL"`
    ClientID  string `env:"CLIENT_ID"`
    ClusterID string `env:"CLUSTER_ID"`
}
