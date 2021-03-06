package main

import (
	"fmt"
)

const (
	minimumAltitude = 0
	minimumBearing  = 0
	maximumBearing  = 1
	maximumAltitude = 1
)

var airboneAirplanes = make([]*airplane, 0)

func (a *airplane) taxiCall() {
	fmt.Println(a.CallSign, "You have permission to taxi")
	a.HasTaxied = true
}

func (a *airplane) takeOff() {
	if a.HasTaxied {
		fmt.Println(a.CallSign, "You have permission to takeoff!!")
		a.Altitude = GenerateRandomInt(maximumAltitude)
		a.Bearing = GenerateRandomInt(maximumBearing)
		airboneAirplanes = append(airboneAirplanes, a)
		go a.initiateTCAS()
	} else {
		fmt.Println(a.CallSign, "You do not have takeoff permission")
	}
}

func (a *airplane) cruisingAltitude() {
	for a.Altitude < maximumAltitude {
		a.Altitude += 5000
	}
	fmt.Println("We have reached cruising altitude at", a.Altitude, "feet!")
}

func (a *airplane) initiateTCAS() {
	//Initiate TCAS
	if a.Altitude == airboneAirplanes[0].Altitude && a.Bearing == airboneAirplanes[0].Bearing {
		fmt.Println("Warning", a.CallSign, "and", airboneAirplanes[0].CallSign, "are on a collison course!")
	}

	a.Altitude = GenerateRandomInt(maximumAltitude)
	a.Bearing = GenerateRandomInt(maximumBearing)
	fmt.Println(a.CallSign, "Please proceed to alitude", a.Altitude, "and bearing", a.Bearing)

	airboneAirplanes[0].Altitude = GenerateRandomInt(maximumAltitude)
	airboneAirplanes[0].Bearing = GenerateRandomInt(maximumBearing)
	fmt.Println(airboneAirplanes[0].CallSign, "Please proceed to alitude", airboneAirplanes[0].Altitude, "and bearing", airboneAirplanes[0].Bearing)
}

func (a *airplane) landAirplane() {
	for a.Altitude > minimumAltitude {
		a.Altitude -= 5000
	}
	if a.Altitude < 0 {
		a.Altitude = 0
	}
	fmt.Println("We have successfully landed", a.CallSign, "hope you enjoyed your flight?")
}

//StartTrip function
func (a *airplane) StartTrip() {
	//Taxi Permission
	go a.taxiCall()
	//Takeoff Permision
	go a.takeOff()
	//Cruising Altitude
	go a.cruisingAltitude()
	//Land Airplane
	go a.landAirplane()
}
