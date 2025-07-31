package parking

import "testing"

func TestLotCreationWithCapacityFive(t *testing.T) {
	_, err := Newlot(5)
	if err != nil {
		t.Errorf("parking lot with capacity 5 should be created")
	}
}

func TestLotShouldNotCreatedWithCapacityZero(t *testing.T) {
	_, err := Newlot(0)
	if err == nil {
		t.Errorf("parking lot should not be created with zero capacity")
	}
}

func TestCanMyCarBeParked(t *testing.T) {
	l, _ := Newlot(5)
	vehicle, err := l.Park("KA03T4567")
	if err != nil {
		t.Errorf("Vehicle should be parked")
	}
	if (*vehicle).lotId == uint(0) {
		t.Errorf("slot id cannot be empty")
	}
}

func TestVehicleNumberCannotbeEmpty(t *testing.T) {
	l, _ := Newlot(5)
	_, err := l.Park("")
	if err == nil {
		t.Error("vehicle number cannot be empty")
	}
}
