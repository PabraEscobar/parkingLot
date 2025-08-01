package parking

import "testing"

func TestLotCreationWithCapacityFive(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	_, err := Owner.Newlot(5)
	if err != nil {
		t.Errorf("parking lot with capacity 5 should be created")
	}
}

func TestLotShouldNotCreatedWithCapacityZero(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	_, err := Owner.Newlot(0)
	if err == nil {
		t.Errorf("parking lot should not be created with zero capacity")
	}
}

func TestCanMyCarBeParked(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(5)
	vehicle, err := l.Park("KA03T4567")
	if err != nil {
		t.Errorf("Vehicle should be parked")
	}
	if (*vehicle).lotId == uint(0) {
		t.Errorf("slot id cannot be zero")
	}
}

func TestVehicleNumberCannotbeEmpty(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(5)
	_, err := l.Park("")
	if err == nil {
		t.Error("vehicle number cannot be empty")
	}
}

func TestUnparkVehicle(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(5)
	l.Park("KA03T4567")
	vehicle, err := l.Unpark("KA03T4567")
	if err != nil {
		t.Errorf("vehicle should be unparked")
	}
	if vehicle == nil {
		t.Errorf("vehicle should be unparked")
	}
}

func TestUnparkVehicleWhichIsNotParked(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(5)
	_, err := l.Unpark("RJ19PA4141")
	if err == nil {
		t.Errorf("vehichle is not parked with these number")
	}
}

func TestCarAlreadyParked(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(5)
	l.Park("RJ19MS1858")
	_, err := l.Park("RJ19MS1858")
	if err == nil {
		t.Errorf("Car already parked can't be parked again")
	}
}

func TestAvailablityNotificationForFullLot(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(2)
	l.Park("TN39AD1232")
	l.Park("RJ78DE1234")
	res := l.AvailabilityNotification()
	if res != "notify owner that parking lot is full" {
		t.Errorf("parking lot is full")
	}
}

func TestAvailablityNotificationforEmptyLot(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(5)
	res := l.AvailabilityNotification()
	if res != "notify owner parking lot have space" {
		t.Errorf("parking lot is available")
	}
}

func TestNewOwner(t *testing.T) {
	_, err := NewOwner("RS Reddy")
	if err != nil {
		t.Errorf("owner should be created")
	}
}

func TestOwnerNotification(t *testing.T) {
	Owner, _ := NewOwner("RS Reddy")
	Owner.Newlot(5)
	res := Owner.Lots[0].AvailabilityNotification()
	if res != "notify owner parking lot have space" {
		t.Errorf("parking lot is available")
	}
}

func TestParkingFullNotificationToOwner(t *testing.T) {
	Owner, _ := NewOwner("Reddy")
	l, _ := Owner.Newlot(2)
	l.Park("TN39AD1232")
	l.Park("RJ78DE1234")
	res := l.AvailabilityNotification()
	if res != "notify owner that parking lot is full" {
		t.Errorf("parking lot is full")
	}
}
