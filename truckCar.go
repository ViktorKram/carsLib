package carslib

import (
	"errors"
	"math"
	"strings"
)

type truckCar struct {
	typeofCar    string
	tankCapacity float64
	avgFuelRate  float64
	Speed        float64
	loadCapacity float64
	loadPenalty  float64
	penaltyStep  float64
}

func NewTruckCar(typeOfCar string, tankCapacity, avgFuelRate, loadCapacity float64) (*truckCar, error) {
	err := isTruckCarCtorValid(typeOfCar, tankCapacity, avgFuelRate, loadCapacity)
	if err != nil {
		return nil, err
	}

	return &truckCar{
		typeofCar:    typeOfCar,
		tankCapacity: tankCapacity,
		avgFuelRate:  avgFuelRate,
		loadCapacity: loadCapacity,
		loadPenalty:  0.04,
		penaltyStep:  200,
	}, nil
}

func (c truckCar) GetType() string {
	return c.typeofCar
}

func (c truckCar) GetTankCapacity() float64 {
	return c.tankCapacity
}

func (c truckCar) GetAvgFuelRate() float64 {
	return c.avgFuelRate
}

func (c truckCar) GetSpeed() float64 {
	return c.Speed
}

func (c truckCar) GetLoadCapacity() float64 {
	return c.loadCapacity
}

func (c truckCar) GetKmLeftWithLoad(load, fuel float64) (float64, error) {
	if load <= 0 {
		return -1, errors.New("load is less or equal zero")
	}
	if fuel <= 0 {
		return -1, errors.New("fuel is less or equal zero")
	}
	if fuel > c.tankCapacity {
		return -1, errors.New("fuel is more than car's tank capacity")
	}
	if load > c.loadCapacity {
		return -1, errors.New("load is more than car's load capacity")
	}

	newAvgFuelRate := c.GetAvgFuelRate()
	penaltySteps := math.Ceil(load / c.penaltyStep)

	for i := 0; i < int(penaltySteps); i++ {
		newAvgFuelRate += newAvgFuelRate * c.loadPenalty
	}

	return fuel / newAvgFuelRate * 100, nil
}

func isTruckCarCtorValid(typeOfCar string, tankCapacity, avgFuelRate, loadCapacity float64) error {
	if strings.TrimSpace(typeOfCar) == "" {
		return errors.New("typeOfCar is null or empty")
	}
	if tankCapacity <= 0 {
		return errors.New("tankCapacity is less or equals zero")
	}
	if avgFuelRate <= 0 {
		return errors.New("avgFuelRate is less or equals zero")
	}
	if loadCapacity <= 0 {
		return errors.New("loadCapacity is less or equals zero")
	}

	return nil
}
