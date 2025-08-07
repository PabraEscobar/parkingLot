package parking

import "errors"

type attendant struct {
	lot ParkingLot
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
