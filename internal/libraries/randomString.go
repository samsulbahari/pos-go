package libraries

import "math/rand"

var text = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString() string {
	b := make([]rune, 15)
	for i := range b {
		b[i] = text[rand.Intn(len(text))]
	}
	return string(b)
}
