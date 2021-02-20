package quickcheck

import (
	"math/rand"
	"reflect"
)

// RandomGenerator random generator
type RandomGenerator struct {
	generator func() interface{}
}

// Type ...
func (g RandomGenerator) Type() reflect.Type {
	return g.Value().Type()
}

// Value ...
func (g RandomGenerator) Value() reflect.Value {
	return reflect.ValueOf(g.Generate())
}

// Generate ...
func (g RandomGenerator) Generate() interface{} {
	return g.generator()
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
	s := []byte{}

	close := Boolean()
	for !close {
		s = append(s, byte(Int()))
		close = Boolean()
	}

	return string(s)
}

// StringContainingOnly returns an arbitrary string value containing only the specified characters
func StringContainingOnly(chars []rune) string {
	randomString := []rune{}

	close := Boolean()
	for !close {
		randomString = append(randomString, chars[Int()%len(chars)])
		close = Boolean()
	}

	return string(randomString)
}

func OneOf(values ...interface{}) interface{} {
	return values[Int()%len(values)]
}

// CreateOneOfGenerator creates a 'one of' random generator
func CreateOneOfGenerator(values ...interface{}) func() interface{} {
	return func() interface{} {
		return values[Int()%len(values)]
	}
}
