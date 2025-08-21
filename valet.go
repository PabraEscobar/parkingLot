package parking

import (
	"errors"
)

type attendant struct {
	lots        []*lot
	parkingFull []bool
}

// attendant implements ParkingFullReceiver
func (a *attendant) ParkingFullReceive(i uint) {
	a.parkingFull[i] = true
}

func (a *attendant) Unpark(vehicle *vehicle) (*vehicle, error) {
	if vehicle == nil {
		return nil, errors.New("nil vehicle cannot be unparked")
	}
	if !a.isParked(vehicle) {
		return nil, errors.New("vehicle not parked in the parking lot")
	}

	for _, lot := range a.lots {
		if !lot.isparked(vehicle) {
			continue
		}
		_, err := lot.Unpark(vehicle)
		if err != nil {
			return nil, err
		}
		a.parkingFull[lot.id] = false
	}
	return vehicle, nil
}

func (a *attendant) Park(vehicle *vehicle) (*vehicle, error) {
	if vehicle == nil {
		return nil, errors.New("nil vehicle cannot be parked")
	}
	if a.isParked(vehicle) {
		return nil, errors.New("car already parked in parking lot")
	}

	lot, err := a.findAvailableParkinglot()
	if err != nil {
		return nil, err
	}

	_, err = lot.Park(vehicle)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (a *attendant) findAvailableParkinglot() (*lot, error) {
	for i, lot := range a.lots {
		if a.parkingFull[i] {
			continue
		}
		return lot, nil
	}
	return nil, errors.New("parking is full")
}

func NewAttendant(lots ...*lot) (*attendant, error) {
	for _, lot := range lots {
		if lot == nil {
			return nil, errors.New("attendant does not exist without parking lot")
		}
	}
	l := make([]*lot, 0, len(lots))
	parkingFull := make([]bool, len(lots)+1)
	l = append(l, lots...)
	a := &attendant{lots: l, parkingFull: parkingFull}
	for _, lot := range lots {
		lot.AddSubscriberParkingFull(a)
	}
	return a, nil
}

func (a *attendant) isParked(vehicle *vehicle) bool {
	for _, lot := range a.lots {
		isParked := lot.isparked(vehicle)
		if isParked {
			return true
		}
	}
	return false
}
