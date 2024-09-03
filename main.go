package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kajanan1212/simple-booking-cli/utils"
	// "strconv"
)

const CONFERENCE_TICKETS uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	fullName    string
	email       string
	userTickets int
}

var waitGroup = sync.WaitGroup{}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if len(os.Args) > 1 {
		log.Fatalf(
			"Too many arguments provided. Please remove the following extra arguments: %v",
			os.Args[1:],
		)
	}

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInputs()

		isValidUsername, isValidEmail, isValidUserTickets :=
			utils.ValidateUserInputs(
				firstName,
				lastName,
				email,
				userTickets,
			)

		if !isValidUsername || !isValidEmail || !isValidUserTickets {
			log.Println("Your input data is invalid.")
			continue
		}

		if userTickets > int(remainingTickets) {
			log.Printf(
				"We only have %v tickets remaining, so you can't book %v tickets.\n",
				remainingTickets, userTickets,
			)
			continue
		}

		bookTicket(userTickets, firstName, lastName, email)

		waitGroup.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		log.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			log.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

	waitGroup.Wait()
}

func greetUsers() {
	log.Printf("Welcome to %v booking application\n", conferenceName)

	log.Printf(
		"We have total of %v tickets and %v are still available.\n",
		CONFERENCE_TICKETS, remainingTickets,
	)

	log.Println("Get your tickets here to attend")
}

func getUserInputs() (
	firstName string,
	lastName string,
	email string,
	userTickets int,
) {
	/*
		var (
			firstName   string
			lastName    string
			email       string
			userTickets int
		)
	*/

	log.Print("Enter your first name: ")
	_, err := fmt.Scan(&firstName)
	if err != nil {
		waitGroup.Wait()
		log.Fatalf("Error reading first name: %v", err)
	}

	log.Print("Enter your last name: ")
	_, err = fmt.Scan(&lastName)
	if err != nil {
		waitGroup.Wait()
		log.Fatalf("Error reading last name: %v", err)
	}

	log.Print("Enter your email address: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		waitGroup.Wait()
		log.Fatalf("Error reading email address: %v", err)
	}

	log.Print("Enter number of tickets: ")
	_, err = fmt.Scan(&userTickets)
	if err != nil {
		waitGroup.Wait()
		log.Fatalf("Error reading number of tickets: %v", err)
	}

	return
}

func bookTicket(
	userTickets int,
	firstName string,
	lastName string,
	email string,
) {
	remainingTickets = remainingTickets - uint(userTickets)

	/*
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["fullName"] = firstName + " " + lastName
		userData["email"] = email
		userData["userTickets"] = strconv.FormatInt(int64(userTickets), 10)
	*/

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		fullName:    firstName + " " + lastName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)
	log.Printf("List of booking is %v.\n", bookings)

	log.Printf(
		"Thank you %v %v for booking %v tickets. "+
			"You will receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email,
	)
	log.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v have been booked.\n", userTickets, firstName, lastName)

	fmt.Println()
	fmt.Println("#################")
	log.Printf("Sending ticket:\n\t%q\nto email address %v.\n", ticket, email)
	fmt.Println("#################")
	fmt.Println()

	waitGroup.Done()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking.fullName)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}
