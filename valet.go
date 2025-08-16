package parking

import (
	"errors"
)

type attendant struct {
	lots        []*lot
	parkingFull bool
}

// attendant implements ParkingFullReceiver
func (a *attendant) ParkingFullReceive(i uint) {
	a.parkingFull = true
}

func (a *attendant) Unpark(vehicle *vehicle) (*vehicle, error) {
	if a.parkingFull {
		a.parkingFull = false
	}
	return a.lots[0].Unpark(vehicle)
}

func (a *attendant) Park(vehicle *vehicle) (*vehicle, error) {
	if a.parkingFull {
		return nil, errors.New("parking is full")
	}
	return a.lots[0].Park(vehicle)
}

func NewAttendant(lots ...*lot) (*attendant, error) {
	for _, lot := range lots {
		if lot == nil {
			return nil, errors.New("attendant does not exist without parking lot")
		}
	}
	l := make([]*lot, 0, len(lots))
	l = append(l, lots...)
	a := &attendant{lots: l}
	for _, lot := range lots {
		lot.SubscribeParkingFullStatus(a)
	}
	return a, nil
}
