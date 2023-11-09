package main

import "fmt"

// Vehicle interface
type Vehicle interface {
	getNumberOfWheels() int
}

// VehicleFactory interface for creating vehicles
type VehicleFactory interface {
	createVehicle() Vehicle
}

// Bicycle struct implementing Vehicle
type Bicycle struct{}

func (b *Bicycle) getNumberOfWheels() int {
	return 2
}

// BicycleFactory struct implementing VehicleFactory
type BicycleFactory struct{}

func (bf *BicycleFactory) createVehicle() Vehicle {
	return &Bicycle{}
}

// EngineVehicle struct embedding Vehicle
type EngineVehicle struct {
	Vehicle
}

// Car struct embedding EngineVehicle
type Car struct {
	EngineVehicle
}

func (c *Car) getNumberOfWheels() int {
	return 4
}

// CarFactory struct implementing VehicleFactory
type CarFactory struct{}

func (cf *CarFactory) createVehicle() Vehicle {
	return &Car{}
}

// MotorCycle struct embedding EngineVehicle
type MotorCycle struct {
	EngineVehicle
}

func (m *MotorCycle) getNumberOfWheels() int {
	return 2
}

// MotorCycleFactory struct implementing VehicleFactory
type MotorCycleFactory struct{}

func (mf *MotorCycleFactory) createVehicle() Vehicle {
	return &MotorCycle{}
}

func main() {
	// Using factories to create instances of vehicles
	var vehicleList []Vehicle
	factories := []VehicleFactory{&MotorCycleFactory{}, &CarFactory{}, &BicycleFactory{}}

	for _, factory := range factories {
		vehicle := factory.createVehicle()
		vehicleList = append(vehicleList, vehicle)
	}

	// Displaying the number of wheels for each vehicle
	for _, vehicle := range vehicleList {
		fmt.Println(vehicle.getNumberOfWheels())
	}
}
