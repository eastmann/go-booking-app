package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainigTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application!")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainigTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

}