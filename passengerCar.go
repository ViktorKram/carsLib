package carslib

import (
	"errors"
	"strings"
)

type passengerCar struct {
	typeofCar           string
	tankCapacity        float64
	avgFuelRate         float64
	seats               uint
	penaltyPerPassenger float64
	Speed               float64
}

func NewPassengerCar(typeOfCar string, tankCapacity, avgFuelRate float64, seats uint) (*passengerCar, error) {
	err := isPassengerCtorValid(typeOfCar, tankCapacity, avgFuelRate)
	if err != nil {
		return nil, err
	}

	return &passengerCar{
		typeofCar:           typeOfCar,
		tankCapacity:        tankCapacity,
		avgFuelRate:         avgFuelRate,
		seats:               seats,
		penaltyPerPassenger: 0.06,
	}, nil
}

func (c passengerCar) GetType() string {
	return c.typeofCar
}

func (c passengerCar) GetTankCapacity() float64 {
	return c.tankCapacity
}

func (c passengerCar) GetAvgFuelRate() float64 {
	return c.avgFuelRate
}

func (c passengerCar) GetSpeed() float64 {
	return c.Speed
}

func (c passengerCar) GetSeatsCount() uint {
	return c.seats
}

func (c passengerCar) GetKmLeftWithPassengers(fuel float64, passengers uint) (float64, error) {
	if fuel <= 0 {
		return 0, errors.New("fuelAmount is less or equals zero")
	}
	if fuel > c.tankCapacity {
		return 0, errors.New("fuelAmount is more than the car's tank capacity")
	}
	if passengers > c.seats {
		return 0, errors.New("passengers are more than the car's seats")
	}

	newAvgFuelRate := c.avgFuelRate

	for i := 0; i < int(passengers); i++ {
		newAvgFuelRate += newAvgFuelRate * c.penaltyPerPassenger
	}

	return fuel / newAvgFuelRate * 100, nil
}

func isPassengerCtorValid(typeOfCar string, tankCapacity, avgFuelRate float64) error {
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
