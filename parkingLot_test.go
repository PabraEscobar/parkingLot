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

func TestVehicleSholdNotBeParkedWhenParkingLotFull(t *testing.T) {
	parking, _ := Newlot(1)
	_, err := parking.Park(car1)
	if err != nil {
		t.Fatalf("test setup failed %v", err)
	}
	_, actualErr := parking.Park(car2)
	expectedErr := errors.New("parking lot is full")
	if actualErr.Error() != expectedErr.Error() {
		t.Fatalf("actualErr should be equal to expectedErr")
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
func TestCarParkingAlreadyParkedInAnySlot(t *testing.T) {
	l, _ := Newlot(5)

	_, err := l.Park(car1)
	if err != nil {
		t.Fatalf("should be able to park car1 %v", err)
	}
	_, err = l.Park(car2)
	if err != nil {
		t.Fatalf("should be able to park car2 %v", err)
	}

	_, err = l.Unpark(car1)
	if err != nil {
		t.Fatalf("should be able to unpark car1 %v", err)
	}

	_, err = l.Park(car2)
	if err == nil {
		t.Fatalf("should give error car already parked")
	}

}

type mockParkingFull struct {
	id             uint
	receivedStatus ParkingStatus
}

func (m *mockParkingFull) ParkingFullReceive(i uint) {
	m.receivedStatus = ParkingFull
	m.id = i
}

func TestNotifiedSubscriberThatParkingFullOfCapacityOne(t *testing.T) {
	parking, _ := Newlot(1)
	mockSubscriber := &mockParkingFull{}
	(parking).AddSubscriberParkingFull(mockSubscriber)
	parking.Park(car1)
	if mockSubscriber.receivedStatus != ParkingFull {
		t.Errorf("When parking lot is full, parking lot should notify the owner that parking lot is full")
	}
}

func TestNotifiedSubscriberThatPakingFull(t *testing.T) {
	parking, _ := Newlot(2)
	mockSubscriber := &mockParkingFull{}
	(parking).AddSubscriberParkingFull(mockSubscriber)
	parking.Park(car1)
	parking.Park(car2)

	if mockSubscriber.receivedStatus != ParkingFull {
		t.Errorf("owner is not notified")
	}
}

func TestNotifiedSubscribersThatParkingFull(t *testing.T) {
	parking, _ := Newlot(2)
	subscriber1 := &mockParkingFull{}
	(parking).AddSubscriberParkingFull(subscriber1)
	subscriber2 := &mockParkingFull{}
	(parking).AddSubscriberParkingFull(subscriber2)
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
	id             uint
	receivedStatus ParkingStatus
}

func (m *mockOwner) Receive(status ParkingStatus, i uint) {
	m.receivedStatus = status
	m.id = i
}

type mockParkingStatusReceiver struct {
	id             uint
	receivedStatus ParkingStatus
}

func (m *mockParkingStatusReceiver) Receive(status ParkingStatus, i uint) {
	m.receivedStatus = status
	m.id = i
}

func TestOnlyNotifiedToOwnerThatParkingAvailableAndFull(t *testing.T) {
	parking, _ := NewlotV2(0, 2)
	owner := &mockOwner{}
	subscriber := &mockParkingStatusReceiver{}
	parking.AddSubscriberParkingStatus(owner)
	parking.AddSubscriberParkingStatus(subscriber)
	parking.Park(car1)
	parking.Park(car2)
	if owner.receivedStatus != ParkingFull {
		t.Errorf("owner is not notified that parking is Full")
	}
	if owner.id != 0 {
		t.Errorf("owner is receiving wrong lot id : %v", owner.id)
	}
	if subscriber.receivedStatus != ParkingFull {
		t.Errorf("owner is not notified that parking is Full")
	}
	if subscriber.id != 0 {
		t.Errorf("owner is receiving wrong lot id : %v", subscriber.id)
	}
	parking.Unpark(car2)
	if owner.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified that parking is available")
	}
	if subscriber.receivedStatus != ParkingAvailable {
		t.Errorf("owner is not notified that parking is available")
	}
}

func TestNotifiedSubscriberThatParkingAvailableWhichSubscribedParkingStatus(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.AddSubscriberParkingStatus(owner)
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
	parking.AddSubscriberParkingStatus(owner)

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
	_, actualErr := lot.Unpark(nil)
	expectedErr := errors.New("nil vehicle cannot be unparked")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("nil vehicle cannot be unparked")
	}
}

func TestNotifySubscriberWhenParkingAvailable(t *testing.T) {
	parking, _ := Newlot(2)
	owner := &mockOwner{}
	parking.AddSubscriberParkingStatus(owner)

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

func TestUnparkCannotDoneForNonExistentVehicle(t *testing.T) {
	lot, _ := Newlot(2)

	lot.Park(car2)
	_, actualErr := lot.Unpark(car1)
	expectedErr := errors.New("vehicle not parked in the parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendance cannot unpark the nonexistent vehicle")
	}
}

func TestNewParkingLotWithId(t *testing.T) {
	lot, err := NewlotV2(1, 2)
	if err != nil {
		t.Fatal("new lot is not created with id 1")
	}
	if lot.id != 1 {
		t.Fatal("id of the lot should be equal to 1")
	}
}

func TestNewlotv2ShouldNotCreateLotWithZeroCapacity(t *testing.T) {
	lot, err := NewlotV2(1, 0)
	if err == nil {
		t.Fatal("lot should not be created with zero capacity")
	}
	if lot != nil {
		t.Fatal("lot should be nil")
	}
}

func TestNotifiedParkingFullWithLotId(t *testing.T) {
	l, _ := NewlotV2(1, 1)
	mockSubscriber := &mockParkingFull{}
	l.AddSubscriberParkingFull(mockSubscriber)
	_, err := l.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed %v", err)
	}

	if mockSubscriber.id != 1 {
		t.Fatal("received id of the lot should be 1")
	}

}
