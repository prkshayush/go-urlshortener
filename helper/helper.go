package helper

import "math/rand"

const letter = "asdfghjklzxcvbnmqwertyuiop"

func GenRandomString(n int) string{
	b := make([]byte, n)

	// this for loop will take random integer and value from letter and assign to b for generating another url
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}