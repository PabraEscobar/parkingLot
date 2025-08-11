package parking

import (
	"errors"
)

type attendant struct {
	lot         *Lot
	parkingFull bool
}

// ParkingFullReceive implements ParkingFullReceiver.
func (a *attendant) ParkingFullReceive() {
	a.parkingFull = true
}

// attendant implements ParkingStatusReceiver
func (a *attendant) ParkingStatusReceive(status ParkingStatus) {
	if status == ParkingAvailable {
		a.parkingFull = false
	}
}

func (a *attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	return a.lot.Unpark(vehicleNumber)
}

func (a *attendant) Park(vehicleNumber string) (*vehicle, error) {
	if a.parkingFull {
		return nil, errors.New("parking is full")
	}
	return a.lot.Park(vehicleNumber)
}

func NewAttendant(lot *Lot) (*attendant, error) {
	if lot == nil {
		return nil, errors.New("attendant does not exist without parking lot")
	}
	a := &attendant{lot: lot}
	lot.SubscribeParkingFullStatus(a)
	lot.subscriberParkingStatus = a
	return a, nil
}
