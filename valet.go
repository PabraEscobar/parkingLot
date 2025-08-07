package parking

import "errors"

type attendant struct {
	status ParkingStatus
	lots   []ParkingLot
}

type ParkingLot interface {
	Park(vehicleNumber string) (*vehicle, error)
	Unpark(vehicleNumber string) (*vehicle, error)
}

func (a *attendant) AddParkingLot(lot ParkingLot) {
	a.lots = append(a.lots, lot)
}

func NewAttendant() *attendant {
	return &attendant{}
}

func (a *attendant) ParkingFullReceive() {
	a.status = ParkingFull
}

func (a *attendant) Park(vehicleNumber string) (*vehicle, error) {
	if a.status != ParkingFull {
		return a.lots[0].Park(vehicleNumber)
	}
	return nil, errors.New("attendant cannot park the vehicle , parking lot is full")
}

func (a *attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	return a.lots[0].Unpark(vehicleNumber)
}
