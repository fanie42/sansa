package example

// Event TODO
type Event interface {
    Apply(*Device)
}

// StartedEvent TODO
type StartedEvent struct {
    ID         DeviceID `json:"id"`
    Annotation string   `json:"annotation"`
}

// StoppedEvent TODO
type StoppedEvent struct {
    ID         DeviceID `json:"id"`
    Annotation string   `json:"annotation"`
}

// Apply TODO
func (event *StartedEvent) Apply(
    device *Device,
) {
    device.State = On
}

// Apply TODO
func (event *StoppedEvent) Apply(
    device *Device,
) {
    device.State = Off
}
