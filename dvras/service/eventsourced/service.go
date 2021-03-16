package eventsourced

import "github.com/fanie42/sansa/dvras"

type service struct {
    data dvras.Repository
    devices
}

// New TODO
func New(
    repository dvras.Repository,
) *dvras.Service {
    return &service{
        repo: repository,
    }
}

// Stop TODO
func (svc *service) AcquireData(
    command *AcquireDataCommand,
) error {
    device, err := svc.devices.Load(command.DeviceID)
    if err != nil {
        return err
    }

    data := dvras.New(dvras.DataID)

    err = sensor.AcquireData(&dvras.Data{
        Timestamp: command.Time,
        Ch1:       command.Ch1,
        Ch2:       command.Ch2,
        PPS:       command.PPS,
    })
    if err != nil {
        return err
    }

    return svc.repo.Save(sensor)
}
