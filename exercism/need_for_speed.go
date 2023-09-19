package speed

type Car struct {
    battery, batteryDrain, speed, distance int
}

type Track struct{
	distance int
}


// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car {100, batteryDrain, speed, 0,}
    return car
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track { distance: distance, }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain {
        car.battery -= car.batteryDrain
    	car.distance += car.speed
    }

    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    batteryNeeded := track.distance / car.speed * car.batteryDrain
    
    return batteryNeeded <= car.battery
}

