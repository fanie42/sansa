package inmem

import "github.com/fanie42/sansa/example"

type repository struct {
    devices map[example.ID]*Device
}

// Load TODO
func (repo *repository) Load(
    id example.ID,
) (*example.Device, error) {
    devices := repo.devices[id]
}

// Save TODO
func (repo *repository) Save(
    device *example.Device,
) error {

}
