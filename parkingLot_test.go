package parking

import (
	"testing"
)

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
	parking, _ := Newlot(5)
	vehicle, err := parking.Park("KA03T4567")
	if err != nil {
		t.Errorf("Vehicle should be parked")
	}
	if (*vehicle).lotId == uint(0) {
		t.Errorf("slot id cannot be zero")
	}
}

func TestVehicleNumberCannotbeEmpty(t *testing.T) {
	parking, _ := Newlot(5)
	_, err := parking.Park("")
	if err == nil {
		t.Error("vehicle number cannot be empty")
	}
}

func TestUnparkVehicle(t *testing.T) {
	parking, _ := Newlot(5)
	vehicle := "KA03T4567"
	parking.Park(vehicle)
	vehicleUnparked, err := parking.Unpark(vehicle)
	if err != nil {
		t.Errorf("vehicle should be unparked")
	}
	if vehicleUnparked == nil {
		t.Errorf("vehicle should be unparked")
	}
}

func TestUnparkVehicleWhichIsNotParked(t *testing.T) {
	parking, _ := Newlot(5)
	vehicle := "KA03T4567"
	_, err := parking.Unpark(vehicle)
	if err == nil {
		t.Errorf("vehichle is not parked with these number")
	}
}

func TestCarAlreadyParked(t *testing.T) {
	l, _ := Newlot(5)
	l.Park("RJ19MS1858")
	_, err := l.Park("RJ19MS1858")
	if err == nil {
		t.Errorf("Car already parked can't be parked again")
	}
}

type mockParkingFull struct {
	receivedStatus parkingStatus
}

func (m *mockParkingFull) NotifyParkingFull() {
	m.receivedStatus = parkingFull
}

type mockParkingAvailable struct {
	receivedStatus parkingStatus
}

func (m *mockParkingAvailable) NotifyParkingAvailable() {
	m.receivedStatus = parkingAvailable
}
func TestNotifiedSubscriberThatParkingFullOfCapacityOne(t *testing.T) {
	parking, _ := Newlot(1)
	mockSubscriber := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(mockSubscriber)
	vehicle := "TN39AD1232"
	parking.Park(vehicle)
	if mockSubscriber.receivedStatus != parkingFull {
		t.Errorf("When parking lot is full, parking lot should notify the owner that parking lot is full")
	}
}

func TestNotifiedSubscriberThatPakingFull(t *testing.T) {
	parking, _ := Newlot(2)
	mockSubscriber := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(mockSubscriber)
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)

	if mockSubscriber.receivedStatus != parkingFull {
		t.Errorf("owner is not notified")
	}
}

func TestNotifiedSubscriberThatParkingAvailable(t *testing.T) {
	parking, _ := Newlot(2)
	mockSubscriber := &mockParkingAvailable{}
	(parking).SubscribeParkingAvailableStatus(mockSubscriber)
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)
	parking.Unpark(anotherVehicle)

	if mockSubscriber.receivedStatus != parkingAvailable {
		t.Errorf("owner is not notified")
	}
}

func TestNotifiedSubscribersThatParkingFull(t *testing.T) {
	parking, _ := Newlot(2)
	subscriber1 := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(subscriber1)
	subscriber2 := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(subscriber2)
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)

	if subscriber1.receivedStatus != parkingFull {
		t.Errorf("subscriber1 is not notified that parking is full")
	}
	if subscriber2.receivedStatus != parkingFull {
		t.Errorf("subscriber2 is not notified that parking is full")
	}
}

type parkingStatus uint

const (
	unknown parkingStatus = iota
	parkingAvailable
	parkingFull
)

type mockOwner struct {
	receivedStatus parkingStatus
}

func (m *mockOwner) NotifyParkingAvailable() {
	m.receivedStatus = parkingAvailable
}
func (m *mockOwner) NotifyParkingFull() {
	m.receivedStatus = parkingFull
}

func TestOnlyNotifiedToOwnerThatParkingAvailable(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	(parking).SubscribeParkingAvailableStatus(owner)
	(parking).SubscribeParkingFullStatus(owner)
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)
	if owner.receivedStatus != parkingFull {
		t.Errorf("owner is not notified that parking is Full")
	}
	parking.Unpark(anotherVehicle)
	if owner.receivedStatus != parkingAvailable {
		t.Errorf("owner is not notified that parking is available")
	}
}
