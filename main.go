package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers(conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your e-mail address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want?: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 1 && len(lastName) > 1
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v.\n", firstName, lastName, userTickets, email)

			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("You have entered a short name.")
			}
			if !isValidEmail {
				fmt.Println("Your e-mail address is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Enter a correct number of tickets to proceed.")
			}
		}
	}
}

func greetUsers (confName string){
	fmt.Printf("Welcome to our %v booking application\n", confName)
}