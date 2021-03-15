package example

// Command TODO
// type Command interface {
//     Execute(*Device) error
// }

// Start TODO
// func (d *Device) Start(
//     command *StartCommand,
//     record func(*Device) error,
//     save func(*Device, []Event) error,
// ) error {
//     events := []Event{}

//     switch device.State {
//     case Off:
//         err := record(d)
//         if err != nil {
//             return fmt.Errorf("unable to start recording: %v", err)
//         }
//         events = append(
//             events,
//             &StartedEvent{
//                 Annotation: command.Annotation,
//             },
//         )
//     case On:
//         return fmt.Errorf("can't start device that is already on")
//     default:
//         return fmt.Errorf("can't start device from unknown state")
//     }

//     for _, event := range events {
//         event.Apply(device)
//     }

//     return save(device, events)
// }
