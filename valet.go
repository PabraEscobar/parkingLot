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
	for lot := 0; lot < len(a.lots); lot++ {
		if a.status != ParkingFull {
			v, err := a.lots[lot].park(vehicleNumber)
			if err != nil {
				if errors.Is(err, ErrParkingLotFull) {
					continue
				} else {
					return nil, errors.New("unable to park the vehicle")
				}
			} else {
				return v, nil
			}
		} else {
			a.status = ParkingAvailable
			continue
		}
	}
	return nil, errors.New("attendant cannot park the vehicle , parking lot is full")
}

func (a *Attendant) Unpark(vehicleNumber string) (*vehicle, error) {
	return a.lots[0].unpark(vehicleNumber)
}
