package quickcheck

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// RandomGenerator random value generator
type RandomGenerator func() reflect.Value

var generatorsByType map[reflect.Type]RandomGenerator = map[reflect.Type]RandomGenerator{
	reflect.TypeOf(""): RandomStringGenerator,
}

// AddRandomGenerator installs a random generator
func AddRandomGenerator(generator RandomGenerator) {
	randomValue := generator()
	t := randomValue.Type()
	_, exists := generatorsByType[t]
	if exists {
		panic("quickcheck: already exists")
	}
	generatorsByType[t] = generator
	return
}

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

// RandomStringGenerator returns a random value of type strings
func RandomStringGenerator() reflect.Value {
	randomString := []byte{}

	close := Boolean()
	for !close {
		randomString = append(randomString, byte(Int()))
		close = Boolean()
	}

	return reflect.ValueOf(string(randomString))
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
func Check(t *testing.T, name string, propertyFunc interface{}) {
	propertyFuncValue := reflect.ValueOf(propertyFunc)
	propertyFuncType := propertyFuncValue.Type()

	// Validate the property func number of return values and  type
	if propertyFuncType.NumOut() != 1 {
		panic("quickcheck: propertyFunc should have on return value")
	}

	if propertyFuncType.Out(0).Kind() != reflect.Bool {
		panic("quickcheck: propertyFunc should return a value of type bool")
	}

	// Check if the property func needs any arbitrary value
	if propertyFuncType.NumIn() == 0 { // No need for any arbitrary value
		returns := propertyFuncValue.Call([]reflect.Value{})
		passed := returns[0].Bool()
		if !passed {
			t.Fatalf("Property %s has failed.", name)
		}
	} else {
		generators := make([]RandomGenerator, 0, propertyFuncType.NumIn())
		for i := 0; i < propertyFuncType.NumIn(); i++ {
			generators = append(generators, generatorsByType[propertyFuncType.In(i)])
		}

		values := make([]reflect.Value, propertyFuncType.NumIn())
		for i := 0; i < 100; i++ {
			for i := 0; i < len(values); i++ {
				values[i] = generators[i]()
			}

			returns := propertyFuncValue.Call(values)
			passed := returns[0].Bool()
			if !passed {
				t.Logf("Args: %v", values)
				t.Fatalf("Property %s has failed. %v of 100 passed.", name, i)
			}
		}
	}
}
