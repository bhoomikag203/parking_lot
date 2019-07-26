package main

import (
	"errors"
)

var ErrParkinglotFull = errors.New("Cannot park a vehicle as parking lot is full")

type parkingLot struct {
	capacity int
	occupied int
}

func (p parkingLot) NewParkingLot(capacity int) parkingLot {
	return p
}

func (p *parkingLot) Park(vehicle string) error {
	if p.IsFull() {
		return ErrParkinglotFull
	}
	p.occupied += 1
	return nil
}

func (p *parkingLot) Unpark(vehicle string) {
	p.occupied -= 1
}

func (p *parkingLot) IsFull() bool {
	if p.occupied == p.capacity {
		return true
	}
	return false
}
