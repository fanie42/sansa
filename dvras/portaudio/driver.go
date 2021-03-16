package portaudio

import (
    "fmt"
    "time"

    pa "github.com/gordonklaus/portaudio"
)

// Driver TODO
type Driver struct {
    cfg Config
    svc dvras.Service
}

// New TODO
func New(
    config Config,
    service dvras.Service,
) *Driver {
    return &Driver{
        cfg: config,
        svc: service,
    }
}

// Run TODO
func (d *Driver) Run() {
    err := pa.Initialize()
    if err != nil {
        fmt.Printf("failed to initialize portaudio: %v", err)
        return
    }
    defer pa.Terminate()

    buffer := make([][]int16, 3)
    for i := range buffer {
        buffer[i] = make([]int16, d.cfg.SampleRate)
    }

    stream, err := pa.OpenDefaultStream(
        3,
        0,
        float64(d.cfg.SampleRate),
        d.cfg.SampleRate,
        buffer,
    )
    if err != nil {
        fmt.Printf("failed to open portaudio stream: %v", err)
        return
    }
    defer stream.Close()

    err := stream.Start()
    if err != nil {
        fmt.Printf("failed to start portaudio stream: %v", err)
        return
    }
    defer stream.Stop()

    for {
        if err := stream.Read(); err != nil {
            fmt.Printf("error: +%v", err)
            // Check if still connected... Do some error handling
            continue
        }
        timestamp := time.Now()
        d.svc.AcquireDataFrame(
            &dvras.DataFrame{
                Time: timestamp,
                Ch1:  buffer[0],
                Ch2:  buffer[1],
                PPS:  buffer[2],
            },
        )
    }

    return
}
