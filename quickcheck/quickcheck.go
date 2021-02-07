package quickcheck

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var generatorsByType map[reflect.Type]RandomGenerator = map[reflect.Type]RandomGenerator{}

func init() {
	rand.Seed(time.Now().Unix())
	AddRandomGenerator(func() interface{} { return String() })
}

// AddRandomGenerator instals
func AddRandomGenerator(generatorFunc func() interface{}) {
	g := RandomGenerator{generatorFunc}

	t := g.Type()
	_, exists := generatorsByType[t]
	if exists {
		panic("quickcheck: already exists")
	}

	generatorsByType[t] = g
	return
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
				values[i] = generators[i].Value()
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
