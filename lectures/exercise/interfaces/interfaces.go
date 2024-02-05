//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int
type Truck string
type Car string
type Motorcycle string
type model string

type LiftHandler interface {
	HandleLift() Lift
}

func (t Truck) HandleLift() Lift {
	return LargeLift
}
func (c Car) HandleLift() Lift {
	return StandardLift
}
func (m Motorcycle) HandleLift() Lift {
	return SmallLift
}

func handleLifts(vehicles []LiftHandler) {
	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		lift := vehicle.HandleLift()
		switch lift {
		case SmallLift:
			fmt.Println("Motorcycle on small lift")
		case StandardLift:
			fmt.Println("Car on standard lift")
		case LargeLift:
			fmt.Println("Truck on large lift")
		}
	}
}

func (t Truck) String() {
	fmt.Sprintf("Truck: %V/ ", string(t))
}
func (c Car) String() {
	fmt.Sprintf("Car: %V/ ", string(c))
}

func (m Motorcycle) String() {
	fmt.Sprintf("Motorcycle: %V/ ", string(m))
}

func main() {
	car := Car("Toyota")
	truck := Truck("Ford")
	motorcycle := Motorcycle("Harley")

	vehicles := []LiftHandler{car, truck, motorcycle}
	handleLifts(vehicles)

	car.String()
	truck.String()
	motorcycle.String()

}
