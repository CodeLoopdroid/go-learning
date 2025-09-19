package main
import "fmt"

func main(){
	var conference_name = "Go Conference"
	const consttickets = 50
	var remaining_tickets = 50

	fmt.Println("Welcome to %v booking application", conference_name)
	fmt.Println("We have total of %v tickets and %v are available", consttickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend")
}