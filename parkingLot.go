package parking

import "errors"

type vehicle struct {
	number string
	lotId  uint
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

func (l *lot) unpark(vehicleNumber string) (*vehicle, error) {
	if vehicleNumber == "" {
		return nil, errors.New("vehicle number is manadatory to unpark the vehicle")
	}
	counter := 0
	for i := 0; i < len(l.vehicles); i++ {
		if l.vehicles[i] != nil {
			counter++
		}
	}
	var lotId uint
	for i := 0; i < len(l.vehicles); i++ {
		if l.vehicles[i] != nil && l.vehicles[i].number == vehicleNumber {
			lotId = uint(i + 1)
			l.vehicles[i].number = ""
			l.vehicles[i] = nil
			if counter == int(l.capacity) {
				l.subscriberParkingStatus.ParkingStatusReceive(ParkingAvailable)
			}
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
	}
	return nil, errors.New("vehicle not parked in the parking lot with provided number")
}

func Newlot(capacity uint) (*lot, error) {
	if capacity == 0 {
		return nil, errors.New("capacity can't be zero")
	}
	l := make([]*vehicle, capacity)
	return &lot{capacity: capacity, vehicles: l}, nil
}

func (l *lot) park(vehicleNumber string) (*vehicle, error) {
	if vehicleNumber == "" {
		return nil, errors.New("vehicle number is mandatory to park")
	}
	var lotId uint
	for i := 0; i < len((*l).vehicles); i++ {
		if (*l).vehicles[i] == nil {
			lotId = uint(i + 1)
			(*l).vehicles[i] = &vehicle{number: vehicleNumber, lotId: uint(i + 1)}
			if int(lotId) == len((*l).vehicles) {
				for j := 0; j < len(l.subscribersParkingFull); j++ {
					if l.subscribersParkingFull[j] != nil {
						l.subscribersParkingFull[j].ParkingFullReceive()
					}
				}
				if l.subscriberParkingStatus != nil {
					l.subscriberParkingStatus.ParkingStatusReceive(ParkingFull)
				}
			}
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
		if (*l).vehicles[i].number == vehicleNumber {
			return nil, errors.New("car already parked in parking lot")
		}
	}
	return nil, errors.New("parking lot is full")
}
