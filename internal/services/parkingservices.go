package services

import (
	"Parking_Lot/internal/models"
	"fmt"
)

type ParkingServices struct {
	parkinglot *models.ParkingLot
}

func (s *ParkingServices) CreateParkingLot(capacity int) {
	if capacity <= 0 {
		fmt.Println("Invalid Slots Entered")
		return
	}
	s.parkinglot = &models.ParkingLot{
		Capacity: capacity,
		Slots:    make([]models.Slot, capacity),
	}
	for i := 0; i < capacity; i++ {
		s.parkinglot.Slots[i] = models.Slot{
			Slotnum: i + 1,
			Car:     nil,
		}
	}
}

func (s *ParkingServices) ParkCar(Regnum string, Color string) int {
	//if no slot is available return -1
	if s.parkinglot == nil {
		fmt.Println("Parking Lot not created")
		return -1
	}
	// if slot available park the car and return slot number
	for i := 0; i < s.parkinglot.Capacity; i++ {

		if s.parkinglot.Slots[i].Car == nil {
			s.parkinglot.Slots[i].Car = &models.Car{
				RegNum: Regnum,
				Color:  Color,
			}
			return s.parkinglot.Slots[i].Slotnum
		}
	}
	return -1
}

func (s *ParkingServices) LeaveCar(Slotnum int) int {
	if s.parkinglot == nil {
		fmt.Println("Parking lot not created")
		return -1
	}
	if Slotnum > s.parkinglot.Capacity || Slotnum <= 0 {
		fmt.Println("Invalid Slot number")
		return -1
	}
	if s.parkinglot.Slots[Slotnum-1].Car == nil {
		fmt.Println("Slot is already empty")
		return -1

	}
	if s.parkinglot.Slots[Slotnum-1].Car != nil {
		s.parkinglot.Slots[Slotnum-1].Car = nil
		fmt.Println("Slot number ", Slotnum, "is free")
	}
	return 1
}

func (s *ParkingServices) Status() {
	// checking all the occupied slots and printing the details
	if s.parkinglot == nil {
		fmt.Println("Parking lot is not created")
		return
	}
	//anycar is parked
	fmt.Printf("SlotNo.\tRegNo\tColour\n")
	for i := 0; i < s.parkinglot.Capacity; i++ {
		if s.parkinglot.Slots[i].Car != nil {
			fmt.Printf("%d\t%s\t%s\n", s.parkinglot.Slots[i].Slotnum, s.parkinglot.Slots[i].Car.RegNum, s.parkinglot.Slots[i].Car.Color)
		}
	}

}

func (s *ParkingServices) GetRegNumByColor(Color string) {
	//checking all the slots and printing the registration number of cars with given color
	if s.parkinglot == nil {
		fmt.Println("Parking lot is not created")
		return
	}
	fmt.Printf("Registration numbers of cars with color %s:\n", Color)
	var found bool
	for i := 0; i < s.parkinglot.Capacity; i++ {
		if s.parkinglot.Slots[i].Car != nil && s.parkinglot.Slots[i].Car.Color == Color {
			fmt.Printf("%s\n", s.parkinglot.Slots[i].Car.RegNum)
			found = true
		}
	}
	if !found {
		fmt.Println("Not Found")
	}
}

func (s *ParkingServices) GetSlotNumByColor(Color string) {
	if s.parkinglot == nil {
		fmt.Println("Parking lot is not created")
		return
	}
	fmt.Printf("Slot numbers of cars with color %s:\n", Color)
	var found bool
	for i := 0; i < s.parkinglot.Capacity; i++ {
		if s.parkinglot.Slots[i].Car != nil && s.parkinglot.Slots[i].Car.Color == Color {
			fmt.Printf("%d\n", s.parkinglot.Slots[i].Slotnum)
			found = true
		}
	}
	if !found {
		fmt.Println("Not Found")
	}
}

func (s *ParkingServices) GetSlotNumByRegNum(Regnum string) {
	if s.parkinglot == nil {
		fmt.Println("Parking lot is not created")
		return
	}
	fmt.Printf("Slot number of car with registration number %s:\n", Regnum)
	var found bool
	for i := 0; i < s.parkinglot.Capacity; i++ {
		if s.parkinglot.Slots[i].Car != nil && s.parkinglot.Slots[i].Car.RegNum == Regnum {
			fmt.Printf("%d\n", s.parkinglot.Slots[i].Slotnum)
			found = true
		}
	}
	if !found {
		fmt.Println("Not Found")
	}
}
