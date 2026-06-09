package main

import (
	"Parking_Lot/internal/parser"
	"Parking_Lot/internal/services"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Parking Lot")
	fmt.Println("Start your journey by giving the command")
	// taking input
	services := &services.ParkingServices{}
	parser := parser.CommandParser{
		Services: services,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		if command == "exit" {
			break
		}
		parser.ParseCommand(command)
	}

}
