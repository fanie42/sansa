package eventsourced

type service struct {
    repo Repository
}

// New TODO
func New(
    repository Repository,
) *Service {
    return &service{
        repo: repository,
    }
}

// Start TODO
func (svc *service) Start(
    command *StartCommand,
) error {
    device, err := svc.repo.Load(command.DeviceID)
    if err != nil {
        return err
    }

    err = device.Start(command.Annotation)
    if err != nil {
        return err
    }

    return svc.repo.Save(device)
}

// Stop TODO
func (svc *service) Stop(
    command *StopCommand,
) error {
    device, err := svc.repo.Load(command.DeviceID)
    if err != nil {
        return err
    }

    err = device.Stop(command.Annotation)
    if err != nil {
        return err
    }

    return svc.repo.Save(device)
}
