package main

import (
	"testing"
)

func TestParkingLot(t *testing.T) {
	t.Run("parking lot with given capacity", func(t *testing.T) {
		capacity, _ := ParkingLot(2)
		want := 2

		if capacity != want {
			t.Errorf("capacity is %d but want %d", capacity, want)
		}

	})

	t.Run("raise an error when capacity is negative", func(t *testing.T) {
		_, err := ParkingLot(-1)
		want := "capacity cannot be negative"

		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}

	})

}
