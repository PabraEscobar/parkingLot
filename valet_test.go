package parking

import (
	"errors"
	"testing"
)

func TestNewAttendant(t *testing.T) {
	lot, _ := Newlot(2)
	_, err := NewAttendant(uint(FirstAvailableSlot), lot)
	if err != nil {
		t.Errorf("Attendant should be exist with parking lot")
	}
}

func TestNewAttendantCannotExistWithNilLot(t *testing.T) {
	_, err := NewAttendant(uint(FirstAvailableSlot), nil)
	if err == nil {
		t.Errorf("Attendant should not exist without the parking lot")
	}
}

func TestAttendantParkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)
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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)
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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)
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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)

	attendant.Park(car1)
	_, actualErr := attendant.Park(car1)
	expectedErr := errors.New("car already parked in parking lot")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park the same vehicle")
	}
}

func TestAttendantCannotParkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)
	attendant.Park(car1)
	_, actualErr := attendant.Park(nil)
	expectedErr := errors.New("nil vehicle cannot be parked")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park nil vehicle")
	}
}

func TestAttendantCannotUnparkNonParkedVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)

	attendant.Park(car2)
	_, actualErr := attendant.Unpark(car1)
	expectedErr := errors.New("vehicle not parked in the parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant cannot unpark the nonexistent vehicle")
	}
}

func TestAttendantCannotUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)
	_, actualErr := attendant.Unpark(nil)

	expectedErr := errors.New("nil vehicle cannot be unparked")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to unpark nil vehicle")
	}
}

func TestAttendantParksInFirstAvailableSlot(t *testing.T) {
	lot, _ := Newlot(3)
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)
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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot)

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
	attendant, err := NewAttendant(uint(FirstAvailableSlot), lot, anotherLot)

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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot, anotherLot)

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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot, anotherLot)

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
	attendant, _ := NewAttendant(uint(FirstAvailableSlot), lot, anotherLot)

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

func TestBalancedAttendantParkVehicleInLeastFilledLot(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)

	//business logic
	balancedAttendant, err := NewAttendant(uint(LeastFilledLot), firstLot, secondLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(balancedAttendant.lots) != 2 {
		t.Fatal("attendant should have 2 lots")
	}

	//logic to test
	balancedAttendant.Park(car1)
	_, actualErr := balancedAttendant.Park(car2)

	if actualErr != nil {
		t.Fatal("car2 should be parked")
	}
	if secondLot.vehicles[0] == nil {
		t.Fatal("balanced attendant should park in the second lot")
	}
	if car2.Equals(secondLot.vehicles[0]) == false {
		t.Fatal("balanced attendant should park in the first slot of second lot")
	}
}

func TestBalancedAttendantCannotParkVehicleWhenParkingIsFull(t *testing.T) {
	//intialization
	lot, _ := NewlotV2(0, 1)
	anotherLot, _ := NewlotV2(1, 1)

	//business logic
	balancedAttendant, err := NewAttendant(uint(LeastFilledLot), lot, anotherLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(balancedAttendant.lots) != 2 {
		t.Fatal("attendant should have 2 lots")
	}

	//logic to test
	balancedAttendant.Park(car1)
	balancedAttendant.Park(car2)

	car3, _ := NewVehicle("MH14FG4567")

	expectedErr := errors.New("parking is full")
	_, actualErr := balancedAttendant.Park(car3)

	if actualErr == expectedErr {
		t.Fatal("car3 cannot be parked when parking lot is full")
	}

}

func TestBalancedAttendantParkVehicleInFirstOrderWhenMultipleLotsHaveLeastFilled(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)
	thirdLot, _ := NewlotV2(2, 3)

	//business logic
	balancedAttendant, err := NewAttendant(uint(LeastFilledLot), firstLot, secondLot, thirdLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(balancedAttendant.lots) != 3 {
		t.Fatal("attendant should have 3 lots")
	}

	//logic to test
	balancedAttendant.Park(car1)
	balancedAttendant.Park(car2)

	car3, _ := NewVehicle("MH14FG4567")

	_, actualErr := balancedAttendant.Park(car3)

	if actualErr != nil {
		t.Fatal("car3 cannot be parked when parking lot is full")
	}
	if secondLot.vehicles[0] == nil {
		t.Fatal("car3 should be parked in the second lot as the first order with the least filled slots")
	}
	if car3.Equals(secondLot.vehicles[0]) {
		t.Fatal("car3 should be parked in the first slot of second lot")
	}
}

func TestBothAttendantParkVehicle(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)
	thirdLot, _ := NewlotV2(2, 3)

	//business logic
	simpleAttendant, _ := NewAttendant(uint(FirstAvailableSlot), firstLot, secondLot, thirdLot)
	balancedAttendant, err := NewAttendant(uint(LeastFilledLot), firstLot, secondLot, thirdLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(simpleAttendant.lots) != 3 {
		t.Fatal("simple attendant should have 3 lots")
	}
	if len(balancedAttendant.lots) != 3 {
		t.Fatal("balanced attendant should have 3 lots")
	}

	//logic to test
	simpleAttendant.Park(car1)
	balancedAttendant.Park(car2)

	car3, _ := NewVehicle("MH14FG4567")

	_, actualErr := balancedAttendant.Park(car3)

	if actualErr != nil {
		t.Fatal("car3 cannot be parked when parking lot is full")
	}
	if secondLot.vehicles[0] == nil {
		t.Fatal("car3 should be parked in the second lot as the first order with the least filled slots")
	}

	if car3.Equals(thirdLot.vehicles[0]) == false {
		t.Fatal("car3 should be parked in the first slot of third lot")
	}
}

func TestBothAttendantParkVehicleBasedOnType(t *testing.T) {
	lot1, _ := NewlotV2(0, 2)
	lot2, _ := NewlotV2(1, 2)

	car3 := &vehicle{number: "car3"}
	car4 := &vehicle{number: "car4"}

	simpleAttendant, _ := NewAttendant(uint(FirstAvailableSlot), lot1, lot2)
	complexAttendant, _ := NewAttendant(uint(LeastFilledLot), lot1, lot2)
	_, err := simpleAttendant.Park(car1)
	if err != nil {
		t.Fatalf("park setup failed for car1 %v", err)
	}

	_, err = simpleAttendant.Park(car2)
	if err != nil {
		t.Fatalf("park stup failed for car2 %v", err)
	}

	if car1.Equals(lot1.vehicles[0]) == false {
		t.Fatal("attendant should park in lot1")
	}
	_, err = complexAttendant.Park(car3)
	if err != nil {
		t.Fatal(err)
	}

	_, err = simpleAttendant.Unpark(car1)
	if err != nil {
		t.Fatal(err)
	}

	_, err = simpleAttendant.Unpark(car2)
	if err != nil {
		t.Fatal(err)
	}

	_, err = complexAttendant.Park(car4)

	if err != nil {
		t.Fatal(err)
	}
	if car4.Equals(lot1.vehicles[0]) == false {
		t.Fatalf("car should have been parked in first slot but was parked in lot 2 %v", lot2.vehicles[1])
	}
}

func TestAttendantWithMFLMethodParkVehicleInMostFilledLot(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)

	//business logic
	mostFilledLotParkAttendant, err := NewAttendant(uint(MostFilledLot), firstLot, secondLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(mostFilledLotParkAttendant.lots) != 2 {
		t.Fatal("attendant should have 2 lots")
	}

	//logic to test
	mostFilledLotParkAttendant.Park(car1)
	_, actualErr := mostFilledLotParkAttendant.Park(car2)

	if actualErr != nil {
		t.Fatal("car2 should be parked")
	}
	if firstLot.vehicles[0] == nil {
		t.Fatal("attendant with the most filled lot method should park in the second lot")
	}
	if car2.Equals(firstLot.vehicles[1]) == false {
		t.Fatal("attendant with the most filled lot method should park in the first slot of second lot")
	}
}

func TestAttendantWithMFLMethodCannotParkVehicleWhenParkingIsFull(t *testing.T) {
	//intialization
	lot, _ := NewlotV2(0, 1)
	anotherLot, _ := NewlotV2(1, 1)

	//business logic
	mostFilledLotParkAttendant, err := NewAttendant(uint(MostFilledLot), lot, anotherLot)

	//assertions
	if err != nil {
		t.Fatal("attendant should be created with multiple lots")
	}
	if len(mostFilledLotParkAttendant.lots) != 2 {
		t.Fatal("attendant should have 2 lots")
	}

	//logic to test
	mostFilledLotParkAttendant.Park(car1)
	mostFilledLotParkAttendant.Park(car2)

	car3, _ := NewVehicle("MH14FG4567")

	expectedErr := errors.New("parking is full")
	_, actualErr := mostFilledLotParkAttendant.Park(car3)

	if actualErr == expectedErr {
		t.Fatal("car3 cannot be parked when parking lot is full")
	}
}

func TestAttendantWithMFLMethodParkVehicleInFirstOrderWhenMultipleLotsHaveMostFilled(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)
	thirdLot, _ := NewlotV2(2, 3)

	//business logic
	balancedAttendant, err1 := NewAttendant(uint(LeastFilledLot), firstLot, secondLot, thirdLot)
	mostFilledLotParkAttendant, err2 := NewAttendant(uint(MostFilledLot), firstLot, secondLot, thirdLot)

	//assertions
	if err1 != nil {
		t.Fatal("balanced attendant should be created with multiple lots")
	}
	if err2 != nil {
		t.Fatal("attendant park with most filled lot plan should be created with multiple lots")
	}
	if len(balancedAttendant.lots) != 3 {
		t.Fatal("attendant should have 3 lots")
	}
	if len(mostFilledLotParkAttendant.lots) != 3 {
		t.Fatal("attendant should have 3 lots")
	}

	//logic to test
	balancedAttendant.Park(car1)
	balancedAttendant.Park(car2)

	car3, _ := NewVehicle("MH14FG4567")

	_, actualErr := mostFilledLotParkAttendant.Park(car3)

	if actualErr != nil {
		t.Fatal("car3 should be parked")
	}
	if firstLot.vehicles[1] == nil {
		t.Fatal("car3 should be parked in the first lot as the first order with the most filled slots")
	}
	if car3.Equals(firstLot.vehicles[1]) == false {
		t.Fatal("car3 should be parked in the second slot of first lot")
	}
}

func TestBalancedAttendentUnparkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(LeastFilledLot), lot)
	var err error

	//logic to test
	_, _ = attendant.Park(car1)
	_, err = attendant.Unpark(car1)

	//assertions
	if err != nil {
		t.Fatal("balanced attendent should unpark the vehicle")
	}
}

func TestBalancedAttendantCannotParkWhenParkinFull(t *testing.T) {
	//initalization
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(uint(LeastFilledLot), lot)
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

func TestBalancedAttendantCannotParkVehicleWhichIsAlreadyParked(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(LeastFilledLot), lot)

	attendant.Park(car1)
	_, actualErr := attendant.Park(car1)
	expectedErr := errors.New("car already parked in parking lot")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park the same vehicle")
	}
}

func TestBalancedAttendantCannotParkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(LeastFilledLot), lot)
	attendant.Park(car1)
	_, actualErr := attendant.Park(nil)
	expectedErr := errors.New("nil vehicle cannot be parked")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park nil vehicle")
	}
}

func TestBalancedAttendantCannotUnparkNonParkedVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(LeastFilledLot), lot)

	attendant.Park(car2)
	_, actualErr := attendant.Unpark(car1)
	expectedErr := errors.New("vehicle not parked in the parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant cannot unpark the nonexistent vehicle")
	}
}

func TestBalancedAttendantCannotUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(LeastFilledLot), lot)
	_, actualErr := attendant.Unpark(nil)

	expectedErr := errors.New("nil vehicle cannot be unparked")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to unpark nil vehicle")
	}
}

func TestAttendantWithMFLMethodUnparkVehicle(t *testing.T) {
	//initalization
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(MostFilledLot), lot)
	var err error

	//logic to test
	_, _ = attendant.Park(car1)
	_, err = attendant.Unpark(car1)

	//assertions
	if err != nil {
		t.Fatal("attendent with most filled lot should unpark the vehicle")
	}
}

func TestAttendantWithMFLMethodCannotParkWhenParkinFull(t *testing.T) {
	//initalization
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(uint(MostFilledLot), lot)
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

func TestAttendantWithMFLMethodCannotParkVehicleWhichIsAlreadyParked(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(MostFilledLot), lot)

	attendant.Park(car1)
	_, actualErr := attendant.Park(car1)
	expectedErr := errors.New("car already parked in parking lot")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park the same vehicle")
	}
}

func TestAttendantWithMFLMethodCannotParkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(MostFilledLot), lot)
	attendant.Park(car1)
	_, actualErr := attendant.Park(nil)
	expectedErr := errors.New("nil vehicle cannot be parked")

	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to park nil vehicle")
	}
}

func TestAttendantWithMFLMethodCannotUnparkNonParkedVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(MostFilledLot), lot)

	attendant.Park(car2)
	_, actualErr := attendant.Unpark(car1)
	expectedErr := errors.New("vehicle not parked in the parking lot")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant cannot unpark the nonexistent vehicle")
	}
}

func TestAttendantWithMFLMethodCannotUnparkNilVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant, _ := NewAttendant(uint(MostFilledLot), lot)
	_, actualErr := attendant.Unpark(nil)

	expectedErr := errors.New("nil vehicle cannot be unparked")
	if actualErr.Error() != expectedErr.Error() {
		t.Errorf("attendant should not be able to unpark nil vehicle")
	}
}

func TestAttendantWithMFLMethodUnparkVehicleWhichParkedByOtherTypeAttendant(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)
	thirdLot, _ := NewlotV2(2, 3)

	//business logic
	balancedAttendant, err1 := NewAttendant(uint(LeastFilledLot), firstLot, secondLot, thirdLot)
	mostFilledLotParkAttendant, err2 := NewAttendant(uint(MostFilledLot), firstLot, secondLot, thirdLot)

	//assertions
	if err1 != nil {
		t.Fatal("balanced attendant should be created with multiple lots")
	}
	if err2 != nil {
		t.Fatal("attendant park with most filled lot plan should be created with multiple lots")
	}
	if len(balancedAttendant.lots) != 3 {
		t.Fatal("balanced attendant should have 3 lots")
	}
	if len(mostFilledLotParkAttendant.lots) != 3 {
		t.Fatal("attendant with most filled park lot approach should have 3 lots")
	}

	//logic to test
	balancedAttendant.Park(car1)
	balancedAttendant.Park(car2)

	car3, _ := NewVehicle("MH14FG4567")

	_, actualErr := mostFilledLotParkAttendant.Unpark(car2)

	if actualErr != nil {
		t.Fatal("car2 should be unparked")
	}
	if car3.Equals(firstLot.vehicles[1]) == true {
		t.Fatal("car3 should be unparked from the second slot of first lot")
	}
}

func TestAttendantWithMFLMethodCannotParkAlreadyParkedByOtherTypeAttendant(t *testing.T) {
	//intialization
	firstLot, _ := NewlotV2(0, 3)
	secondLot, _ := NewlotV2(1, 3)
	thirdLot, _ := NewlotV2(2, 3)

	//business logic
	balancedAttendant, err1 := NewAttendant(uint(LeastFilledLot), firstLot, secondLot, thirdLot)
	mostFilledLotParkAttendant, err2 := NewAttendant(uint(MostFilledLot), firstLot, secondLot, thirdLot)

	//assertions
	if err1 != nil {
		t.Fatal("balanced attendant should be created with multiple lots")
	}
	if err2 != nil {
		t.Fatal("attendant park with most filled lot plan should be created with multiple lots")
	}
	if len(balancedAttendant.lots) != 3 {
		t.Fatal("balanced attendant should have 3 lots")
	}
	if len(mostFilledLotParkAttendant.lots) != 3 {
		t.Fatal("attendant with most filled lot approach should have 3 lots")
	}

	//logic to test
	balancedAttendant.Park(car1)
	balancedAttendant.Park(car2)

	_, actualErr := mostFilledLotParkAttendant.Park(car2)

	if actualErr == nil {
		t.Fatal("car2 cannot park which is already being parked by other attendant")
	}
}
