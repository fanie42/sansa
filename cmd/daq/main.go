package main

import (
    "fmt"
    "os"

    "github.com/fanie42/sansa/pkg/config/env"
    "github.com/fanie42/sansa/pkg/lemi011b"
    "github.com/fanie42/sansa/pkg/local"
    "github.com/fanie42/sansa/pkg/nats"
    "github.com/fanie42/sansa/pkg/serial"
)

// Config TODO
type Config struct {
    Local  local.Config  `env:"LOCAL"`
    Serial serial.Config `env:"SERIAL"`
    Nats   nats.Config   `env:"NATS"`
}

var config Config

func init() {
    os.Setenv("LOCAL_BASE_PREFIX", "marl111_")
    os.Setenv("LOCAL_BASE_FORMAT", "20060102_15")
    os.Setenv("LOCAL_BASE_POSTFIX", "0000")
    os.Setenv("LOCAL_BASE_EXTENSION", "dat")

    os.Setenv("LOCAL_PATH_LOCATION", "./data")
    os.Setenv("LOCAL_PATH_FORMAT", "./2006/01/02")

    os.Setenv("SERIAL_NAME", "COM5")
    os.Setenv("SERIAL_BAUD", "19200")

    os.Setenv("NATS_URL", "nats://172.18.30.100:4222")
    os.Setenv("NATS_CLUSTER_ID", "marion")
    os.Setenv("NATS_CLIENT_ID", "marl111_daq")

    env.Configure(&config)

    fmt.Println(config)
}

func main() {
    repo := local.NewLemi011bRepo(config.Local)
    pres := nats.NewPresenter(config.Nats)
    svc := lemi011b.NewService(repo, pres)
    ctrl := serial.New(config.Serial, svc)

    ctrl.Run()
}
