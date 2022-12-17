package carslib

import "errors"

type Car interface {
	GetType() string
	GetTankCapacity() float64
	GetAvgFuelRate() float64
	GetSpeed() float64
}

func GetKmLeft(car Car, fuelAmount float64) (float64, error) {
	if fuelAmount <= 0 {
		return -1, errors.New("fuelAmount is less or equals zero")
	}

	if fuelAmount > car.GetTankCapacity() {
		return -1, errors.New("fuelAmount is too much for the car's fuel tank")
	}

	return fuelAmount / car.GetAvgFuelRate() * 100, nil
}

func GetTravelTime(car Car, fuelAmount float64, distanceByKm float64) (float64, error) {
	if fuelAmount <= 0 {
		return -1, errors.New("fuelAmount is less or equals zero")
	}

	if fuelAmount > car.GetTankCapacity() {
		return -1, errors.New("fuelAmount is too much for the car's fuel tank")
	}

	if distanceByKm <= 0 {
		return -1, errors.New("distanceByKm is less or equals zero")
	}

	fueldNeeded := car.GetAvgFuelRate() / 100 * distanceByKm

	if fueldNeeded > fuelAmount {
		return -1, errors.New("fuelAmount is too low for this distance")
	}

	if car.GetSpeed() <= 0 {
		return -1, errors.New("car's speed is less or equals zero")
	}

	return distanceByKm / car.GetSpeed(), nil
}
