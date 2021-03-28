package rest

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "reflect"

    "github.com/fanie42/sansa/dvras"
)

// Controller TODO
type Controller struct {
    app dvras.Service
}

// New TODO
func New(
    application dvras.Service,
) *Controller {
    ctrl := &Controller{
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *Controller) Run() {
    http.Handle("/", http.FileServer(http.Dir("C:/Users/Stephanus/Desktop/sansa/web/")))
    http.HandleFunc("/start", ctrl.start)
    http.HandleFunc("/stop", ctrl.stop)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Start TODO
func (ctrl *Controller) start(
    w http.ResponseWriter,
    r *http.Request,
) {
    command := dvras.StartCommand{}
    err := json.NewDecoder(r.Body).Decode(&command)
    fmt.Printf("%v\n", command)
    if err != nil {
        t := reflect.TypeOf(err)

        fmt.Println(err.Error(), t.String())

        response, _ := json.Marshal(
            map[string]string{
                "message": err.Error(),
            },
        )
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write(response)

        return
    }

    // device := ctrl.devices.Load(command.AggregateID)
    // err := command.Execute(device, s.devices.Save)
    // if err != nil {
    //     response, _ := json.Marshal(
    //         map[string]string{
    //             "message": err.Error(),
    //         },
    //     )
    //     w.Header().Set("Content-Type", "application/json")
    //     w.WriteHeader(http.StatusUnprocessableEntity)
    //     w.Write(response)
    // }

    err = ctrl.app.Start(&command)
    if err != nil {
        t := reflect.TypeOf(err)

        fmt.Println(err.Error(), t.String())

        response, _ := json.Marshal(
            map[string]string{
                "message": err.Error(),
            },
        )
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusConflict)
        w.Write(response)
        return
    }

    response, _ := json.Marshal(map[string]string{
        "message": "Successfully started",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)

    return
}

// Stop TODO
func (ctrl *Controller) stop(
    w http.ResponseWriter,
    r *http.Request,
) {
    command := dvras.StopCommand{}
    err := json.NewDecoder(r.Body).Decode(&command)
    fmt.Printf("%v\n", command)
    if err != nil {
        t := reflect.TypeOf(err)

        fmt.Println(err.Error(), t.String())

        response, _ := json.Marshal(
            map[string]string{
                "message": err.Error(),
            },
        )
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write(response)
        return
    }

    // device := ctrl.devices.Load(command.AggregateID)
    // err := command.Execute(device, s.devices.Save)
    // if err != nil {
    //     response, _ := json.Marshal(
    //         map[string]string{
    //             "message": err.Error(),
    //         },
    //     )
    //     w.Header().Set("Content-Type", "application/json")
    //     w.WriteHeader(http.StatusUnprocessableEntity)
    //     w.Write(response)
    // }

    err = ctrl.app.Stop(&command)
    if err != nil {
        t := reflect.TypeOf(err)

        fmt.Println(err.Error(), t.String())

        response, _ := json.Marshal(
            map[string]string{
                "message": err.Error(),
            },
        )
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusConflict)
        w.Write(response)
        return
    }

    response, _ := json.Marshal(map[string]string{
        "message": "Successfully stopped",
    })

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)

    return
}
