package processor

import (
	"math"
	"math/big"
)

func unifyRandom(randomInt *big.Int) float64 {
	max := new(big.Int).Lsh(big.NewInt(1), 256)

	randomFloat := new(big.Float).SetInt(randomInt)
	maxFloat := new(big.Float).SetInt(max)
	randomF, _ := new(big.Float).Quo(randomFloat, maxFloat).Float64()
	return randomF
}

func getRandomAccordingDistribution(u float64, downBound, upBound float64) float64 {
	probability := upBound - downBound
	var x float64
	if u < 0.5 {
		x = -probability + math.Sqrt(2*probability*probability*u)
	} else {
		x = probability - math.Sqrt(2*probability*probability*(1-u))
	}

	return (upBound+downBound)/2 + x
}
