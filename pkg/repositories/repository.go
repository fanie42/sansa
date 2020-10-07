package repositories

import "github.com/fanie42/sansa/pkg/domains"

// Repository TODO
type Repository interface {
    Save(m domains.Model) error
    GetByID(id string) (domains.Model, error)
}
