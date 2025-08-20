package parking

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)
	// attendant.plan = Simple
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
		t.Errorf("vehicle number should be match with provided number")
	}
}

func TestAttendentUnparkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)
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
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)
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
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)

	attendant.Park(car1)
	_, actualErr := attendant.Park(car1)
	expectedErr := errors.New("car already parked in parking lot")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park the same vehicle")
	}
}

func TestAttendantCannotParkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)
	attendant.Park(car1)
	_, actualErr := attendant.Park(nil)
	expectedErr := errors.New("nil vehicle cannot be parked")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park nil vehicle")
	}
}

func TestAttendantCannotUnparkNonParkedVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)

	attendant.Park(car2)
	_, actualErr := attendant.Unpark(car1)
	expectedErr := errors.New("vehicle not parked in the parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendance cannot unpark the nonexistent vehicle")
	}
}

func TestAttendantCannotUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)
	_, actualErr := attendant.Unpark(nil)

	expectedErr := errors.New("nil vehicle cannot be unparked")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to unpark nil vehicle")
	}
}

func TestAttendantParksInFirstAvailableSlot(t *testing.T) {
	lot, _ := Newlot(3)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)
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
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot)

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
	attendant, err := NewAttendantv2(ParkInFirstEmptyLot, lot, anotherLot)

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
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot, anotherLot)

	_, err := attendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = attendant.Park(car2)
	if err != nil {
		t.Fatalf("park setup failed for car2 %v", err)
	}

	if anotherLot.vehicles[0] == nil {
		t.Fatal("attendant should able to park in another Lot")
	}
}

func TestAttendantUnparkVehicleParkedInAnyLot(t *testing.T) {
	lot, _ := NewlotV2(0, 1)
	anotherLot, _ := NewlotV2(1, 1)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot, anotherLot)

	_, err := attendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = attendant.Park(car2)
	if err != nil {
		t.Fatalf("park setup failed for car2 %v", err)
	}

	_, err = attendant.Unpark(car2)
	if err != nil {
		t.Fatalf("unpark setup failed for car2 %v", err)
	}

	if anotherLot.vehicles[0] != nil {
		t.Fatal("attendant should able to unpark from another Lot")
	}

}

func TestAttendantCanParkSameVehicleAfterUnPark(t *testing.T) {
	lot, _ := NewlotV2(0, 1)
	anotherLot, _ := NewlotV2(1, 1)
	attendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot, anotherLot)

	_, err := attendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = attendant.Park(car2)
	if err != nil {
		t.Fatalf("park setup failed for car2 %v", err)
	}

	_, err = attendant.Unpark(car1)
	if err != nil {
		t.Fatalf("unpark setup failed for car1 %v", err)
	}

	_, err = attendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 after unpark %v", err)
	}
}

func TestNewAttendantv2(t *testing.T) {
	lot, _ := NewlotV2(0, 1)
	_, err := NewAttendantv2(1, lot)
	if err != nil {
		t.Fatal("NewAttendant should be created")
	}
}

func TestNewAttendantv2ShouldNotCreateAttendantWithNilLot(t *testing.T) {
	lot, _ := NewlotV2(0, 1)

	_, err := NewAttendantv2(1, lot, nil)
	if err == nil {
		t.Fatal("NewAttendant should be created with empty lot")
	}
}

func TestParkUsingAttendantWithComplexPlan(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 2)

	complexAttendant, _ := NewAttendantv2(ParkInLeastFilledLot, lot1, lot2)

	_, err := complexAttendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}
	_, err = complexAttendant.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 %v", err)
	}

	if lot2.vehicles[0] == nil {
		t.Fatal("attendant should park in lot2")
	}
}

func TestAttendantWithComplexPlanCannotParkVehicleWhenParkingFull(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 1)

	complexAttendant, _ := NewAttendantv2(ParkInLeastFilledLot, lot1, lot2)
	car3 := &vehicle{number: "Rj19"}
	car4 := &vehicle{number: "Rj18"}

	_, err := complexAttendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}
	_, err = complexAttendant.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 %v", err)
	}

	_, err = complexAttendant.Park(car3)
	if err != nil {
		t.Fatalf("park setup failed for car3 %v", err)
	}
	_, err = complexAttendant.Park(car4)
	if err == nil {
		t.Fatal("car4 cannot be parked parking lot is full")
	}
}

func TestBothAttendantShouldBeAbleToParkAndUnpark(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 2)

	car3 := &vehicle{number: "car3"}
	car4 := &vehicle{number: "car4"}

	simpleAttendant, _ := NewAttendantv2(ParkInFirstEmptyLot, lot1, lot2)
	complexAttendant, _ := NewAttendantv2(ParkInLeastFilledLot, lot1, lot2)

	_, err := simpleAttendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = simpleAttendant.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 by simpleAttendant %v", err)
	}

	if car1.Equals(lot1.vehicles[0]) == false {
		t.Fatal("attendant should park car1 in lot1")
	}

	_, err = complexAttendant.Park(car3)
	if err != nil {
		t.Fatalf("Park setup failed for car3 by complexAttendant %v", err)
	}

	_, err = simpleAttendant.Unpark(car1)
	if err != nil {
		t.Fatalf("unpark setup failed for car1 by simpleAttendant %v", err)
	}

	_, err = complexAttendant.Unpark(car2)
	if err != nil {
		t.Fatalf("unpark setup failed for car2 by complexAttendant %v", err)
	}

	_, err = complexAttendant.Park(car4)
	if err != nil {
		t.Fatalf("Park setup failed for car2 by complexAttendant %v", err)
	}
	if car4.Equals(lot1.vehicles[0]) == false {
		t.Fatalf("car should have been parked in first slot but was parked in lot 2 %v", lot2.vehicles[1])
	}
}

func TestParkUsingAttendantWithMaximumFilledLotParkingPlan(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 2)

	attendantWithMaxFilledLotPlan, _ := NewAttendantv2(ParkInMaximumFilledLot, lot1, lot2)

	_, err := attendantWithMaxFilledLotPlan.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = attendantWithMaxFilledLotPlan.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 %v", err)
	}

	if car2.Equals(lot1.vehicles[1]) == false {
		t.Fatal("attendant should park in lot2")
	}
}

func TestAttendantWithMaximumFilledLotPlanCannotParkVehicleWhenParkingFull(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 1)

	attendantWithMaxFilledLotPlan, _ := NewAttendantv2(ParkInMaximumFilledLot, lot1, lot2)
	car3 := &vehicle{number: "Rj19"}
	car4 := &vehicle{number: "Rj18"}

	_, err := attendantWithMaxFilledLotPlan.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}
	_, err = attendantWithMaxFilledLotPlan.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 %v", err)
	}

	_, err = attendantWithMaxFilledLotPlan.Park(car3)
	if err != nil {
		t.Fatalf("park setup failed for car3 %v", err)
	}
	_, err = attendantWithMaxFilledLotPlan.Park(car4)
	if err == nil {
		t.Fatal("car4 cannot be parked parking lot is full")
	}
}

func TestAllTheAttendantsAbleToParkAndUnpark(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 2)

	car3 := &vehicle{number: "car3"}
	car4 := &vehicle{number: "car4"}
	car5 := &vehicle{number: "car5"}
	car6 := &vehicle{number: "car6"}

	attendantWithFirstEmptyLotPlan, _ := NewAttendantv2(ParkInFirstEmptyLot, lot1, lot2)
	attendantWithLeastFilledLotPlan, _ := NewAttendantv2(ParkInLeastFilledLot, lot1, lot2)
	attendantWithMaximumFilledLotPlan, _ := NewAttendantv2(ParkInMaximumFilledLot, lot1, lot2)

	_, err := attendantWithFirstEmptyLotPlan.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 by attendantWithFirstEmptyLotPlan %v", err)
	}

	_, err = attendantWithFirstEmptyLotPlan.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 by attendantWithFirstEmptyLotPlan %v", err)
	}

	if car1.Equals(lot1.vehicles[0]) == false {
		t.Fatal("attendant should park car1 in lot1")
	}

	if car2.Equals(lot1.vehicles[1]) == false {
		t.Fatal("attendant should park car2 in lot1")
	}

	_, err = attendantWithLeastFilledLotPlan.Park(car3)
	if err != nil {
		t.Fatalf("Park setup failed for car3 by attendantWithLeastFilledLotPlan %v", err)
	}

	_, err = attendantWithFirstEmptyLotPlan.Unpark(car1)
	if err != nil {
		t.Fatalf("unpark setup failed for car1 by attendantWithFirstEmptyLotPlan %v", err)
	}

	_, err = attendantWithMaximumFilledLotPlan.Unpark(car2)
	if err != nil {
		t.Fatalf("unpark setup failed for car2 by attendantWithLeastFilledLotPlan %v", err)
	}

	_, err = attendantWithMaximumFilledLotPlan.Park(car5)
	if err != nil {
		t.Fatalf("park setup failed for car1 by attendantWithMaximumFilledLotPlan %v", err)
	}

	if car5.Equals(lot2.vehicles[1]) == false {
		t.Fatal("car1 should parked in lot2 by attendantWithMaximumFilledLotPlan")
	}

	_, err = attendantWithMaximumFilledLotPlan.Park(car6)
	if err != nil {
		t.Fatalf("park setup failed for car2 by attendantWithLeastFilledLotPlan %v", err)
	}

	if car6.Equals(lot1.vehicles[0]) == false {
		t.Fatal("car1 should parked in lot1 by attendantWithMaximumFilledLotPlan")
	}

	_, err = attendantWithLeastFilledLotPlan.Park(car4)
	if err != nil {
		t.Fatalf("Park setup failed for car4 by attendantWithLeastFilledLotPlan %v", err)
	}

	if car4.Equals(lot1.vehicles[1]) == false {
		t.Fatalf("car4 should have been parked in lot1 by attendantWithLeastFilledLotPlan %v", lot2.vehicles[1])
	}

}

type mockParkingLot struct {
	mock.Mock
}

func (m *mockParkingLot) Park(veh *vehicle) (*vehicle, error) {
	args := m.Called(veh)
	return args.Get(0).(*vehicle), args.Error(1)
}

func (m *mockParkingLot) Unpark(veh *vehicle) (*vehicle, error) {
	args := m.Called(veh)
	return args.Get(0).(*vehicle), args.Error(1)
}

func (m *mockParkingLot) isparked(veh *vehicle) bool {
	args := m.Called(veh)
	return args.Bool(0)
}

func (m *mockParkingLot) parkedVehicleCount() int {
	args := m.Called()
	return args.Int(0)
}

func (m *mockParkingLot) AddSubscriberParkingFull(subscriber ParkingFullReceiver) {
	m.Called(subscriber)
}

func (m *mockParkingLot) AddSubscriberParkingStatus(subscriber ParkingStatusReceiver) {
	m.Called(subscriber)
}

func TestAttendantHavingLeastFilledLotParkingPlanParkVehicleInLeastFilledLot(t *testing.T) {
	//intialization
	mockLot1 := new(mockParkingLot)
	mockLot2 := new(mockParkingLot)

	mockLot1.On("AddSubscriberParkingFull", mock.Anything).Return()
	mockLot2.On("AddSubscriberParkingFull", mock.Anything).Return()

	mockLot1.On("AddSubscriberParkingStatus", mock.Anything).Return()
	mockLot2.On("AddSubscriberParkingStatus", mock.Anything).Return()

	mockLot1.On("isparked", car1).Return(false)
	mockLot1.On("isparked", car2).Return(false)

	mockLot2.On("isparked", car2).Return(false)
	mockLot2.On("isparked", car1).Return(false)

	attendant, _ := NewAttendantv2(ParkInLeastFilledLot, mockLot1, mockLot2)

	//assertions
	mockLot1.On("Park", car1).Return(car1, nil).Once()
	mockLot2.On("Park", car2).Return(car2, nil).Once()
	mockLot1.On("parkedVehicleCount").Return(0).Once()
	mockLot1.On("parkedVehicleCount").Return(1).Once()
	mockLot2.On("parkedVehicleCount").Return(0)

	//business logic
	_, err := attendant.Park(car1)

	assert.NoError(t, err, "attendant should able to park the car1")

	_, err = attendant.Park(car2)

	assert.NoError(t, err, "attendant should able to park the car2")

	// verify assertions
	mockLot1.AssertExpectations(t)
	mockLot2.AssertExpectations(t)

}

func TestAttendantHavingMaximumFilledLotParkingPlanParkInMostFilledLot(t *testing.T) {
	//intialization
	mockLot1 := new(mockParkingLot)
	mockLot2 := new(mockParkingLot)

	mockLot1.On("AddSubscriberParkingFull", mock.Anything).Return()
	mockLot2.On("AddSubscriberParkingFull", mock.Anything).Return()

	mockLot1.On("AddSubscriberParkingStatus", mock.Anything).Return()
	mockLot2.On("AddSubscriberParkingStatus", mock.Anything).Return()

	mockLot1.On("isparked", car1).Return(false)
	mockLot1.On("isparked", car2).Return(false)

	mockLot2.On("isparked", car2).Return(false)
	mockLot2.On("isparked", car1).Return(false)

	attendantWithFirstEmptyLotPlan, _ := NewAttendantv2(ParkInFirstEmptyLot, mockLot1, mockLot2)
	attendantWithMaxFilledLotPlan, _ := NewAttendantv2(ParkInMaximumFilledLot, mockLot1, mockLot2)

	//assertions
	mockLot1.On("Park", car2).Return(car2, nil).Once()
	mockLot1.On("Park", car1).Return(car1, nil).Once()
	mockLot1.On("parkedVehicleCount").Return(1).Once()
	mockLot2.On("parkedVehicleCount").Return(0).Once()

	//business logic
	_, err := attendantWithFirstEmptyLotPlan.Park(car2)

	assert.NoError(t, err, "attendantWithFirstEmptyLotPlan should able to park the car3")

	_, err = attendantWithMaxFilledLotPlan.Park(car1)

	assert.NoError(t, err, "attendantWithMaxFilledLotPlan should able to park the car1")

	// verify assertions
	mockLot1.AssertExpectations(t)
	mockLot2.AssertExpectations(t)

}
