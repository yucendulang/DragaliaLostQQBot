package repeatBot

import "math/rand"

func RandomSummonCard() int {
	r := rand.Intn(2)
	if r != 0 {
		return RandomSummonCard() * (r + 1)
	}
	return r + 1
}
