package main

import (
	"errors"
)

func ParkingLot(capacity int) (int, error) {
	if capacity < 0 {
		return capacity, errors.New("capacity cannot be negative")
	}
	return capacity, nil
}
