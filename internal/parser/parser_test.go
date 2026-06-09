package parser

import (
	"Parking_Lot/internal/services"
	"testing"
)

func TestParseCreateParkingLot(t *testing.T) {
	service := &services.ParkingServices{}
	parser := CommandParser{
		Services: service,
	}
	parser.ParseCommand("create_parking_lot 6")
	slot := service.ParkCar("A", "white")

	if slot != 1 {
		t.Errorf("Expected slot 1, got %d", slot)
	}
}

func TestParseParkCommand(t *testing.T) {

	service := &services.ParkingServices{}

	parser := CommandParser{
		Services: service,
	}

	parser.ParseCommand("create_parking_lot 2")

	parser.ParseCommand(
		"park RJ-01-AB-1234 White",
	)

	slot := service.ParkCar(
		"RJ-02-XY-9999",
		"Black",
	)

	if slot != 2 {
		t.Errorf(
			"Expected slot 2, got %d",
			slot,
		)
	}
}

func TestParseLeaveCommand(t *testing.T) {

	service := &services.ParkingServices{}

	parser := CommandParser{
		Services: service,
	}

	parser.ParseCommand("create_parking_lot 2")

	parser.ParseCommand(
		"park RJ-01-AB-1234 White",
	)

	parser.ParseCommand("leave 1")

	slot := service.ParkCar(
		"NEW-CAR",
		"Black",
	)

	if slot != 1 {
		t.Errorf(
			"Expected slot 1 after leave, got %d",
			slot,
		)
	}
}

func TestInvalidCommand(t *testing.T) {

	service := &services.ParkingServices{}

	parser := CommandParser{
		Services: service,
	}

	parser.ParseCommand("some_random_command")
}
