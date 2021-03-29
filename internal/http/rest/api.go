package rest

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/fanie42/sansa"
    "github.com/gorilla/mux"
)

// Config TODO
type Config struct {
    Port uint
}

// API TODO
type API struct {
    r  *mux.Router
    es sansa.EventStore
}

// New TODO
func New(
    router *mux.Router,
    eventstore sansa.EventStore,
) *API {
    api := &API{
        r:  router,
        es: eventstore,
    }

    router.HandleFunc("/api/{id}", api.load).Methods("GET")
    router.HandleFunc("/api", api.save).Methods("PUT")

    return api
}

// Run TODO
func (api *API) Run() {
    http.ListenAndServe(":"+strconv.FormatUint(api.config.Port, 10), api.r)
}

func (api *API) load(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    aggregate, err := api.es.Load(id)
    if err != nil {
        w.respond(
            http.StatusNotFound,
            map[string]string{
                "error": err.Error(),
            },
        )
    }

    w.respond(http.StatusOK, aggregate)
}

func (api *API) save(w http.ResponseWriter, r *http.Request) {
    event := &sansa.Event{}
    err := json.Unmarshal(r.Body, event)
    if err != nil {
        w.respond(
            http.StatusInternalServerError,
            map[string]string{
                "error": err.Error(),
            },
        )
    }

    err := api.es.Save(event)
    if err != nil {
        w.respond(
            http.StatusInternalServerError,
            map[string]string{
                "error": err.Error(),
            },
        )
    }
}

func (w http.ResponseWriter) respond(status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        log.Printf("error: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}
