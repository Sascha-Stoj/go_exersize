package types

// User Interface - nur EINMAL definiert
type User interface {
	GetName() string
	GetAge() int
	SendEmail(message string) error
}
