package lemi011

import (
    "fmt"
    "strconv"
    "strings"
    "time"

    "github.com/fanie42/sansa/pkg/domains"
    "github.com/fanie42/sansa/pkg/repositories"
)

// Implements the domain.Service interface
type lemi011 struct {
    repo    repositories.Repository
    streams []chan domains.Model
}

// NewService TODO
func NewService(
    repository repositories.Repository,
    // options ...Option,
) domains.Service {
    svc := &lemi011{
        repo:    repository,
        streams: []chan domains.Model{},
    }

    // for _, option := range options {
    //     err := option(svc)
    //     if err != nil {
    //         fmt.Println("could not apply option to lemi011")
    //     }
    // }

    // svc.config.Watch()

    return svc
}

// Stream TODO
func (svc *lemi011) Stream() <-chan domains.Model {
    // We firs need to register a new subscriber I think
    stream := make(chan domains.Model)

    svc.streams = append(svc.streams, stream)

    return stream
}

// Build TODO
func (svc *lemi011) Build(s string) (domains.Model, error) {
    ss := strings.Split(s, ", ")
    l := len(ss)
    if l != 5 {
        return nil, fmt.Errorf("expected 4 data fields, but got: %d", l)
    }

    timestamp, err := time.Parse(
        "2006-01-02 15:04:05.000000",
        ss[0],
    )
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse timestamp value: %v to time.Time with error: %v",
            ss[0],
            err,
        )
    }

    x, err := strconv.ParseInt(ss[1], 10, 64)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse x value: %v to int64 with error: %v",
            ss[1],
            err,
        )
    }

    y, err := strconv.ParseInt(ss[2], 10, 64)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse y value: %v to int64 with error: %v",
            ss[2],
            err,
        )
    }

    z, err := strconv.ParseInt(ss[3], 10, 64)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse z value: %v to int64 with error: %v",
            ss[3],
            err,
        )
    }

    temperature, err := strconv.ParseInt(ss[4], 10, 64)
    if err != nil {
        return nil, fmt.Errorf(
            "could not parse temperature value: %v to int64 with error: %v",
            ss[4],
            err,
        )
    }

    return &data{
        id:          "lemi011" + strconv.FormatInt(timestamp.UnixNano(), 10),
        timestamp:   timestamp,
        x:           x,
        y:           y,
        z:           z,
        temperature: temperature,
    }, nil
}

// Create TODO
func (svc *lemi011) Create(m domains.Model) error {
    // Send to all subscribers
    for _, stream := range svc.streams {
        stream <- m
    }

    // svc.stream <- m // This will only work if the stream is actually read from
    // somewhere too.

    // This sends to repo, yes, but what about the others?
    return svc.repo.Save(m)
}
