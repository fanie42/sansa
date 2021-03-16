package inmem

import (
    "fmt"

    "github.com/fanie42/sansa/dvras"
)

type repository struct {
    sensors map[dvras.SensorID]*dvras.Sensor
    devices DeviceGateway
}

// New TODO
func New(
    deviceGateway DeviceGateway,
) dvras.Repository {
    return &repository{
        sensors: make(map[dvras.SensorID]*dvras.Sensor),
        devices: deviceGateway,
    }
}

// Load TODO
func (repo *repository) Load(
    id dvras.SensorID,
) (*dvras.Sensor, error) {
    sensor := repo.sensors[id]

    device, err := repo.devices.GetDeviceByID(sensor.DeviceID)
    if err != nil {
        return nil, fmt.Errorf("no device with this id. You need to register the device first")
    }

}

// Save TODO
func (repo *repository) Save(
    sensor *dvras.Sensor,
) error {
    repo.sensors[sensor.ID] = sensor

}
