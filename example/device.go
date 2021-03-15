package example

import "fmt"

// Device TODO
type Device struct {
    id       ID
    deviceID device.ID
    state    State

    events  []Event
    version int
}

// New TODO
func New(
    id ID,
) *Device {
    return &Device{
        id:      id,
        state:   Off,
        events:  []Event{},
        version: 0,
    }
}

// Apply TODO
// func (d *Device) Apply(event Event, new bool) {
//     switch e := event.(type) {
//     case *Started:
//         d.state = On
//     case *Stopped:
//         d.state = Off
//     default:
//         return
//     }

//     if !new {
//         d.version++
//     }
// }

// ID TODO
func (d *Device) ID() ID {
    return d.id
}

// State TODO
func (d *Device) State() State {
    return d.state
}

// Events TODO
func (d *Device) Events() []Event {
    return d.events
}

// Version TODO
func (d *Device) Version() int {
    return d.version
}

// Start TODO
func (d *Device) Start(
    annotation string,
) error {
    switch d.State {
    case On:
        return fmt.Errorf("can't start device that's already on")
    case Off:
        return nil
    default:
        return fmt.Errorf("can't start device from unknown state")
    }

    d.raise(&Started{
        ID:         d.id,
        Annotation: annotation,
    })

    return nil
}

// Stop TODO
func (d *Device) Stop(
    annotation string,
) error {
    switch d.State {
    case On:
        return nil
    case Off:
        return fmt.Errorf("can't start device that's already on")
    default:
        return fmt.Errorf("can't start device from unknown state")
    }

    d.raise(&Stopped{
        ID:         d.id,
        Annotation: annotation,
    })

    return nil
}

func (d *Device) raise(event Event) {
    d.events = append(d.events, event)
    event.Apply(d)
}
