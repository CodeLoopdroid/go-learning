package main
import "fmt"

func main(){
	conference_name  := "Go Conference" // You can’t use := outside a function (like at package level). There you must use var.
	x := 5
	x, y := 10, 20  // x already exists, but y is new → allowed
	fmt.Println(x, y)
	const consttickets int = 50
	var remaining_tickets uint = 50 // Explicitly define the datatype
	
	fmt.Printf("DataType of %T, DataType of %T, DataType of %T", conference_name, consttickets, remaining_tickets)
	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are available\n", consttickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")

	var firstName,lastName,email string
	var userTickets int

	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email")
	fmt.Scan(&email)

	fmt.Println("Enter the ")


	userTickets = 2

	fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)
}