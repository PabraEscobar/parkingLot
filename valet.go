package parking

import (
	"errors"
	"math"
)

type attendant struct {
//TODO reveal intention
	id          int
	lots        []*lot
	//TODO reveal intention lots is plural and has array but here it is singular but array, fix is not a suffix s but come up with better name
	parkingFull []bool
}

// attendant implements ParkingFullReceiver
func (a *attendant) ParkingFullReceive(i uint) {
	a.parkingFull[i] = true
}

func (a *attendant) Receive(status ParkingStatus, i uint) {
	if status == ParkingAvailable {
		a.parkingFull[i] = false
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

	var lot *lot
	var err error

	//TODO remove both if condition checks on every park call
	if a.id == 1 {
		lot, err = a.firstEmptylot()
	}

	if a.id == 2 {
		lot, err = a.lotWithleastVehicles()
	}

	if err != nil {
		return nil, err
	}

	_, err = lot.Park(vehicle)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (a *attendant) firstEmptylot() (*lot, error) {
	for i, lot := range a.lots {
		if a.parkingFull[i] {
			continue
		}
		return lot, nil
	}
	return nil, errors.New("parking is full")
}

func (a *attendant) lotWithleastVehicles() (*lot, error) {
	Count := math.MaxInt
	lotId := -1
	for i, lot := range a.lots {
		if a.parkingFull[i] {
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
	a := &attendant{lots: l, parkingFull: parkingFull}
	for _, lot := range lots {
		lot.AddSubscriberParkingFull(a)
		lot.AddSubscriberParkingStatus(a)
	}
	return a, nil
}

func NewAttendantv2(choice int, lots ...*lot) (*attendant, error) {
	attendant, err := NewAttendant(lots...)
	if err != nil {
		return nil, err
	}
	if choice != 1 && choice != 2 {
		return nil, errors.New("invalid choice for attendant creation")
	}
	attendant.id = choice
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
