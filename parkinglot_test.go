package main

import (
	"testing"
)

func TestNewParkingLot(t *testing.T) {
	t.Run("create a parking lot", func(t *testing.T) {
		got := parkingLot{capacity: 10, occupied: 0}
		want := parkingLot{10, 0}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("to check if the parking lot is full", func(t *testing.T) {
		pl := parkingLot{capacity: 1, occupied: 1}
		want := true

		got := pl.IsFull()

		if got != want {
			t.Errorf("parking lot is not full. Got %v, want %v", got, want)
		}
	})

}

func TestPark(t *testing.T) {
	t.Run("to park one vehicle", func(t *testing.T) {
		pl := parkingLot{capacity: 1, occupied: 0}

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

func TestUnpark(t *testing.T) {

	t.Run("to be able to unpark a vehicle", func(t *testing.T) {
		pl := parkingLot{capacity: 1, occupied: 0}
		pl.Park("vehicle1")

		pl.Unpark("vehicle1")

		if pl.occupied != 0 {
			t.Errorf("could not unpark")
		}
	})

}
