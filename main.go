package main

import "fmt"

func main() {
	var conferenceName = "Go Conferecne"
	const conferenceTickets = 50 
	var remainingTickets = 50
	fmt.Printf("Hello, welcome to our %v booking aplication!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still avalible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")

}