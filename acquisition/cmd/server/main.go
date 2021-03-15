package main

import "github.com/fanie42/sansa/example/http/rest"

func main() {
    eventstore := timescaledb.NewEventStore()

    repository := eventsourced.NewRepository(eventstore)

    service := eventsourced.NewService(repository)

    server := rest.NewServer(service)

    server.Run()
}
