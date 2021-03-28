package main

import (
    "fmt"

    "github.com/fanie42/sansa/dvras"
    "github.com/fanie42/sansa/dvras/http/rest"
    "github.com/fanie42/sansa/dvras/portaudio"
    "github.com/fanie42/sansa/dvras/repository/inmem"
    "github.com/google/uuid"
    pa "github.com/gordonklaus/portaudio"
)

func main() {
    err := pa.Initialize()
    if err != nil {
        fmt.Printf("failed to initialize portaudio: %v", err)
        return
    }
    defer pa.Terminate()

    repo := inmem.New()
    app := portaudio.New(
        &portaudio.Config{
            SampleRate: 44100,
            DeviceID:   dvras.DeviceID(uuid.New()),
        },
        repo,
    )
    controller := rest.New(app)

    controller.Run()
}
