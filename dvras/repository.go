package dvras

// Repository TODO
type Repository interface {
    Load(DataID) (*Data, error)
    Save(*Data) error
}
