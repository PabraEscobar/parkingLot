package parking

import "errors"

type vehicle struct {
	number string
	lotId  uint
}

type Lot struct {
	capacity uint
	vehicles []*vehicle
}

func (l *Lot) IsSlotAvailable() (bool, error) {
	if len(l.vehicles) == int(l.capacity) {
		return false, errors.New("no slot is available")
	}
	return true, nil
}

func Newlot(capacity uint) (*Lot, error) {
	if capacity == 0 {
		return nil, errors.New("capacity can't be zero")
	}
	return &Lot{capacity: capacity, vehicles: nil}, nil
}

func (l *Lot) Park(vehicleNumber string) (bool, error) {
	if vehicleNumber == "" {
		return false, errors.New("vehicle number is mandatory to park")
	}
	return true, nil
}
