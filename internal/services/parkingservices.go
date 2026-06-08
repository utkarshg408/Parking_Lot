package services

import (
	"Parking_Lot/internal/models"
	"fmt"
)

type ParkingServices struct {
	parkinglot *models.ParkingLot
}

func (s *ParkingServices) createParkingLot(capacity int) {
	s.parkinglot = &models.ParkingLot{
		Capacity: capacity,
		Slots:    make([]models.Slot, capacity),
	}
	for i := 0; i < capacity; i++ {
    s.parkinglot.Slots[i] = models.Slot{
        Slotnum: i + 1,
        Car: nil,
    }
}
}

func (s *ParkingServices) parkCar(Regnum string, Color string) int {
	//if no slot is available return -1
	if s.parkinglot == nil {
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

func (s *ParkingServices) leavecar(Slotnum int) {
	  if s.parkinglot == nil {
		 fmt.Println("Parking lot not created")
		 return 0
	  }
	   if Slotnum > s.parkinglot.Capacity || Slotnum <= 0{
		  fmt.Println("Invalid Slot number")
		  return 0
	   }
}
