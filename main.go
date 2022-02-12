package main

import "fmt"

func main() {
	var conferenceName = "Go Conferecne"
	const conferenceTickets = 50 
	var remainingTickets uint = 50
	fmt.Printf("Hello, welcome to our %v booking aplication!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still avalible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")

	var userName string


	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Printf("Hello %v",userName)

}