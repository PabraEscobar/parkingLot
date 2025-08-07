package parking

import "testing"

func TestNewAttendant(t *testing.T) {
	attendant := NewAttendant()
	if attendant == nil {
		t.Errorf("Attendant should be exist")
	}
}

func TestAttendantReceiveNotificationParkingFull(t *testing.T) {
	lot, _ := Newlot(1)
	attendant := NewAttendant()
	attendant.AddParkingLot(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	attendant.lots[0].park(vehicle)
	if attendant.status != ParkingFull {
		t.Errorf("attendant should be notified when parking is full")
	}
}

func TestAttendantCanParkTheVehicleWhenParkingIsNotFull(t *testing.T) {
	lot, _ := Newlot(1)
	attendant := NewAttendant()
	attendant.AddParkingLot(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	_, err := attendant.Park(vehicle)
	if err != nil {
		t.Errorf("attendant can park the vehicle")
	}
}

func TestAttendantCannotParkTheVehicleWhenParkingFull(t *testing.T) {
	lot, _ := Newlot(1)
	attendant := NewAttendant()
	attendant.AddParkingLot(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	anotherVehicle := "KA02FG4567"
	attendant.Park(vehicle)
	_, err := attendant.Park(anotherVehicle)
	if err == nil {
		t.Errorf("attendant cannot park the vehicle when the parking is full")
	}
}

func TestAttendantCanUnparkTheVehicle(t *testing.T) {
	lot, _ := Newlot(2)
	attendant := NewAttendant()
	attendant.AddParkingLot(lot)
	lot.SubscribeParkingFullStatus(attendant)
	vehicle := "KA03FG2345"
	attendant.Park(vehicle)
	_, err := attendant.Unpark(vehicle)
	if err != nil {
		t.Errorf("attendant should unpark the vehicle")
	}
}

func TestAttendentCanManageMultipleLots(t *testing.T) {
	lot, _ := Newlot(1)
	anotherLot, _ := Newlot(2)
	attendant := NewAttendant()
	attendant.AddParkingLot(lot)
	attendant.AddParkingLot(anotherLot)
	if len(attendant.lots) != 2 {
		t.Errorf("Attendant should be able to manage multiple lots")
	}
}
