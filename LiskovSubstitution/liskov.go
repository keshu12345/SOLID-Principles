package main

import "fmt"

type Vehicle interface {
	getNumberOfWheels() int
}

type Bicycle struct {
	Vehicle
}

type EngineVehicle struct {
	Vehicle
}

func (ev *EngineVehicle) hasEngine() bool {
	return true
}

type Car struct {
	EngineVehicle
}

func (c *Car) getNumberOfWheels() int {
	return 4
}
func (c *MotorCycle) getNumberOfWheels() int {
	return 2
}
func (c *Bicycle) getNumberOfWheels() int {
	return 2
}

type MotorCycle struct {
	EngineVehicle
}

func main() {

	var vlist Vehicle
	var vehicleList []Vehicle
	vlist = &MotorCycle{}

	vehicleList = append(vehicleList, vlist)
	vlist = &Car{}
	vehicleList = append(vehicleList, vlist)
	vlist = &Bicycle{}
	vehicleList = append(vehicleList, vlist)
	for _, vehicle := range vehicleList {
		fmt.Println(vehicle.getNumberOfWheels())
	}
}
