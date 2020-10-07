package inmem

import (
    "fmt"

    "github.com/fanie42/sansa/pkg/domains"
    "github.com/fanie42/sansa/pkg/repositories"
)

type repository struct {
    items map[string]domains.Model
}

// NewRepository TODO
func NewRepository() repositories.Repository {
    return &repository{
        items: make(map[string]domains.Model),
    }
}

// Save TODO
func (repo *repository) Save(item domains.Model) error {
    repo.items[item.ID()] = item

    return nil
}

// GetByID TODO
func (repo *repository) GetByID(id string) (domains.Model, error) {
    result, ok := repo.items[id]
    if !ok {
        return nil, fmt.Errorf("no entity with id: %s", id)
    }
    return result, nil
}
