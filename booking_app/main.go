package main

import (
	"booking-app/booking_app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const numberConferenceTicketsAvailable = 50

var remainingConferenceTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for remainingConferenceTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidNumberUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingConferenceTickets)

		if isValidName && isValidEmail && isValidNumberUserTickets {
			remainingConferenceTickets = remainingConferenceTickets - userTickets
			bookings = append(bookings, getUserDataMap(firstName, lastName, email, userTickets))
			fmt.Printf("Thank you %v %v for booking %v Tickets. You will receive a confirmation mail on %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Tickets are still available for %v\n", remainingConferenceTickets, conferenceName)
			fmt.Printf("These are all of our bookings: %v\n", bookings)
		} else {
			if !isValidName {
				fmt.Printf("First name and last name should be at least 2 characters long \n")
			}
			if !isValidEmail {
				fmt.Printf("Email should be at least 5 characters long and contain @ sign \n")
			}
			if !isValidNumberUserTickets {
				fmt.Printf("Number of tickets should be more than 0 and less than remaining tickets \n")
			}
		}
	}
	fmt.Printf("The first names of our bookings are: %v\n", getFirstNames())
}

func getFirstNames() []string {
	userFirstName := []string{}
	for _, val := range bookings {
		userFirstName = append(userFirstName, val["firstName"])
	}
	return userFirstName
}

func greetUsers() {
	fmt.Printf("Welcome to our conference %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v Tickets and %v Tickets are still available\n", numberConferenceTicketsAvailable, remainingConferenceTickets)
	fmt.Println("You can book your tickets here")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your eMail: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets you wanna buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getUserDataMap(firstName string, lastName string, email string, userTickets uint) map[string]string {
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	return userData
}
