package carslib

import (
	"errors"
	"strings"
)

type sportCar struct {
	typeofCar    string
	tankCapacity float64
	avgFuelRate  float64
	Speed        float64
}

func NewSportCar(typeOfCar string, tankCapacity, avgFuelRate float64) (*sportCar, error) {
	err := isPassengerCtorValid(typeOfCar, tankCapacity, avgFuelRate)
	if err != nil {
		return nil, err
	}

	return &sportCar{
		typeofCar:    typeOfCar,
		tankCapacity: tankCapacity,
		avgFuelRate:  avgFuelRate,
	}, nil
}

func (c sportCar) GetType() string {
	return c.typeofCar
}

func (c sportCar) GetTankCapacity() float64 {
	return c.tankCapacity
}

func (c sportCar) GetAvgFuelRate() float64 {
	return c.avgFuelRate
}

func (c sportCar) GetSpeed() float64 {
	return c.Speed
}

func isSportCarCtorValid(typeOfCar string, tankCapacity, avgFuelRate float64) error {
	if strings.TrimSpace(typeOfCar) == "" {
		return errors.New("typeOfCar is null or empty")
	}
	if tankCapacity <= 0 {
		return errors.New("tankCapacity is less or equals zero")
	}
	if avgFuelRate <= 0 {
		return errors.New("avgFuelRate is less or equals zero")
	}

	return nil
}
