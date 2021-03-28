package main

import (
    "context"
    "log"
    "net/http"

    "github.com/fanie42/sansa/dvras/http/rest"
    "github.com/fanie42/sansa/internal/timescaledb"
    "github.com/jackc/pgx/v4/pgxpool"
)

func main() {
    dbpool, err := pgxpool.Connect(
        context.Background(),
        "postgres://postgres:admin@172.18.30.100:5432/eventstore",
    )
    if err != nil {
        log.Fatalf("could not connect to timescaledb: %v", err)
    }
    defer dbpool.Close()

    eventstore := timescaledb.New(dbpool)

    api := rest.New(eventstore)
    log.Fatal(http.ListenAndServe(":8080", api))
}
