package postgres

import (
    "context"
    "fmt"
    "os"

    "github.com/fanie42/sansa/pkg/domains"
    "github.com/fanie42/sansa/pkg/repositories"

    "github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
    config Config
    dbpool pgxpool.Pool
}

// NewRepository TODO
func NewRepository(
    config Config,
) repositories.Repository {
    repo := &repository{
        config: config,
    }

    dbpool, err := pgxpool.Connect(
        context.Background(),
        config.URL,
    )

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }

    repo.dbpool = dbpool

    return repo
}

// Save TODO - If the connection is closed, it should put the value on a
// queue to accommodate for when the config is changed??? Persist the queue?
func (repo *repository) Save(m domains.Model) error {
    err = dbpool.QueryRow(
        context.Background(),
        m.SQL(),
    ).Scan(&greeting)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }
}

// GetByID TODO
func (repo *repository) GetByID(id string) (domains.Model, error) {
    result, ok := repo.items[id]
    if !ok {
        return nil, fmt.Errorf("no entity with id: %s", id)
    }
    return result, nil
}
