package parking

import "errors"

type attendant struct {
	status ParkingStatus
	lot    ParkingLot
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

func (a *attendant) ParkingFullReceive() {
	a.status = ParkingFull
}

func (a *attendant) Park(vehicleNumber string) (*vehicle, error) {
	if a.status != ParkingFull {
		return a.lot.Park(vehicleNumber)
	}
	return nil, errors.New("attendant cannot park the vehicle , parking lot is full")
}

func (a *attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	return a.lot.Unpark(vehicleNumber)
}
