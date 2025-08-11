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
	vehicleNumber := "KA03T4567"
	actualVehicle, err := parking.Park(vehicleNumber)
	if err != nil {
		t.Errorf("Vehicle should be parked")
	}
	expectedVehicle := vehicle{number: vehicleNumber}
	if !actualVehicle.Equals(&expectedVehicle) {
		t.Errorf("actual vehicle should be same as expected vehicle")
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
	receivedStatus ParkingStatus
}

func (m *mockParkingFull) ParkingFullReceive() {
	m.receivedStatus = ParkingFull
}

func TestNotifiedSubscriberThatParkingFullOfCapacityOne(t *testing.T) {
	parking, _ := Newlot(1)
	mockSubscriber := &mockParkingFull{}
	(parking).SubscribeParkingFullStatus(mockSubscriber)
	vehicle := "TN39AD1232"
	parking.Park(vehicle)
	if mockSubscriber.receivedStatus != ParkingFull {
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
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)

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
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)
	if owner.receivedStatus != ParkingFull {
		t.Errorf("owner is not notified that parking is Full")
	}
	parking.Unpark(anotherVehicle)
	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified that parking is available")
	}
}

func TestNotifiedSubscriberThatParkingAvailableWhichSubscribedParkingStatus(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.subscriberParkingStatus = owner
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"
	parking.Park(vehicle)
	parking.Park(anotherVehicle)
	parking.Unpark(anotherVehicle)

	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified")
	}
}

func TestNotifiedSubscriberThatParkingAvailableWhichSubscribedParkingStatusEdgeCase(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.subscriberParkingStatus = owner
	vehicle := "TN39AD1232"
	anotherVehicle := "RJ78DE1234"

	parking.Park(vehicle)
	parking.Park(anotherVehicle)
	parking.Unpark(vehicle)

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
