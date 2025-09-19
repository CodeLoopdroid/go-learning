package main
import "fmt"

func main(){
	var conference_name string = "Go Conference"
	const consttickets int = 50
	var remaining_tickets uint = 50 // Explicitly define the datatype
	
	fmt.Printf("DataType of %T, DataType of %T, DataType of %T", conference_name, consttickets, remaining_tickets)
	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are available\n", consttickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")

	var newname string
	var userTickets int

	newname = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets\n", newname, userTickets)
}