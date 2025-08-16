package parking

import (
	"errors"
)

type attendant struct {
	lots        []*lot
	parkingFull []bool
}

// attendant implements ParkingFullReceiver
func (a *attendant) ParkingFullReceive(i uint) {
	a.parkingFull[i] = true
}

func (a *attendant) Unpark(vehicle *vehicle) (*vehicle, error) {
	if a.parkingFull[0] {
		a.parkingFull[0] = false
	}
	return a.lots[0].Unpark(vehicle)
}

func (a *attendant) Park(vehicle *vehicle) (*vehicle, error) {
	if a.isParked(vehicle) {
		return nil, errors.New("car already parked in parking lot")
	}
	for i, lot := range a.lots {
		if a.parkingFull[i] {
			continue
		}
		_, err := lot.Park(vehicle)
		if err != nil {
			return nil, err
		}
		return vehicle, nil
	}
	return nil, errors.New("parking is full")
}

func NewAttendant(lots ...*lot) (*attendant, error) {
	for _, lot := range lots {
		if lot == nil {
			return nil, errors.New("attendant does not exist without parking lot")
		}
	}
	l := make([]*lot, 0, len(lots))
	parkingFull := make([]bool, len(lots)+1)
	l = append(l, lots...)
	a := &attendant{lots: l, parkingFull: parkingFull}
	for _, lot := range lots {
		lot.SubscribeParkingFullStatus(a)
	}
	return a, nil
}

func (a *attendant) isParked(vehicle *vehicle) bool {
	for _, lot := range a.lots {
		flag := lot.Isparked(vehicle)
		if flag {
			return true
		}
	}
	return false
}
