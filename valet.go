package parking

import (
	"errors"
	"math"
)

type parkingLot interface {
	Park(vehicle *vehicle) (*vehicle, error)
	Unpark(car *vehicle) (*vehicle, error)
	isparked(vehicle *vehicle) bool
	parkedVehicleCount() int
	AddSubscriberParkingFull(subscriber ParkingFullReceiver)
	AddSubscriberParkingStatus(subscriber ParkingStatusReceiver)
}

type findAvailableLot func(a *attendant) (parkingLot, error)

type ParkingPlan int

const (
	ParkInFirstEmptyLot ParkingPlan = iota
	ParkInLeastFilledLot
	ParkInMaximumFilledLot
)

var parkingPlanMap = map[ParkingPlan]findAvailableLot{
	ParkInFirstEmptyLot:    findFirstEmptylot,
	ParkInLeastFilledLot:   findLotWithleastVehicles,
	ParkInMaximumFilledLot: findLotWithMaximumVehicles,
}

type attendant struct {
	findAvailableLotFn findAvailableLot
	approach           ParkingPlan
	lots               []parkingLot
	lotsFullStatus     []bool
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

	availablelot := a.findAvailableLotFn
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

func findFirstEmptylot(a *attendant) (parkingLot, error) {
	for i, lot := range a.lots {
		if a.lotsFullStatus[i] {
			continue
		}
		return lot, nil
	}
	return nil, errors.New("parking is full")
}

func findLotWithleastVehicles(a *attendant) (parkingLot, error) {
	Count := math.MaxInt
	lotId := -1
	for i, lot := range a.lots {
		if a.lotsFullStatus[i] {
			continue
		}
		vehicleCount := lot.parkedVehicleCount()
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

func findLotWithMaximumVehicles(a *attendant) (parkingLot, error) {
	Count := -1
	lotId := -1
	for i, lot := range a.lots {
		if a.lotsFullStatus[i] {
			continue
		}
		vehicleCount := lot.parkedVehicleCount()
		if Count < vehicleCount {
			Count = vehicleCount
			lotId = i
		}
	}
	if lotId == -1 {
		return nil, errors.New("parking is full")
	}
	return a.lots[lotId], nil
}

func NewAttendant(lots ...parkingLot) (*attendant, error) {
	for _, lot := range lots {
		if lot == nil {
			return nil, errors.New("attendant does not exist without parking lot")
		}
	}
	l := make([]parkingLot, 0, len(lots))
	parkingFull := make([]bool, len(lots)+1)
	l = append(l, lots...)
	a := &attendant{lots: l, lotsFullStatus: parkingFull}
	a.findAvailableLotFn = findFirstEmptylot
	for _, lot := range lots {
		lot.AddSubscriberParkingFull(a)
		lot.AddSubscriberParkingStatus(a)
	}
	return a, nil
}

func NewAttendantv2(plan ParkingPlan, lots ...parkingLot) (*attendant, error) {
	attendant, err := NewAttendant(lots...)
	if err != nil {
		return nil, err
	}

	attendant.approach = plan
	findAvailableLotFn, ok := parkingPlanMap[plan]
	if !ok {
		return nil, errors.New("invalid parking plan")
	}
	attendant.findAvailableLotFn = findAvailableLotFn

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
