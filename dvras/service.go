package dvras

// Service TODO
type Service interface {
    Start(command *StartCommand) error
    Stop(command *StopCommand) error
}
