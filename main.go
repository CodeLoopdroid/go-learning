package main

import (
	"fmt"
	"strings"
)

func main() {
	conference_name := "Go Conference" // You can’t use := outside a function (like at package level). There you must use var.
	// x := 5
	// x, y := 10, 20 // x already exists, but y is new → allowed
	// fmt.Println(x, y)
	const consttickets int = 50
	var remaining_tickets uint = 50 // Explicitly define the datatype
	// var booking = [50]int{1,2,3}
	// var booking [50]string
	var booking = []string{}

	fmt.Printf("DataType of %T, DataType of %T, DataType of %T", conference_name, consttickets, remaining_tickets)
	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are available\n", consttickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")

	var firstName, lastName, email string
	var userTickets uint

	for {

		var l2 uint
		fmt.Scan(&l2)

		fmt.Print("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your Email :")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets :")
		fmt.Scan(&userTickets)

		if userTickets < remaining_tickets {

			remaining_tickets = remaining_tickets - userTickets
			// booking[0] = lastName + " " + firstName
			booking = append(booking, lastName+" "+firstName)

			// fmt.Printf("The whole array: %v\n", booking)
			// fmt.Printf("The firts value: %v\n", booking[0])
			// fmt.Printf("Array length: %v\n", len(booking))
			// fmt.Printf("Array type %T\n", booking)

			fmt.Printf("Thankyou %s %v for booking %v tickets. You will receive a conformation email at %v\n", firstName, lastName, email, userTickets)
			fmt.Printf("Remaining tickets %v\n", remaining_tickets)

			arrfirstnam := []string{}

			for _, j := range booking {
				arrfirstnam = append(arrfirstnam, strings.Fields(j)[0])
			}

			// fmt.Printf("These are all our bookings: %v\n", booking)
			fmt.Printf("The first names of booking are: %v\n", arrfirstnam)
			cond := remaining_tickets == 0 
			if cond {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		}else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remaining_tickets, userTickets)
		}
	}
}
