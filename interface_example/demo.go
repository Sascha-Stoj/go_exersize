package main

import (
	"fmt"
	"interface-example/models"
	"interface-example/types"
)

func main() {
	admin := models.Admin{Name: "Max", Age: 30}

	// Admin erfüllt BEIDE Interfaces gleichzeitig!
	var user types.User = admin
	var person types.Person = admin

	fmt.Println("=== Admin erfüllt beide Interfaces! ===\n")

	// Mit User Interface
	fmt.Println("Als User:")
	fmt.Println("Name:", user.GetName())
	fmt.Println("Age:", user.GetAge())
	user.SendEmail("Hello from User interface")

	fmt.Println("\nAls Person:")
	// Mit Person Interface
	fmt.Println("Name:", person.GetName())
	fmt.Println("Age:", person.GetAge())
	person.SendEmail("Hello from Person interface")

	// Beide funktionieren in Funktionen
	fmt.Println("\n=== In Funktionen ===")
	processUser(admin)
	processPerson(admin)
}

func processUser(u types.User) {
	fmt.Printf("processUser: %s\n", u.GetName())
}

func processPerson(p types.Person) {
	fmt.Printf("processPerson: %s\n", p.GetName())
}
