package parking

import (
	"errors"
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

var (
	car1, _ = NewVehicle("TN39AD1232")
	car2, _ = NewVehicle("RJ78DE1234")
)

func TestCanMyCarBeParked(t *testing.T) {
	parking, _ := Newlot(5)
	actualVehicle, err := parking.Park(car1)
	if err != nil {
		t.Errorf("Vehicle should be parked %v", err)
	}
	expectedVehicle := car1
	if !expectedVehicle.Equals(actualVehicle) {
		t.Errorf("actual vehicle should be same as expected vehicle")
	}
}

func TestNilVehicleShouldNotBeParked(t *testing.T) {
	parking, _ := Newlot(5)
	_, err := parking.Park(nil)
	if err == nil {
		t.Error("vehicle number cannot be empty")
	}
}

func TestUnparkVehicle(t *testing.T) {
	parking, _ := Newlot(5)
	parking.Park(car1)
	vehicleUnparked, err := parking.Unpark(car1)
	if err != nil {
		t.Errorf("vehicle should be unparked")
	}
	if vehicleUnparked == nil {
		t.Errorf("vehicle should be unparked")
	}
}

func TestUnparkVehicleWhichIsNotParked(t *testing.T) {
	parking, _ := Newlot(5)
	_, err := parking.Unpark(car1)
	if err == nil {
		t.Errorf("vehichle is not parked with these number")
	}
}

func TestCarAlreadyParked(t *testing.T) {
	l, _ := Newlot(5)

	l.Park(car1)
	_, actualErr := l.Park(car1)
	expectedErr := errors.New("car already parked in parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("Car already parked can't be parked again")
	}
}

type mockParkingFull struct {
	receivedStatus ParkingStatus
}

func (m *mockParkingFull) ParkingFullReceive() {
	m.receivedStatus = ParkingFull
}

func TestNotifiedSubscriberThatParkingFullOfCapacityOne(t *testing.T) {
	parking, _ := Newlot(1)
	mockSubscriber := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(mockSubscriber)
	parking.Park(car1)
	if mockSubscriber.receivedStatus != ParkingFull {
		t.Errorf("When parking lot is full, parking lot should notify the owner that parking lot is full")
	}
}

func TestNotifiedSubscriberThatPakingFull(t *testing.T) {
	parking, _ := Newlot(2)
	mockSubscriber := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(mockSubscriber)
	parking.Park(car1)
	parking.Park(car2)

	if mockSubscriber.receivedStatus != ParkingFull {
		t.Errorf("owner is not notified")
	}
}

func TestNotifiedSubscribersThatParkingFull(t *testing.T) {
	parking, _ := Newlot(2)
	subscriber1 := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(subscriber1)
	subscriber2 := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(subscriber2)
	parking.Park(car1)
	parking.Park(car2)

	if subscriber1.receivedStatus != ParkingFull {
		t.Errorf("subscriber1 is not notified that parking is full")
	}
	if subscriber2.receivedStatus != ParkingFull {
		t.Errorf("subscriber2 is not notified that parking is full")
	}
}

type mockOwner struct {
	receivedStatus ParkingStatus
}

func (m *mockOwner) ParkingStatusReceive(status ParkingStatus) {
	m.receivedStatus = status
}

func TestOnlyNotifiedToOwnerThatParkingAvailableAndFull(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	(parking).subscriberParkingStatus = owner
	parking.Park(car1)
	parking.Park(car2)
	if owner.receivedStatus != ParkingFull {
		t.Errorf("owner is not notified that parking is Full")
	}
	parking.Unpark(car2)
	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified that parking is available")
	}
}

func TestNotifiedSubscriberThatParkingAvailableWhichSubscribedParkingStatus(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.subscriberParkingStatus = owner
	parking.Park(car1)
	parking.Park(car2)
	parking.Unpark(car2)
	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified")
	}
}

func TestNotifiedSubscriberThatParkingAvailableWhichSubscribedParkingStatusEdgeCase(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.subscriberParkingStatus = owner

	parking.Park(car1)
	parking.Park(car2)
	parking.Unpark(car2)

	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified")
	}
}

func TestEqualityForVehicles(t *testing.T) {
	vehicleOne := &vehicle{number: "RJ19"}
	vehicleTwo := &vehicle{number: "RJ19"}
	flag := vehicleOne.Equals(vehicleTwo)
	if flag == false {
		t.Errorf("vehicleOne and vehicleTwo should be equal")
	}
}

func TestEqualityForVehiclesWithDifferentLotID(t *testing.T) {
	vehicleOne := &vehicle{number: "RJ19"}
	vehicleTwo := &vehicle{number: "RJ19"}
	flag := vehicleOne.Equals(vehicleTwo)
	if flag == false {
		t.Errorf("vehicleOne and vehicleTwo should be equal")
	}
}

func TestVehicleNotEqualToNil(t *testing.T) {
	vehicleOne := &vehicle{number: "RJ19"}
	flag := vehicleOne.Equals(nil)
	if flag == true {
		t.Errorf("vehicleOne should not be equal to nil")
	}
}

func TestParkAfterParkingAvailable(t *testing.T) {
	lot, _ := Newlot(1)
	lot.Park(car1)
	lot.Unpark(car1)
	_, err := lot.Park(car1)
	if err != nil {
		t.Errorf("parking should take place after parking available")
	}
}

func TestUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	_, err := lot.Unpark(nil)
	if err == nil {
		t.Errorf("nil vehicle cannot be unparked")
	}
}

func TestNotifySubscriberWhenParkingAvailable(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.subscriberParkingStatus = owner

	parking.Park(car1)
	parking.Park(car2)
	parking.Unpark(car1)

	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified")
	}
}

func TestNewVehicle(t *testing.T) {
	_, err := NewVehicle("23BH6543IS")
	if err != nil {
		t.Errorf("vehicle should be created with provided number")
	}
}

func TestNewVehicleShouldNotCreatedWithEmptyNumber(t *testing.T) {
	_, err := NewVehicle("")
	if err == nil {
		t.Errorf("vehicle should be created with provided number")
	}
}
