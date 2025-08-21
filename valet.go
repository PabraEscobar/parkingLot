package parking

import (
	"errors"
	"math"
)

type method uint

const (
	UnknownMethod method = iota
	FirstAvailableSlot
	LeastFilledLot
	MostFilledLot
)

type attendant struct {
	lots        []*lot
	parkingFull []bool
	parkingMethod method
	parkingStrategyFunction parkingStrategyFunction
}

type parkingStrategyFunction func(a *attendant) (*lot, int, error)

var parkingStrategyMap = map[method]parkingStrategyFunction{
    FirstAvailableSlot: (*attendant).findFirstAvailableLot,
    LeastFilledLot:     (*attendant).findLeastFilledLot,
    MostFilledLot:      (*attendant).findMostFilledLot,
}

// attendant implements ParkingStatusReceiver
func (a *attendant) Receive(id uint, status ParkingStatus) {
	if status == ParkingAvailable {
		a.parkingFull[id] = false
	} else {
		a.parkingFull[id] = true
	}
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
		if a.lots[i].count > 0 {
			a.lots[i].count--
		}
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

	var lot *lot
	var lotId int = -1
	var err error

	switch a.parkingMethod {
	case FirstAvailableSlot:
		lot, lotId, err = a.findFirstAvailableLot()
	case LeastFilledLot:
		lot, lotId, err = a.findLeastFilledLot()
	case MostFilledLot:
		lot, lotId, err = a.findMostFilledLot()
	}

	if err != nil {
		return nil, err
	}

	_, err = lot.Park(vehicle)
	if err != nil {
		return nil, err
	}
	a.lots[lotId].count++

	return vehicle, nil
}

func (a *attendant) findFirstAvailableLot() (*lot, int, error) {
	for i, lot := range a.lots {
		if a.parkingFull[i] {
			continue
		}
		return lot, i, nil
	}
	return nil, -1, errors.New("parking is full")
}

func (a *attendant) findLeastFilledLot() (*lot, int, error) {
	minimumCount := math.MaxInt64
	var lotWithLeastFilledSlots *lot
	var lotId int = -1
	for i, lot := range a.lots {
		if !a.parkingFull[i] && a.lots[i].count < uint(minimumCount) {
			minimumCount = int(a.lots[i].count)
			lotWithLeastFilledSlots = lot
			lotId = i
		}
	}
	if lotWithLeastFilledSlots != nil {
		return lotWithLeastFilledSlots, lotId, nil
	}
	return nil, -1, errors.New("parking is full")
}

func (a *attendant) findMostFilledLot() (*lot, int, error) {
	maximumCount := -1
	var lotWithMostFilledSlots *lot
	var lotId int = -1
	for i, lot := range a.lots {
		if !a.parkingFull[i] && int(a.lots[i].count) > maximumCount {
			maximumCount = int(a.lots[i].count)
			lotWithMostFilledSlots = lot
			lotId = i
		}
	}
	if lotWithMostFilledSlots != nil {
		return lotWithMostFilledSlots, lotId, nil
	}
	return nil, -1, errors.New("parking is full")
}

func NewAttendant(parkingMethod uint, lots ...*lot) (*attendant, error) {
	for _, lot := range lots {
		if lot == nil {
			return nil, errors.New("attendant does not exist without parking lot")
		}
	}
	l := make([]*lot, 0, len(lots))
	parkingFull := make([]bool, len(lots)+1)
	l = append(l, lots...)
	a := &attendant{lots: l, parkingFull: parkingFull, parkingMethod: method(parkingMethod)}

	parkingStrategyFn, ok := parkingStrategyMap[a.parkingMethod]

	if ok {
		a.parkingStrategyFunction = parkingStrategyFn
	} else {
		return nil, errors.New("invalid parking method")
	}

	for _, lot := range lots {
		lot.AddSubscriberParkingStatus(a)
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
