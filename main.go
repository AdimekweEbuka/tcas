package main

import (
	"fmt"
	"time"
)

func main() {
	airplane1 := airplane{
		Name:     "Emirates",
		CallSign: "EMR2021",
	}
	airplane2 := airplane{
		Name:     "AirPeace",
		CallSign: "ARP2022",
	}
	airplane3 := airplane{
		Name:     "DANA AIRWAYS",
		CallSign: "DAN3243",
	}

	airplane1.StartTrip()
	fmt.Printf("\n")
	airplane2.StartTrip()
	fmt.Printf("\n")
	airplane3.StartTrip()
	time.Sleep(time.Second * 5)
}

// GITHUB
// git add .
// git commit -m "my message"
// git push origin master
