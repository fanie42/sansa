package traditional

import "fmt"

// NotFoundError TODO
type NotFoundError struct {
    ID string
}

// Error TODO
func (e *NotFoundError) Error() string {
    return fmt.Sprintf(
        "device with ID=%q not found",
        e.ID,
    )
}

// CommandFailedError TODO
type CommandFailedError struct {
    ID      string
    Message string
}

// Error TODO
func (e *CommandFailedError) Error() string {
    return fmt.Sprintf(
        "invalid command for device with ID=%q, error: %q",
        e.ID, e.Message,
    )
}
