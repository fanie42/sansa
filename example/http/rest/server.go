package rest

import (
    "encoding/json"
    "net/http"

    "github.com/fanie42/sansa/example"
    "github.com/fanie42/sansa/example/service/eventsourced"
)

// THIS CAN BE OUR SERVICE DIRECTLY!! Because we already implemented ALLL
// domain logic via commands! We can have the repo here directly instead of
// the service.

// Server TODO
type Server struct {
    devices eventsourced.Devices
}

// NewServer TODO
func NewServer(
    devices eventsourced.Devices,
) *Server {
    s := &Server{
        devices: repository,
    }

    return s
}

// Start TODO
func (s *Server) Start(
    w http.ResponseWriter,
    r *http.Request,
) {
    command := eventsourced.StartCommand{}
    err := json.Unmarshal(r.Body, &command)
    if err != nil {
        response, _ := json.Marshal(map[string]string{
            "message": err.Error(),
        })

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write(response)
    }

    device := s.devices.Load(command.AggregateID)
    err := command.Execute(device, s.devices.Save)
    if err != nil {
        response, _ := json.Marshal(map[string]string{
            "message": err.Error(),
        })

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write(response)
    }

    response, _ := json.Marshal(map[string]string{
        "message": "Successfully Started",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)

    return
}

// Stop TODO
func (s *Server) Stop(
    w http.ResponseWriter,
    r *http.Request,
) {
    command := example.StopCommand{}
    err := json.Unmarshal(r.Body, &command)
    if err != nil {
        response, _ := json.Marshal(map[string]string{
            "message": err.Error(),
        })

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write(response)
    }

    device := s.devices.Load(command.AggregateID)
    err := command.Execute(device, s.devices.Save)
    if err != nil {
        response, _ := json.Marshal(map[string]string{
            "message": err.Error(),
        })

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write(response)
    }

    response, _ := json.Marshal(map[string]string{
        "message": "Successfully Started",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)

    return
}
