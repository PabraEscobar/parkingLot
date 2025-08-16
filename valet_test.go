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

func TestAttendantParkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)
	expectedVehicle := car1
	var actualVehicle *vehicle
	var err error

	//logic to test
	actualVehicle, err = attendant.Park(car1)

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
	expectedVehicle := car1
	var actualVehicle *vehicle
	var err error

	//logic to test
	_, _ = attendant.Park(car1)
	actualVehicle, err = attendant.Unpark(car1)

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
	attendant.Park(car1)
	_, actualErr = attendant.Park(car2)
	expectedErr := errors.New("parking is full")

	//assertions
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("actual error should be parking is full")
	}
}

func TestAttendantCannotParkVehicleWhichIsAlreadyParked(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	attendant.Park(car1)
	_, actualErr := attendant.Park(car1)
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

	attendant.Park(car2)
	_, actualErr := attendant.Unpark(car1)
	expectedErr := errors.New("vehicle not parked in the parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendance cannot unpark the nonexistent vehicle")
	}
}

func TestAttendantCannotUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	_, actualErr := attendant.Unpark(nil)

	if actualErr == nil {
		t.Errorf("attendant should not be able to unpark with empty vehicle number")
	}
}

func TestAttendantParksInFirstAvailableSlot(t *testing.T) {
	lot, _ := Newlot(3)
	attendant, _ := NewAttendant(lot)
	attendant.Park(car1)
	attendant.Park(car2)

	_, err2 := attendant.Unpark(car1)
	if err2 != nil {
		t.Error(err2)
	}
	actualVehicle, err1 := attendant.Park(car1)
	expectedVehicle := car1
	if !expectedVehicle.Equals(actualVehicle) {
		t.Errorf("vehicle should be parked in the first available slot %v", err1)
	}
}

func TestParkingVehicleAfterUnparkWhenParkingFull(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(lot)

	attendant.Park(car1)
	attendant.Park(car2)

	attendant.Unpark(car1)
	_, err := attendant.Park(car1)
	if err != nil {
		t.Errorf("attendent should able to park the vehicle : %v", err)
	}
}

func TestAttendantManageMultipleLot(t *testing.T) {
	//intialization
	lot, _ := NewlotV2(1, 1)
	anotherLot, _ := NewlotV2(2, 1)

	//business logic
	attendant, err := NewAttendant(lot, anotherLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(attendant.lots) != 2 {
		t.Fatal("attendant should have 2 lots")
	}
}

func TestAttendantCanParkCarInNextLot(t *testing.T) {
	lot, _ := NewlotV2(0, 1)
	anotherLot, _ := NewlotV2(1, 1)
	attendant, _ := NewAttendant(lot, anotherLot)

	_, err := attendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = attendant.Park(car2)
	if err != nil {
		t.Fatalf("park setup failed for car2 %v", err)
	}
}
