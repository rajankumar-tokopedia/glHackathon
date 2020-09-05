package helper

import "math/rand"

func GetRandomFloats() float32 {
	res := 1 + rand.Float32()*(99)
	return res
}
