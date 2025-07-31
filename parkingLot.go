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
	l := make([]*vehicle, capacity)
	return &Lot{capacity: capacity, vehicles: l}, nil
}

func (l *Lot) Park(vehicleNumber string) (*vehicle, error) {
	if vehicleNumber == "" {
		return nil, errors.New("vehicle number is mandatory to park")
	}
	var lotId uint
	for i := 0; i < len((*l).vehicles); i++ {
		if (*l).vehicles[i] == nil {
			lotId = uint(i + 1)
			(*l).vehicles[i] = &vehicle{number: vehicleNumber, lotId: uint(i + 1)}
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
		if (*l).vehicles[i].number == "" {
			lotId = uint(i + 1)
			(*l).vehicles[i] = &vehicle{number: vehicleNumber, lotId: uint(i + 1)}
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
	}
	return nil, errors.New("parking lot full")
}
