package parking

import "errors"

type attendant struct {
	lot ParkingLot
}

func (a *attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	return a.lot.Unpark(vehicleNumber)
}

func (a *attendant) Park(vehicleNumber string) (*vehicle, error) {
	return a.lot.Park(vehicleNumber)
}

type ParkingLot interface {
	Park(vehicleNumber string) (*vehicle, error)
	Unpark(vehicleNumber string) (*vehicle, error)
}

func NewAttendant(lot ParkingLot) (*attendant, error) {
	if lot == nil {
		return nil, errors.New("attendant does not exist without parking lot")
	}
	return &attendant{lot: lot}, nil
}
