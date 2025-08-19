package parking

import (
	"errors"
	"math"
)

type findAvailableLot func(a *attendant) (*lot, error)

type ParkingPlan int

const (
	ParkInFirstEmptyLot ParkingPlan = iota
	ParkInLeastFilledLot
)

type attendant struct {
	availableLotForPark findAvailableLot
	approach            ParkingPlan
	lots                []*lot
	lotsFullStatus      []bool
}

// attendant implements ParkingFullReceiver
func (a *attendant) ParkingFullReceive(i uint) {
	a.lotsFullStatus[i] = true
}

func (a *attendant) Receive(status ParkingStatus, i uint) {
	if status == ParkingAvailable {
		a.lotsFullStatus[i] = false
	}
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

	availablelot := a.availableLotForPark
	lot, err := availablelot(a)

	if err != nil {
		return nil, err
	}

	_, err = lot.Park(vehicle)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func findFirstEmptylot(a *attendant) (*lot, error) {
	for i, lot := range a.lots {
		if a.lotsFullStatus[i] {
			continue
		}
		return lot, nil
	}
	return nil, errors.New("parking is full")
}

func findLotWithleastVehicles(a *attendant) (*lot, error) {
	Count := math.MaxInt
	lotId := -1
	for i, lot := range a.lots {
		if a.lotsFullStatus[i] {
			continue
		}
		vehicleCount := lot.vehicleCount()
		if Count > vehicleCount {
			Count = vehicleCount
			lotId = i
		}
	}
	if lotId == -1 {
		return nil, errors.New("parking is full")
	}
	return a.lots[lotId], nil
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
	a := &attendant{lots: l, lotsFullStatus: parkingFull}
	for _, lot := range lots {
		lot.AddSubscriberParkingFull(a)
		lot.AddSubscriberParkingStatus(a)
	}
	return a, nil
}

func NewAttendantv2(plan ParkingPlan, lots ...*lot) (*attendant, error) {
	attendant, err := NewAttendant(lots...)
	if err != nil {
		return nil, err
	}
	attendant.approach = plan
	if plan == ParkInFirstEmptyLot {
		attendant.availableLotForPark = findFirstEmptylot
	} else {
		attendant.availableLotForPark = findLotWithleastVehicles
	}
	return attendant, nil
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
