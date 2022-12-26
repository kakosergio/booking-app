package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// create a wait group for threads to be executed
var wg = sync.WaitGroup{}

func main() {

	greetUsers()


		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// add a new thread to be waited
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			getFirstNames()

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
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
		// tells the program to wait for the thread
		wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n%v\nto e-mail address %v\n", ticket, email)
	fmt.Println("##############")

	// tells the program that the thread is finished
	wg.Done()
}
