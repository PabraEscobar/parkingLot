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
	_, err := l.Park("KA03T4567")
	if err != nil {
		t.Errorf("Vehicle should be parked")
	}
}

func TestVehicleNumberCannotbeEmpty(t *testing.T) {
	l, _ := Newlot(5)
	_, err := l.Park("")
	if err == nil {
		t.Error("vehicle number cannot be empty")
	}
}

func TestCheckIfSlotAvailable(t *testing.T) {
	l, _ := Newlot(2)
	_, err := l.IsSlotAvailable()
	if err != nil {
		t.Errorf("Slot is available")
	}
}
