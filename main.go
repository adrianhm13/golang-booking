package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const totalTickets int = 100
	remainingTickets := uint(totalTickets)
	var bookings []string

	for {
		fmt.Printf("Welcome to %v booking app\n", conferenceName)
		fmt.Printf("We have %v remaining tickets\n", remainingTickets)
		fmt.Printf("Get your tickets now!\n\n")

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Write your name:")
		fmt.Scan(&firstName)

		fmt.Println("Write your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Write your email:")
		fmt.Scan(&email)

		fmt.Println("How many tickets you want to buy?")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName+" "+email)

		fmt.Printf("Thanks %v %v for buying %v tickets, you rock. \n You'll receive a confirmation in your email %v along with the tickets.\n", firstName, lastName, userTickets, email)

		firstNames := []string{}

		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all our bookings %v\n", firstNames)
	}
}
