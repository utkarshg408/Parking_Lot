package parser

import (
	"Parking_Lot/internal/services"
	"fmt"
	"strconv"
	"strings"
)

type CommandParser struct {
	services *services.ParkingServices
}

func (cp *CommandParser) ParseCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		fmt.Println("Empty command")
		return
	}
	switch parts[0] {

	case "create_parking_lot":
		capacity, err := strconv.Atoi(parts[1])
		if len(parts) < 2 {
			fmt.Println("Invalid command")
			return
		}
		if err != nil {
			fmt.Println("Invalid capacity")
			return
		}
		cp.services.CreateParkingLot(capacity)
		fmt.Printf("Created a parking lot with %d slots\n", capacity)

	case "park":
		var Regnum, Color string
		if len(parts) < 3 {
			fmt.Println("Invalid Command. Use : park <registration_num> <color>")
			return
		}

		if len(parts) >= 3 {
			Regnum = parts[1]
			Color = parts[2]
		}

		slotnum := cp.services.ParkCar(Regnum, Color)
		if slotnum != -1 {
			fmt.Printf("Allocated slot number: %d\n", slotnum)
		} else {
			fmt.Println("Sorry, parking lot is full")
		}
	case "leave":
		var Slotnum int
		var err error
		if len(parts) < 2 {
			fmt.Println("Invalid Command. Use : leave <slot_number>")
			return
		}

		if len(parts) >= 2 {
			Slotnum, err = strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid slot number")
				return
			}
		}
		cp.services.LeaveCar(Slotnum)

	case "status":
		cp.services.Status()

	case "registration_numbers_for_cars_with_colour":
		var Color string
		if len(parts) < 2 {
			fmt.Println("Invalid Command. Use : registration_numbers_for_cars_with_colour <color>")
			return
		}
		if len(parts) >= 2 {
			Color = parts[1]
		}
		cp.services.GetRegNumByColor(Color)

	case "slot_numbers_for_cars_with_colour":
		var Color string
		if len(parts) < 2 {
			fmt.Println("Invalid Command . Use slot_numbers_for_cars_with_colour <color>")
			return
		}
		if len(parts) >= 2 {
			Color = parts[1]
		}
		cp.services.GetSlotNumByColor(Color)

	case "slot_number_for_registration_number":
		var Regnum string
		if len(parts) < 2 {
			fmt.Println("Invalid Command. Use : slot_number_for_registration_number <color>")
			return
		}
		if len(parts) >= 2 {
			Regnum = parts[1]

		}
		cp.services.GetSlotNumByRegNum(Regnum)

	default:
		fmt.Println("Invalid command")
	}

}
