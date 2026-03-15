package models

import (
	"fmt"
)

type Customer struct {
	Name  string
	Age   int
	Email string
}

// Customer implementiert types.User automatisch
func (c Customer) GetName() string {
	return c.Name
}

func (c Customer) GetAge() int {
	return c.Age
}

func (c Customer) SendEmail(message string) error {
	fmt.Printf("📧 Customer email to %s (%s): %s\n", c.Name, c.Email, message)
	return nil
}

// Zusätzliche Customer-spezifische Methode
func (c Customer) PlaceOrder(item string) {
	fmt.Printf("🛒 %s ordered: %s\n", c.Name, item)
}
