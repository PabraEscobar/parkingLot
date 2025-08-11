package parking

import (
	"errors"
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

func TestAttendantParkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	vehicleNumber := "KA03FG2345"
	expectedVehicle := &vehicle{lotId: 1, number: vehicleNumber}
	var actualVehicle *vehicle
	var err error

	//logic to test
	actualVehicle, err = attendant.Park(vehicleNumber)

	//assertions
	if err != nil {
		t.Errorf("attendent should park the vehicle")
	}
	if !expectedVehicle.Equals(actualVehicle) {
		t.Errorf("vehicle number should be match with provided number")
	}
}

func TestAttendentUnparkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	vehicleNumber := "KA03FG2345"
	expectedVehicle := &vehicle{lotId: 1, number: vehicleNumber}
	var actualVehicle *vehicle
	var err error

	//logic to test
	_, _ = attendant.Park(vehicleNumber)
	actualVehicle, err = attendant.Unpark(vehicleNumber)

	//assertions
	if err != nil {
		t.Errorf("attendent should park the vehicle")
	}
	if !expectedVehicle.Equals(actualVehicle) {
		t.Errorf("vehicle number should be match with provided number")
	}
}

func TestAttendantCannotParkWhenParkinFull(t *testing.T) {
	//initalization
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(lot)
	vehicleNumber := "KA03FG2345"
	anotherVehicle := "KA03FG2344"

	var actualErr error

	//logic to test
	attendant.Park(vehicleNumber)
	_, actualErr = attendant.Park(anotherVehicle)
	expectedErr := errors.New("parking is full")

	//assertions
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("actual error should be parking is full")
	}
}

func TestAttendantCannotParkVehicleWhichIsAlreadyParked(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	vehicleNumber := "KA03FG2345"

	attendant.Park(vehicleNumber)
	_, actualErr := attendant.Park(vehicleNumber)
	expectedErr := errors.New("car already parked in parking lot")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park the same vehicle twice")
	}
}

func TestAttendantCannotParkWithEmptyVehicleNumber(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	_, actualErr := attendant.Park("")
	expectedErr := errors.New("vehicle number is mandatory to park")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park vehicle with empty number")
	}
}

func TestAttendantCannotUnparkNonexistentVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	vehicleNumber := "KA03FG2345"
	unparkVehicle := "KA04DF6789"

	attendant.Park(vehicleNumber)
	_, actualErr := attendant.Unpark(unparkVehicle)
	expectedErr := errors.New("vehicle not parked in the parking lot with provided number")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendance cannot unpark the nonexistent vehicle")
	}
}

func TestAttendantCannotUnparkWithEmptyVehicleNumber(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	_, actualErr := attendant.Unpark("")
	expectedErr := errors.New("vehicle number is manadatory to unpark the vehicle")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to unpark with empty vehicle number")
	}
}

func TestAttendantParksInFirstAvailableSlot(t *testing.T) {
	lot, _ := Newlot(3)
	attendant, _ := NewAttendant(lot)
	firstVehicle := "KA03FG2345"
	secondVehicle := "KA04DF6789"

	attendant.Park(firstVehicle)
	attendant.Park(secondVehicle)

	attendant.Unpark(firstVehicle)

	anotherVehicle := "KA03HJ6789"

	v, _ := attendant.Park(anotherVehicle)

	if v.lotId != 1 {
		t.Errorf("vehicle should be parked in the first available slot, got slot %d", v.lotId)
	}
}
