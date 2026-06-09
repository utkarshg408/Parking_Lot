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
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error Opening the file")
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			command := scanner.Text()
			parser.ParseCommand(command)
			if command == "exit" {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
		return
	}

	// Interactive Mode
	fmt.Println("Start your journey by giving the command")
	fmt.Println("Type 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		command := scanner.Text()

		if command == "exit" {
			fmt.Println("Exiting Parking Lot...")
			break
		}

		parser.ParseCommand(command)
	}

}
