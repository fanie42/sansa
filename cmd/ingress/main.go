package main

import (
    "context"
    "fmt"

    "github.com/fanie42/sansa/pkg/controllers/nats"
    "github.com/fanie42/sansa/pkg/domains/lemi011"
    "github.com/fanie42/sansa/pkg/repositories/filesystem"
)

func main() {
    repository := filesystem.NewRepository(config.Filesystem)
    service := lemi011.NewService(repository)
    natsController := nats.NewController(config.Nats, service)

    ctx := context.Background()

    go natsController.Listen(ctx)

    <-ctx.Done()
    fmt.Println(ctx.Err())
}
