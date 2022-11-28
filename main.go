package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainigTickets uint = 50

	// remainigTickets = -1

	fmt.Printf("conferenceTickets is %T, remainigTickets is %T, conferenceName is %T \n", conferenceTickets, remainigTickets, conferenceName)

	fmt.Printf("Welcome to our %q booking application! \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainigTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// Ask user for his name

	fmt.Println("Please, enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please, enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please, enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please, enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
}
