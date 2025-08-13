package parking

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewAttendant(t *testing.T) {
	lot, _ := Newlot(2)
	_, err := NewAttendant(lot)
	if err != nil {
		t.Errorf("Attendant should be exist with parking lot")
	}
}

func TestNewAttendantCannotExistWithNilLot(t *testing.T) {
	_, err := NewAttendant(nil)
	if err == nil {
		t.Errorf("Attendant should not exist without the parking lot")
	}
}

var (
	vehNumber        = "TN39AD1232"
	anotherVehNumber = "RJ78DE1234"
	carOne           = &vehicle{number: vehNumber}
	carTwo           = &vehicle{number: anotherVehNumber}
)

func TestAttendantParkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	expectedVehicle := &vehicle{number: vehNumber}
	var actualVehicle *vehicle
	var err error

	//logic to test
	actualVehicle, err = attendant.Park(carOne)

	//assertions
	if err != nil {
		t.Errorf("attendent should park the vehicle")
	}
	if !expectedVehicle.Equals(actualVehicle) {
		fmt.Println(expectedVehicle.number, actualVehicle.number)
		t.Errorf("vehicle number should be match with provided number")
	}
}

func TestAttendentUnparkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	expectedVehicle := &vehicle{number: vehNumber}
	var actualVehicle *vehicle
	var err error

	//logic to test
	_, _ = attendant.Park(carOne)
	actualVehicle, err = attendant.Unpark(vehNumber)

	//assertions
	if err != nil {
		t.Errorf("attendent should unpark the vehicle")
	}
	if !expectedVehicle.Equals(actualVehicle) {
		t.Errorf("vehicle number should be match with provided number")
	}
}

func TestAttendantCannotParkWhenParkinFull(t *testing.T) {
	//initalization
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(lot)
	var actualErr error

	//logic to test
	attendant.Park(carOne)
	_, actualErr = attendant.Park(carTwo)
	expectedErr := errors.New("parking is full")

	//assertions
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("actual error should be parking is full")
	}
}

func TestAttendantCannotParkVehicleWhichIsAlreadyParked(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	attendant.Park(carOne)
	_, actualErr := attendant.Park(carOne)
	expectedErr := errors.New("car already parked in parking lot")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park the same vehicle")
	}
}

func TestAttendantCannotParkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	_, actualErr := attendant.Park(nil)
	expectedErr := errors.New("vehicle number is mandatory to park")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park vehicle with empty number")
	}
}

func TestAttendantCannotUnparkNonParkedVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	// unparkVehicle := carOne

	attendant.Park(carTwo)
	_, actualErr := attendant.Unpark(vehNumber)
	expectedErr := errors.New("vehicle not parked in the parking lot with provided number")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendance cannot unpark the nonexistent vehicle")
	}
}

func TestAttendantCannotUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	_, actualErr := attendant.Unpark("")

	if actualErr == nil {
		t.Errorf("attendant should not be able to unpark with empty vehicle number")
	}
}

func TestAttendantParksInFirstAvailableSlot(t *testing.T) {
	lot, _ := Newlot(3)
	attendant, _ := NewAttendant(lot)
	attendant.Park(&vehicle{number: vehNumber})
	attendant.Park(carTwo)

	_, err2 := attendant.Unpark(vehNumber)
	if err2 != nil {
		t.Error(err2)
	}
	actualVehicle, err1 := attendant.Park(&vehicle{number: vehNumber})
	expectedVehicle := &vehicle{number: vehNumber}
	if !expectedVehicle.Equals(actualVehicle) {
		t.Errorf("vehicle should be parked in the first available slot %v", err1)
	}
}

func TestParkingVehicleAfterUnparkWhenParkingFull(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	attendant.Park(&vehicle{number: vehNumber})
	attendant.Park(carTwo)

	attendant.Unpark(vehNumber)
	_, err := attendant.Park(&vehicle{number: vehNumber})
	if err != nil {
		t.Errorf("attendent should able to park the vehicle : %v", err)
	}
}
