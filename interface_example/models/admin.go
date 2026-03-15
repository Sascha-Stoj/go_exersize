package models

import (
	"fmt"
)

type Admin struct {
	Name string
	Age  int
}

// Admin implementiert types.User automatisch
func (a Admin) GetName() string {
	return a.Name
}

func (a Admin) GetAge() int {
	return a.Age
}

func (a Admin) SendEmail(message string) error {
	fmt.Printf("📧 Admin email to %s: %s\n", a.Name, message)
	return nil
}

// Zusätzliche Admin-spezifische Methode
func (a Admin) ManageUsers() {
	fmt.Printf("👑 %s is managing users\n", a.Name)
}
