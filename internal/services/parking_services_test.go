package services

import (
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	service := &ParkingServices{}
	service.CreateParkingLot(10)

	if service.parkinglot == nil {
		t.Fatal("Parking Lot not created")

	}
	if service.parkinglot.Capacity < 0 {
		t.Fatal("Invalid Capacity")
	}

	if service.parkinglot.Capacity != 10 {
		t.Errorf("Expected Slots -> 10 , Got -> %d", service.parkinglot.Capacity)

	}
}

func TestParkCar(t *testing.T) {
	service := &ParkingServices{}
	service.CreateParkingLot(2)

	slot := service.ParkCar(
		"RJ-01-JP-20",
		"white",
	)
	if slot != 1 {
		t.Errorf("Expected Slot = 1 , Actual Slot = %d", slot)
	}
}

func TestFullParking(t *testing.T) {
	service := &ParkingServices{}
	service.CreateParkingLot(1)
	service.ParkCar("A", "white")
	slot := service.ParkCar("B", "red")

	if slot != -1 {
		t.Errorf("Expected -1 but got %d", slot)
	}
}

func TestLeaveCar(t *testing.T) {
	service := &ParkingServices{}
	service.CreateParkingLot(2)
	service.ParkCar("A", "white")
	service.LeaveCar(1)

	if service.parkinglot.Slots[0].Car != nil {
		t.Error("Slot should be empty")
	}
}

func InvalidCapacity(t *testing.T) {
	service := &ParkingServices{}
	service.CreateParkingLot(0)
	if service.parkinglot != nil {
		t.Error("Parking Lot should be created")
	}
}

func InvalidSlot(t *testing.T) {
	service := &ParkingServices{}
	// leave without slot
	result := service.LeaveCar(1)
	if result != -1 {
		t.Error("Error! Leaves Car Without Creating Slot")
	}
	// slot more than created ones
	service.CreateParkingLot(2)
	leave := service.LeaveCar(10)

	if leave != -1 {
		t.Error("Error! More than the slot initialised")
	}
	// negative slot

	leave1 := service.LeaveCar(-1)
	if leave1 != -1 {
		t.Error("Negative Slot cant be Chosen")
	}
}

func TestReuseFreedSlot(t *testing.T) {
	service := &ParkingServices{}

	service.CreateParkingLot(2)

	service.ParkCar("A", "White")

	service.LeaveCar(1)

	slot := service.ParkCar("B", "Black")

	if slot != 1 {
		t.Errorf("Expected slot 1, got %d", slot)
	}
}
