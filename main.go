package main

import (
// "fmt"
)

func main() {
	p := &parkingLot{capacity: 10, occupied: 0}
	p.Park("vehicle1")
	p.Park("vehicle2")
}
