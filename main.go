package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conferecne"
	const conferenceTickets = 50 
	var remainingTickets uint = 5
	fmt.Printf("Hello, welcome to our %v booking aplication!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still avalible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")
	firstNames := []string{}

for{

	
	var firstName string
	var lastName string
	var email string
	var tickets uint
	var bookings  []string

	
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&tickets)
	
	bookings = append(bookings,  firstName + " " + lastName) 
	
	if tickets > remainingTickets {
		fmt.Printf("We only have %v tickets remaining, so you ca not book %v", remainingTickets, tickets)
		continue
	}

	remainingTickets = remainingTickets - tickets
	fmt.Printf("Hello %v %v your email adress is %v. You have booked %v tickets\n",firstName, lastName, email, tickets)
	fmt.Printf("%v tickets remaining\n", remainingTickets)


	for _,booking := range bookings {
		var names  = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	} 
		fmt.Println(firstNames)

	if remainingTickets == 0 {
		fmt.Println("No more tickets")
		break
	}
}

}