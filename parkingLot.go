package parking

import "errors"

type vehicle struct {
	number string
	lotId  uint
}

type Lot struct {
	capacity uint
	vehicles []*vehicle
}

type Owner struct {
	Name string
	Lots []Lot
}

func NewOwner(name string) (*Owner, error) {
	if name == "" {
		return nil, errors.New("Owner name should not be empty")
	}
	return &Owner{Name: name}, nil
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
			return &vehicle{number: vehicleNumber, lotId: lotId}, nil
		}
		if (*l).vehicles[i].number == vehicleNumber {
			return nil, errors.New("car already parked in parking lot")
		}
	}
	return nil, errors.New("parking lot full")
}

func (l *Lot) AvailabilityNotification() string {
	counter := 0
	for i := 0; i < len(l.vehicles); i++ {
		if l.vehicles[i] != nil && l.vehicles[i].number != "" {
			counter++
		}
	}
	if counter == int(l.capacity) {
		return "parking lot is full"
	}
	return "parking lot is available"
}
