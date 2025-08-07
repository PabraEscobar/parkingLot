package parking

import "testing"

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

func TestAttendantReceiveNotificationParkingFull(t *testing.T) {
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	attendant.lot.Park(vehicle)
	if attendant.status != ParkingFull {
		t.Errorf("attendant should be notified when parking is full")
	}
}

func TestAttendantCanParkTheVehicleWhenParkingIsNotFull(t *testing.T) {
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	_, err := attendant.Park(vehicle)
	if err != nil {
		t.Errorf("attendant can park the vehicle")
	}
}

func TestAttendantCannotParkTheVehicleWhenParkingFull(t *testing.T) {
	lot, _ := Newlot(1)
	attendant, _ := NewAttendant(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	anotherVehicle := "KA02FG4567"
	attendant.Park(vehicle)
	_, err := attendant.Park(anotherVehicle)
	if err == nil {
		t.Errorf("attendant cannot park the vehicle when the parking is full")
	}
}
