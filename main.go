package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const totalTickets int = 100
	remainingTickets := uint(totalTickets)

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have %v remaining tickets\n", remainingTickets)
	fmt.Printf("Get your tickets now!\n\n")

	var bookings [totalTickets]string

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
	bookings[0] = firstName + " " + lastName + " " + email

	fmt.Printf("Bookings array %v \n", bookings)
	fmt.Printf("Booking space %v \n", bookings[0])
	fmt.Printf("Booking type %T \n", bookings)
	fmt.Printf("Booking length %v \n", len(bookings))

	fmt.Printf("Thanks %v %v for buying %v tickets, you rock. \n You'll receive a confirmation in your email %v along with the tickets.\n", firstName, lastName, userTickets, email)

	fmt.Printf("We have %v remaining tickets \n", remainingTickets)
}
