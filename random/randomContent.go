package random

import "math/rand"

var sentenseSuffix = []string{"呢!", "呐!"}

func RandomGetSuffix() string {
	r := rand.Intn(len(sentenseSuffix))
	return sentenseSuffix[r]
}
