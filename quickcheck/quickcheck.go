package quickcheck

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Int returns an arbitrary non-negative integer value
func Int() int {
	return rand.Int()
}

// Boolean returns an arbitrary boolean value
func Boolean() bool {
	return Int()%2 == 0
}

// String returns an arbitrary string value
func String() string {
	randomString := []byte{}

	close := Boolean()
	for !close {
		randomString = append(randomString, byte(Int()))
		close = Boolean()
	}

	return string(randomString)
}

// StringContainingOnly returns an arbitrary string value contains only the specified chars
func StringContainingOnly(chars []rune) string {
	randomString := []rune{}

	close := Boolean()
	for !close {
		randomString = append(randomString, chars[Int()%len(chars)])
		close = Boolean()
	}

	return string(randomString)
}

// Check validates a property
func Check(t *testing.T, name string, property func() bool) {
	for i := 0; i < 100; i++ {
		if !property() {
			t.Fatalf("Property %s has failed. %v of 100 passed.", name, i)
		}
	}
}
