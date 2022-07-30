package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var confrenceName = "Go Conference"

const confrenceTickets = 100

var remainingTickets uint = 60
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Print("Hello World")
	// untuk different line gunakan

	greeting()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidatorUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := cetakFirstName()
		fmt.Printf("ini adalah nama costumer yang booking: %v\n", firstNames)

		// tell the pembeli the tickets is sold out
		// using if

		if remainingTickets == 0 {
			fmt.Println("Our Conference booked is over. You can join next year")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("Firstname or lastname is too short")
		}
		if !isValidEmail {
			fmt.Println("Email Yang Kamu Masukan Tidak Valid")

		}
		if !isValidTicketNumber {
			fmt.Println("Number of Tickets invalid")

		}

	}
	wg.Wait()

}

// list
// var bookings = [60]string{"Jojo", "Bizare", "Monkey", "Luffy", "Roronoa", "Zoro"}
// slice

func greeting() {
	fmt.Printf(" Welcome to %v Our Conference\n", confrenceName)
	fmt.Printf("We have total %v tickets and %v are still avaiable\n", confrenceTickets, remainingTickets)
}
func cetakFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// // array
	// var bookings [50]string
	// slice

	// scanln : digunakan untuk menyimpan variabel
	fmt.Println("Enter your Firstname : ")
	// & adalah pointer yang akan mereference ke memory yang kita tuju(usertickets)
	fmt.Scan(&firstName)
	fmt.Println("Enter your Lastname : ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email : ")
	fmt.Scan(&email)
	fmt.Println("how Many tickets you want : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Booking is %v\n", bookings)

	// fmt.Printf("All Array : %v\n", bookings)
	// fmt.Printf("isi Array : %v\n", bookings[0])

	// fmt.Printf("Value of Array : %T\n", bookings)

	// fmt.Printf("Lenght array: %v\n", len(bookings))

	fmt.Printf("Terimakasih %v %v for booking %v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("you'll receive confirmation to your email at %v\n", email)
	fmt.Printf("%v tickets for %v\n", remainingTickets, confrenceName)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("sending tickets: %v\n to email address %v\n", tickets, email)
	fmt.Println("##############")
	wg.Done()
}
