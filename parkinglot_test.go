package main

import (
	"testing"
)

func TestParkingLot(t *testing.T) {
	capacity := ParkingLot(2)
	want := 2

	if capacity != want {
		t.Errorf("capacity is %d but want %d", capacity, want)
	}

}
