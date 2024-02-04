package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true

}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true

}
func (lot *ParkingLot) vaccateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false

}
func main() {

	lot := ParkingLot{spaces: make([]Space, 4)}
	fmt.Println(lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupying spaces 1 and 2", lot)

	lot.vaccateSpace(2)
	fmt.Println("After vaccating space 2", lot)
}
