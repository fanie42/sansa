package domains

// Service TODO
type Service interface {
    Build(s string) (Model, error)
    Create(m Model) error
    Stream() <-chan Model
}
