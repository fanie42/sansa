package inmem

type device struct {
    id
}

// DeviceGateway TODO
type deviceGateway interface {
    GetDeviceByID() *device
}
