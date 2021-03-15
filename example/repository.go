package example

// Repository TODO
type Repository interface {
    Load(ID) (*Device, error)
    Save(*Device) error
}
