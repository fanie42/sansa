package main

import (
    "context"
    "fmt"
    "time"

    "github.com/fanie42/sansa/pkg/controllers/nats"
    "github.com/fanie42/sansa/pkg/controllers/serial"
    "github.com/fanie42/sansa/pkg/domains/lemi011"
    "github.com/fanie42/sansa/pkg/repositories/filesystem"
)

func main() {
    repository := filesystem.NewRepository(config.Filesystem)
    service := lemi011.NewService(repository)
    serialController := serial.NewController(config.Serial, service)
    natsController := nats.NewController(config.Nats, service)

    ctx, cancel := context.WithTimeout(
        context.Background(),
        time.Second*30,
    )
    defer cancel()

    go serialController.Listen(ctx)
    go natsController.Serve(ctx)

    <-ctx.Done()
    fmt.Println(ctx.Err())
}
