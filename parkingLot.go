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
	counter := 0
	for i := 0; i < len(l.vehicles); i++ {
		if l.isFreeSlot(i) {
			continue
		}
		counter++
		if l.vehicles[i].Equals(car) {
			l.vehicles[i] = nil
			if counter == int(l.capacity) && l.subscriberParkingStatus != nil {
				l.subscriberParkingStatus.ParkingStatusReceive(ParkingAvailable)
			}
			return &vehicle{number: car.number}, nil
		}
	}
	return nil, errors.New("vehicle not parked in the parking lot with provided number")
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

func (l *lot) Park(vehicle *vehicle) (*vehicle, error) {
	if vehicle == nil {
		return nil, errors.New("vehicle number is mandatory to park")
	}
	var counter int
	for i := 0; i < len((*l).vehicles); i++ {
		if l.isFreeSlot(i) {
			counter = i + 1
			(*l).vehicles[i] = vehicle
			if counter == len((*l).vehicles) {
				for j := 0; j < len(l.subscribersParkingFull); j++ {
					if l.subscribersParkingFull[j] != nil {
						l.subscribersParkingFull[j].ParkingFullReceive()
					}
				}
				if l.subscriberParkingStatus != nil {
					l.subscriberParkingStatus.ParkingStatusReceive(ParkingFull)
				}
			}
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
