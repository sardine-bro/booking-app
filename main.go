package main

import  (
	"fmt"
	"strings"
	) 

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	bookings := []string {}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for a name
	
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter your how many tickets you want: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
			fmt.Println("Please try again.")
			continue
		}
	
		remainingTickets -= userTickets
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
		
		firstNames := []string {}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are the first names of the bookings we currently have: %v\n", firstNames)
// if no tickets are left get the fuck out of the program
		noTicketsRemaining := remainingTickets == 0 
		if noTicketsRemaining {
			println("AH AH AH YOU DIDNT SAY THE MAGIC WORD")
			println("We are out of tickets... try again next time!")
			break
		}
	}


}