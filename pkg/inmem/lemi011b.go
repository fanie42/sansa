package inmem

import (
    "fmt"
    "sync"

    "github.com/fanie42/sansa/pkg/lemi011b"
)

// THIS IS A GATEWAY

type lemi011bRepo struct {
    mtx  sync.RWMutex
    data map[lemi011b.UUID]*lemi011b.Data
}

// NewLemi011bRepo TODO
func NewLemi011bRepo() lemi011b.DataRepo {
    return &lemi011bRepo{
        data: make(map[lemi011b.UUID]*lemi011b.Data),
    }
}

func (repo *lemi011bRepo) Save(d *lemi011b.Data) error {
    repo.mtx.Lock()
    defer repo.mtx.Unlock()
    repo.data[d.ID] = d
    fmt.Println(d)
    return nil
}
