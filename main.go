package main
import "fmt"

func main(){
	var conference_name = "Go Conference"
	const consttickets = 50
	var remaining_tickets = 50

	fmt.Printf("Welcome to %v booking application\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are available\n", consttickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend\n")
}