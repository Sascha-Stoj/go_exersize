package types

// Ein zweites Interface mit den GLEICHEN Methoden wie User!
type Person interface {
	GetName() string
	GetAge() int
	SendEmail(message string) error
}
