package parking

import "errors"

type vehicle struct {
	number string
}

func (v *vehicle) Equals(vehicleTwo *vehicle) bool {
	if vehicleTwo == nil {
		return false
	}
	if v.number == vehicleTwo.number {
		return true
	}
	return false
}

type lot struct {
	capacity                uint
	vehicles                []*vehicle
	subscribersParkingFull  []ParkingFullReceiver
	subscriberParkingStatus ParkingStatusReceiver
}
type ParkingStatus uint

const (
	Unknown ParkingStatus = iota
	ParkingAvailable
	ParkingFull
)

type ParkingStatusReceiver interface {
	ParkingStatusReceive(status ParkingStatus)
}

type ParkingFullReceiver interface {
	ParkingFullReceive()
}

func (l *lot) SubscribeParkingFullStatus(subscriber ParkingFullReceiver) {
	l.subscribersParkingFull = append(l.subscribersParkingFull, subscriber)
}

func (l *lot) Unpark(car *vehicle) (*vehicle, error) {
	if car == nil {
		return nil, errors.New("vehicle number is manadatory to unpark the vehicle")
	}
	for i := 0; i < len(l.vehicles); i++ {
		if l.isFreeSlot(i) {
			continue
		}
		if l.vehicles[i].Equals(car) {
			l.vehicles[i] = nil
			l.notifyParkingAvailable()
			return &vehicle{number: car.number}, nil
		}
	}
	return nil, errors.New("vehicle not parked in the parking lot with provided number")
}

func (l *lot) notifyParkingAvailable() {
	vehicleCount := 0
	for i := 0; i < len(l.vehicles); i++ {
		if l.vehicles[i] != nil {
			vehicleCount++
		}
	}
	if uint(vehicleCount+1) != l.capacity {
		return
	}
	if l.subscriberParkingStatus != nil {
		l.subscriberParkingStatus.ParkingStatusReceive(ParkingAvailable)
	}
}

func (l *lot) isFreeSlot(i int) bool {
	return l.vehicles[i] == nil
}

func Newlot(capacity uint) (*lot, error) {
	if capacity == 0 {
		return nil, errors.New("capacity can't be zero")
	}
	l := make([]*vehicle, capacity)
	return &lot{capacity: capacity, vehicles: l}, nil
}

func (l *lot) notifyParkingFull() {
	vehicleCount := 0
	for i := 0; i < len(l.vehicles); i++ {
		if l.vehicles[i] != nil {
			vehicleCount++
		}
	}
	if uint(vehicleCount) != l.capacity {
		return
	}
	if l.subscriberParkingStatus != nil {
		l.subscriberParkingStatus.ParkingStatusReceive(ParkingFull)
	}
	if len(l.subscribersParkingFull) > 0 {
		for j := 0; j < len(l.subscribersParkingFull); j++ {
			l.subscribersParkingFull[j].ParkingFullReceive()
		}
	}
}

func (l *lot) Park(vehicle *vehicle) (*vehicle, error) {
	if vehicle == nil {
		return nil, errors.New("vehicle number is mandatory to park")
	}
	for i := 0; i < len((*l).vehicles); i++ {
		if l.isFreeSlot(i) {
			(*l).vehicles[i] = vehicle
			l.notifyParkingFull()
			return vehicle, nil
		}
		if !l.isFreeSlot(i) {
			if (*l).vehicles[i].Equals(vehicle) {
				return nil, errors.New("car already parked in parking lot")
			}
		}
	}
	return nil, errors.New("parking lot is full")
}
