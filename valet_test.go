package parking

import "testing"

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