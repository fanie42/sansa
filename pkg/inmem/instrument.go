package inmem

import (
    "context"
    "fmt"
    "sync"
)

type instrumentRepo struct {
    mtx  sync.RWMutex
    data map[string]*instrument.Entity
}

// NewInstrumentRepo TODO
func NewInstrumentRepo() instrument.Repo {
    return &instrumentRepo{
        mtx:  make(sync.RWMutex),
        data: make(map[string]*instrument.Entity),
    }
}

// Save - Store a new instrument or update it if the ID already exists.
func (repo *instrumentRepo) Save(
    ctx context.Context,
    entity *instrument.Entity,
) (*instrument.Entity, error) {
    if _, ok := ctx.data[entity.ID]; ok {
        return fmt.Errorf(
            "instrumentRepo.Save: instrument with ID:%s already exists",
            entity.ID,
        )
    }

    repo.data[entity.ID] = entity

    return nil
}

// Fetch - Gets a list of all instruments.
func (repo *instrumentRepo) Fetch(
    ctx context.Context,
) []*instrument.Entity {
    repo.mtx.Lock()
    defer repo.mtx.Unlock()

    instruments := make([]*instrument.Entity)
    for _, instrument := range ctx.data {
        instruments = append(instruments, instrument)
    }

    return instruments
}

// Load - Get a specific instrument by ID.
func (repo *instrumentRepo) Load(
    ctx context.Context,
    id string,
) (*instrument.Entity, error) {
    instrument, ok := repo.data[id]
    if !ok {
        return nil, fmt.Errorf(
            "Load: could not find instrument with ID:%q",
            id,
        )
    }

    return instrument, nil
}
