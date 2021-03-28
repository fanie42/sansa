package timescaledb

import (
    "context"
    "io/ioutil"
    "log"

    "github.com/fanie42/sansa"
    "github.com/jackc/pgx/v4/pgxpool"
)

type eventstore struct {
    db *pgxpool.Pool
}

// New TODO
func New(database *pgxpool.Pool) sansa.EventStore {
    es := &eventstore{
        db: database,
    }

    sql, err := ioutil.ReadFile("init.sql")
    if err != nil {
        log.Printf("no config file: %v", err)
    }

    tag, err := es.db.Exec(context.Background(), sql)
    if err != nil {
        log.Printf("no config file: %v, tag: %v", err, tag)
    }

    return es
}

// Save TODO
func (es *eventstore) Save(aggregate sansa.Aggregate) error {
    sql := "INSERT INTO events (id, aggregate, type, sequence, data) " +
        "VALUES ($1, $2, $3, $4, $5);"

    for _, event := range aggregate.Changes {
        tag, err := es.db.Exec(context.Background(), sql,
            aggregate.ID,
            aggregate.Type,
            aggregate.EventType,
            aggregate.Sequence,
            aggregate.Data,
        )

        if err != nil {
            return err
        }
    }

    return nil
}

// Load TODO
func (es *eventstore) Load(id AggregateID) (sansa.Aggregate, error) {
    sql := "SELECT (id, aggregate, type, sequence, data) " +
        "FROM events WHERE "

    for _, event := range aggregate.Changes {
        tag, err := es.db.Exec(context.Background(), sql,
            aggregate.ID,
            aggregate.Type,
            aggregate.EventType,
            aggregate.Sequence,
            aggregate.Data,
        )

        if err != nil {
            return err
        }
    }

    return nil
}
