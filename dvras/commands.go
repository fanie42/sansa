package dvras

// AcquireDataCommand TODO
// type AcquireDataCommand struct {
//     // DeviceID  UUID
//     Timestamp time.Time
//     Ch1       []int16
//     Ch2       []int16
//     PPS       []int16
// }

// StartCommand TODO
type StartCommand struct {
    // DeviceID   UUID
    Annotation string `json:"annotation"`
}

// StopCommand TODO
type StopCommand struct {
    // DeviceID   UUID
    Annotation string `json:"annotation"`
}
