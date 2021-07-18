package helpers

import (
	"math/rand"
	"time"
)

func GetResultForProbability(probability float64) bool {
	probability = probability * 100

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	if randomNumber < int(probability) {
		return true
	} else {
		return false
	}
}
