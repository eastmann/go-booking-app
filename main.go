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
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for his name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %q booked %v tickets.\n", userName, userTickets)
}
