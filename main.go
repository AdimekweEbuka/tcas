package main

import ( 
	"fmt"
	"math/rand"
	"time"
)
const (
	minimumAltitude = 1000
	maximumAltitude = 1550
	minimumBearing = 0
	maximumBearing = 60
)

// Airport is an airport struct
type Airport struct {
	name string
}

// Airplane is an airplane struct
type Airplane struct {
	name string
	callSign string
	altitude int
	bearing int
	hasTaxied 	bool
}

func (a *Airplane) taxi() {
	fmt.Println(a.callSign, "you have permission to taxi!")
	a.hasTaxied = true
}

func (a *Airplane) takeOff() {
	
	if a.hasTaxied {
		fmt.Println(a.callSign, "you have permission to takeoff!")
		rand.Seed(time.Now().UnixNano())
		a.bearing = generateRandomInt(minimumBearing, maximumBearing)

		a.altitude = generateRandomInt(minimumAltitude, maximumAltitude)
	
		
	} else {
		fmt.Println(a.callSign, "please proceed to the taxi station")
	}
}

func generateRandomInt(min, max int) int {
	return (rand.Intn(max - min + 1) + min)
}

func initiateTCAS(airplaneOne, airplaneTwo Airplane) {

	initiate:
	if airplaneOne.altitude == airplaneTwo.altitude && airplaneOne.bearing == airplaneTwo.bearing {
		fmt.Println("WARNING!!! COLLISION AHEAD!!!")
	}
	rand.Seed(time.Now().UnixNano())
	airplaneOne.bearing = generateRandomInt(minimumBearing, maximumBearing)
	airplaneTwo.bearing = generateRandomInt(minimumBearing, maximumBearing)
	airplaneOne.altitude = generateRandomInt(minimumAltitude, maximumAltitude)
	airplaneTwo.altitude = generateRandomInt(minimumAltitude, maximumAltitude)
	goto initiate
}

// func generateRandomFloat(min, max float64) float64 {
// 	return (min + rand.Float64() * (max - min))
// }

// func roundUp(number float64) float64 {
// 	return (math.Floor(number * 100) / 100)
// }

func main() {
	airplane1 := Airplane {
		name: "Emirates",
		callSign: "EMR076",
	}

	airplane2 := Airplane {
		name: "Air Peace",
		callSign: "AP327",
	}

	airplane1.taxi()
	airplane1.takeOff()

	airplane2.taxi()
	airplane2.takeOff()

	initiateTCAS(airplane1, airplane2)
}
