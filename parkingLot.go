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

func Newlot(capacity uint) (*Lot, error) {
	if capacity == 0 {
		return nil, errors.New("capacity can't be zero")
	}
	return &Lot{capacity: capacity, vehicles: nil}, nil
}
