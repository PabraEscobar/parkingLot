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

// attendant implements ParkingStatusReceiver
func (a *Attendant) ParkingStatusReceive(status ParkingStatus) {
	if status == ParkingAvailable {
		a.parkingFull = false
	}
}

func (a *Attendant) Unpark(vehicleNumber string) (*vehicle, error) {
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
	lot.subscriberParkingStatus = a
	return a, nil
}
