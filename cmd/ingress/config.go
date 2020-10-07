package main

import (
    "os"

    "github.com/fanie42/sansa/pkg/configurators/env"
    "github.com/fanie42/sansa/pkg/controllers/nats"
    "github.com/fanie42/sansa/pkg/repositories/filesystem"
)

// Config TODO
type Config struct {
    Filesystem filesystem.Config `env:"FILESYSTEM"`
    Nats       nats.Config       `env:"NATS"`
    // Lemi011    lemi011.Config    `env:"LEMI011"`
}

var config Config

func init() {
    os.Setenv("FILESYSTEM_BASE_PREFIX", "marlem1_")
    os.Setenv("FILESYSTEM_BASE_PATTERN", "20060102_15")
    os.Setenv("FILESYSTEM_BASE_POSTFIX", "0000")
    os.Setenv("FILESYSTEM_BASE_EXTENSION", "dat")

    os.Setenv("FILESYSTEM_PATH_LOCATION", "./data")
    os.Setenv("FILESYSTEM_PATH_PATTERN", "./2006/01/02")

    os.Setenv("NATS_URL", "nats://172.18.30.100:4222")
    os.Setenv("NATS_CLUSTERID", "marion")
    os.Setenv("NATS_CLIENTID", "lemi011_ingress")

    env.Configure(&config)
}
