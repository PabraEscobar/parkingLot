package parking

import "errors"

type Attendant struct {
	status ParkingStatus
	lots   []parkingLot
}

type parkingLot interface {
	park(vehicleNumber string) (*vehicle, error)
	unpark(vehicleNumber string) (*vehicle, error)
}

func (a *Attendant) AddParkingLot(lot parkingLot) {
	a.lots = append(a.lots, lot)
}

func NewAttendant() *Attendant {
	return &Attendant{}
}

func (a *Attendant) ParkingFullReceive() {
	a.status = ParkingFull
}

func (a *Attendant) Park(vehicleNumber string) (*vehicle, error) {
	if a.status != ParkingFull {
		return a.lots[0].park(vehicleNumber)
	}
	return nil, errors.New("attendant cannot park the vehicle , parking lot is full")
}

func (a *Attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	return a.lots[0].unpark(vehicleNumber)
}
