package parking

import "errors"

type vehicle struct {
	number string
}

func NewVehicle(number string) (*vehicle, error) {
	if number == "" {
		return nil, errors.New("vehicle number is mandatory")
	}
	return &vehicle{number: number}, nil
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
	slots                   []*vehicle
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
	for i := 0; i < len(l.slots); i++ {
		if l.isFreeSlot(i) {
			continue
		}
		if l.slots[i].Equals(car) {
			l.slots[i] = nil
			l.notifyParkingAvailable()
			return car, nil
		}
	}
	return nil, errors.New("vehicle not parked in the parking lot with provided number")
}

func (l *lot) notifyParkingAvailable() {
	vehicleCount := 0
	for i := 0; i < len(l.slots); i++ {
		if l.slots[i] != nil {
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
	return l.slots[i] == nil
}

func Newlot(capacity uint) (*lot, error) {
	if capacity == 0 {
		return nil, errors.New("capacity can't be zero")
	}
	l := make([]*vehicle, capacity)
	return &lot{capacity: capacity, slots: l}, nil
}

func (l *lot) notifyParkingFull() {
	vehicleCount := 0
	for i := 0; i < len(l.slots); i++ {
		if l.slots[i] != nil {
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
	if l.Isparked(vehicle) {
		return nil, errors.New("car already parked in parking lot")
	}

	//find available slot
	for i := 0; i < len(l.slots); i++ {
		if l.isFreeSlot(i) {
			l.slots[i] = vehicle
			l.notifyParkingFull()
			return vehicle, nil
		}
	}
	return nil, errors.New("parking lot is full")
}

func (l *lot) Isparked(vehicle *vehicle) bool {
	for i := 0; i < len(l.slots); i++ {
		if vehicle.Equals(l.slots[i]) {
			return true
		}
	}
	return false
}
