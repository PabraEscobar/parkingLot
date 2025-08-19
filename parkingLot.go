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
	capacity                 uint
	vehicles                 []*vehicle
	subscribersParkingFull   []ParkingFullReceiver
	subscribersParkingStatus []ParkingStatusReceiver
	id                       uint
}

type ParkingStatus uint

const (
	Unknown ParkingStatus = iota
	ParkingAvailable
	ParkingFull
)

type ParkingStatusReceiver interface {
	Receive(status ParkingStatus, id uint)
}

type ParkingFullReceiver interface {
	ParkingFullReceive(i uint)
}

func (l *lot) AddSubscriberParkingFull(subscriber ParkingFullReceiver) {
	l.subscribersParkingFull = append(l.subscribersParkingFull, subscriber)
}

func (l *lot) AddSubscriberParkingStatus(subscriber ParkingStatusReceiver) {
	l.subscribersParkingStatus = append(l.subscribersParkingStatus, subscriber)
}

func (l *lot) Unpark(car *vehicle) (*vehicle, error) {
	if car == nil {
		return nil, errors.New("nil vehicle cannot be unparked")
	}
	for i := 0; i < len(l.vehicles); i++ {
		if l.isFreeSlot(i) {
			continue
		}
		//TODO indentation
		if l.vehicles[i].Equals(car) {
			l.vehicles[i] = nil
			l.notifyParkingAvailable()
			return car, nil
		}
	}
	return nil, errors.New("vehicle not parked in the parking lot")
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
	for _, subscriber := range l.subscribersParkingStatus {
		subscriber.Receive(ParkingAvailable, l.id)
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

func NewlotV2(id, capacity uint) (*lot, error) {
	l, err := Newlot(capacity)
	if err != nil {
		return nil, err
	}
	l.id = id
	return l, nil
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
	for _, subscriber := range l.subscribersParkingStatus {
		subscriber.Receive(ParkingFull, l.id)
	}
	for _, subscriber := range l.subscribersParkingFull {
		subscriber.ParkingFullReceive(l.id)
	}
}

func (l *lot) Park(vehicle *vehicle) (*vehicle, error) {
	if vehicle == nil {
		return nil, errors.New("vehicle number is mandatory to park")
	}

	if l.isparked(vehicle) {
		return nil, errors.New("car already parked in parking lot")
	}

	//find available slot
	for i := 0; i < len(l.vehicles); i++ {
		if l.isFreeSlot(i) {
			l.vehicles[i] = vehicle
			l.notifyParkingFull()
			return vehicle, nil
		}
	}
	return nil, errors.New("parking lot is full")
}

func (l *lot) isparked(vehicle *vehicle) bool {
	for i := 0; i < len(l.vehicles); i++ {
		if vehicle.Equals(l.vehicles[i]) {
			return true
		}
	}
	return false
}

func (l *lot) vehicleCount() int {
	count := 0
	for _, vehicle := range l.vehicles {
		if vehicle != nil {
			count++
		}
	}
	return count
}
