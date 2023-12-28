package main

import  "fmt" 

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string
	
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

	bookings[0] = firstName + " " + lastName
	remainingTickets -= userTickets

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The type of the array: %T\n", bookings)
	fmt.Printf("The size of the array: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}