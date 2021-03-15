package example

// State TODO
type State int

const (
    // On TODO
    On State = iota
    // Off TODO
    Off
)

// String TODO
func (s *State) String() string {
    switch s {
    case On:
        return "on"
    case Off:
        return "off"
    default:
        return ""
    }
}
