package parking

import (
	"errors"
	"math"
)

type attendant struct {
	lots        []*lot
	parkingFull []bool
	count       []uint
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

	for i, lot := range a.lots {
		if !lot.isparked(vehicle) {
			continue
		}
		_, err := lot.Unpark(vehicle)
		if err != nil {
			return nil, err
		}
		if a.count[i] > 0 {
			a.count[i]--
		}
		a.parkingFull[lot.id] = false
	}
	return vehicle, nil
}

func (a *attendant) Park(typeOfAttendant uint, vehicle *vehicle) (*vehicle, error) {
	if vehicle == nil {
		return nil, errors.New("nil vehicle cannot be parked")
	}
	if a.isParked(vehicle) {
		return nil, errors.New("car already parked in parking lot")
	}

	var lot *lot
	var lotId int = -1
	var err error

	if typeOfAttendant == 1 {
		lot, lotId, err = a.findAvailableParkinglot()
	} else {
		lot, lotId, err = a.findLeastParkedVehiclesLot()
	}

	if err != nil {
		return nil, err
	}

	_, err = lot.Park(vehicle)
	if err != nil {
		return nil, err
	}
	a.count[lotId]++

	return vehicle, nil
}

func (a *attendant) findLeastParkedVehiclesLot() (*lot, int, error) {
	minimumCount := math.MaxInt64
	var lotWithLeastFilledSlots *lot
	var lotId int = -1
	for i, lot := range a.lots {
		if !a.parkingFull[i] && a.count[i] < uint(minimumCount) {
			minimumCount = int(a.count[i])
			lotWithLeastFilledSlots = lot
			lotId = i
		}
	}
	if lotWithLeastFilledSlots != nil {
		return lotWithLeastFilledSlots, lotId, nil
	}
	return nil, -1, errors.New("parking is full")
}

func (a *attendant) findAvailableParkinglot() (*lot, int, error) {
	for i, lot := range a.lots {
		if a.parkingFull[i] {
			continue
		}
		return lot, i, nil
	}
	return nil, -1, errors.New("parking is full")
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
	a := &attendant{lots: l, parkingFull: parkingFull, count: make([]uint, len(lots))}
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
