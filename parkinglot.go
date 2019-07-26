package main

type parkingLot struct {
	capacity int
	occupied int
}

func (p parkingLot) NewParkingLot(capacity int) parkingLot {
	return p
}

func (p *parkingLot) Park(vehicle string) {
	p.occupied += 1
}
