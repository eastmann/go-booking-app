package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainigTickets uint = 50
	// Array
	// var bookings [50]string
	// Slice
	// var bookings []string
	// var bookings = []string{}
	bookings := []string{}

	// remainigTickets = -1

	fmt.Printf("conferenceTickets is %T, remainigTickets is %T, conferenceName is %T \n", conferenceTickets, remainigTickets, conferenceName)

	fmt.Printf("Welcome to our %q booking application! \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", conferenceTickets, remainigTickets)
	fmt.Println("Get your tickets here to attend.")

	// var bookings = [50]string{"Nana", "Nicole", "Peter"}
	// var bookings = [50]string{}

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask user for his name

	fmt.Println("Please, enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please, enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please, enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please, enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainigTickets = remainigTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %q. \n", remainigTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v \n", bookings)
}
