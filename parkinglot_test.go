package main

import (
	"testing"
)

func TestNewParkingLot(t *testing.T) {
	got := parkingLot{capacity: 10, occupied: 0}
	want := parkingLot{10, 0}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestPark(t *testing.T) {
	t.Run("to park one vehicle", func(t *testing.T) {
		pl := parkingLot{capacity: 10, occupied: 0}

		pl.Park("vehicle1")

		if pl.occupied != 1 {
			t.Errorf("Couldn't park a vehicle. got %v want %v", pl.occupied, 1)
		}
	})

	/*t.Run("should not park a vehicle w", func(t *testing.T) {
		pl := parkingLot{capacity: 10, occupied: 0}

		pl.Park("vehicle1")

		if pl.occupied != 1 {
			t.Errorf("Couldn't park a vehicle. got %v want %v", pl.occupied, 1)
		}
	})*/
}
