package parking

import "errors"

type vehicle struct {
	number string
	lotId  uint
}

type ParkingStatus uint

const (
	Unknown ParkingStatus = iota
	ParkingAvailable
	ParkingFull
)

type Lot struct {
	capacity    uint
	vehicles    []*vehicle
	subscribers []ParkingStatusReceiver
}

type ParkingStatusReceiver interface {
	Notify(status ParkingStatus)
}

func (l *Lot) SubscribeParkingStatus(subscriber ParkingStatusReceiver) {
	l.subscribers = append(l.subscribers, subscriber)
}

func (l *Lot) Unpark(vehicleNumber string) (*vehicle, error) {
	if vehicleNumber == "" {
		return nil, errors.New("vehicle number is manadatory to unpark the vehicle")
	}
	var lotId uint
	for i := 0; i < len(l.vehicles); i++ {
		if l.vehicles[i] != nil && l.vehicles[i].number == vehicleNumber {
			lotId = uint(i + 1)
			l.vehicles[i].number = ""
			l.vehicles[i] = nil
			if int(lotId) == len((*l).vehicles) {
				l.availabilityNotification(ParkingAvailable)
			}
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
	}
	return nil, errors.New("vehicle not parked in the parking lot with provided number")
}

func Newlot(capacity uint) (*Lot, error) {
	if capacity == 0 {
		return nil, errors.New("capacity can't be zero")
	}
	l := make([]*vehicle, capacity)
	return &Lot{capacity: capacity, vehicles: l}, nil
}

func (l *Lot) Park(vehicleNumber string) (*vehicle, error) {
	if vehicleNumber == "" {
		return nil, errors.New("vehicle number is mandatory to park")
	}
	var lotId uint
	for i := 0; i < len((*l).vehicles); i++ {
		if (*l).vehicles[i] == nil {
			lotId = uint(i + 1)
			(*l).vehicles[i] = &vehicle{number: vehicleNumber, lotId: uint(i + 1)}
			if int(lotId) == len((*l).vehicles) {
				l.availabilityNotification(ParkingFull)
			}
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
		if (*l).vehicles[i].number == vehicleNumber {
			return nil, errors.New("car already parked in parking lot")
		}
	}
	return nil, errors.New("parking lot is full")
}

func (l *Lot) availabilityNotification(status ParkingStatus) {
	for i := 0; i < len(l.subscribers); i++ {
		if l.subscribers[i] != nil {
			l.subscribers[i].Notify(status)
		}
	}
}
