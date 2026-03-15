package main

import (
	"fmt"
	"interface-example/models"
	"interface-example/types"
)

func main() {
	fmt.Println("=== Interface Example ===\n")

	// Beide Typen erfüllen das types.User Interface
	admin := models.Admin{Name: "Max Müller", Age: 30}
	customer := models.Customer{Name: "Anna Schmidt", Age: 25, Email: "anna@example.com"}

	// Als Interface-Variablen verwenden
	var user1 types.User = admin
	var user2 types.User = customer

	// Funktion, die mit User Interface arbeitet
	fmt.Println("--- Notify Users ---")
	notifyUser(user1, "System maintenance at 10 PM")
	notifyUser(user2, "Your order has shipped!")

	// Liste von verschiedenen User-Typen
	fmt.Println("\n--- Notify All Users ---")
	users := []types.User{admin, customer}
	notifyAllUsers(users, "Important announcement!")

	// Typ-spezifische Methoden (brauchen den konkreten Typ)
	fmt.Println("\n--- Type-specific methods ---")
	admin.ManageUsers()
	customer.PlaceOrder("Laptop")
}

// Funktion akzeptiert JEDES User Interface
func notifyUser(u types.User, msg string) {
	fmt.Printf("Notifying %s (%d years old)\n", u.GetName(), u.GetAge())
	u.SendEmail(msg)
	fmt.Println()
}

// Funktion für mehrere User
func notifyAllUsers(users []types.User, msg string) {
	for _, user := range users {
		user.SendEmail(msg)
	}
}
