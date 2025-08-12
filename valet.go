package parking

import (
	"errors"
)

type Attendant struct {
	lot         *Lot
	parkingFull bool
}

// attendant implements ParkingFullReceiver
func (a *Attendant) ParkingFullReceive() {
	a.parkingFull = true
}

func (a *Attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	if a.parkingFull{
		a.parkingFull=false
	}
	return a.lot.Unpark(vehicleNumber)
}

func (a *Attendant) Park(vehicleNumber string) (*vehicle, error) {
	if a.parkingFull {
		return nil, errors.New("parking is full")
	}
	return a.lot.Park(vehicleNumber)
}

func NewAttendant(lot *Lot) (*Attendant, error) {
	if lot == nil {
		return nil, errors.New("attendant does not exist without parking lot")
	}
	a := &Attendant{lot: lot}
	lot.SubscribeParkingFullStatus(a)
	return a, nil
}
