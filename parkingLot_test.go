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
