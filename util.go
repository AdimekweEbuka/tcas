package main

import (
	"math/rand"
	"time"
)

//GenerateRandomInt gets random integer
func GenerateRandomInt(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	number := r1.Intn(max)
	return number
}

// tcas
