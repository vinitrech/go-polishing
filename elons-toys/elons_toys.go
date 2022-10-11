package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	return (trackDistance/car.speed)*car.batteryDrain <= car.battery
}
