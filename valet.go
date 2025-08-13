package parking

import (
	"errors"
)

type attendant struct {
	lot         *lot
	parkingFull bool
}

// attendant implements ParkingFullReceiver
func (a *attendant) ParkingFullReceive() {
	a.parkingFull = true
}

func (a *attendant) Unpark(vehicle *vehicle) (*vehicle, error) {
	if a.parkingFull {
		a.parkingFull = false
	}
	return a.lot.Unpark(vehicle)
}

func (a *attendant) Park(vehicle *vehicle) (*vehicle, error) {
	if a.parkingFull {
		return nil, errors.New("parking is full")
	}
	return a.lot.Park(vehicle)
}

func NewAttendant(lot *lot) (*attendant, error) {
	if lot == nil {
		return nil, errors.New("attendant does not exist without parking lot")
	}
	a := &attendant{lot: lot}
	lot.SubscribeParkingFullStatus(a)
	return a, nil
}
