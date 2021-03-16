package traditional

type service struct {
    wRepo example.WriteRepository
    rRepo ReadRepository
}

// ReadRepository TODO
type ReadRepository interface {
    GetState(example.DeviceID)
}

// New TODO
func New(
    writeRepo example.WriteRepository,
    readRepo ReadRepository,
) example.Service {
    return &service{
        wRepo: writeRepo,
        rRepo: readRepo,
    }
}

// Start TODO
func (svc *service) Start(
    command *example.StartCommand,
) error {
    device, err := svc.wRepo.Load(command.DeviceID)
    if err != nil {
        return &NotFoundError{ID: command.DeviceID}
    }

    err = device.Start()
    if err != nil {
        return &CommandFailedError{
            ID:      command.DeviceID,
            Message: err.String(),
        }
    }

    err = svc.wRepo.Save(device)
}

// Stop TODO
func (svc *service) Stop(
    command *example.StopCommand,
) error {
    device, err := svc.wRepo.Load(command.DeviceID)
    if err != nil {
        return &NotFoundError{ID: command.DeviceID}
    }

    err = device.Stop()
    if err != nil {
        return &CommandFailedError{
            ID:      command.DeviceID,
            Message: err.String(),
        }
    }

    err = svc.wRepo.Save(device)
}

// GetState TODO
func (svc *service) GetState(
    query *example.GetStateQuery,
) (*GetStateResponse, error) {
    device, err := svc.rRepo.GetDeviceByID(query.DeviceID)
    if err != nil {
        return nil, &NotFoundError{ID: command.DeviceID}
    }

    response := &GetStateResponse{
        State: device.State,
    }

    return response, nil
}
