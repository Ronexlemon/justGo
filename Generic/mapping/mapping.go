package mapping

import (
	"fmt"

	"golang.org/x/exp/constraints"
)
type SunConstant interface{
	int64 | float64 | string
}

func MapValues[T constraints.Ordered](values []T, mapfunc func(T) T) []T {
	var newValues []T

	for _, v := range values {
		newValue := mapfunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

// without Generic
// start
func sumInts(m map[string]int64) int64 {

	var s int64

	for _, value := range m {
		s += value
	}
	return s
}

func sumFloats(m map[string]float64) float64 {

	var s float64

	for _, value := range m {
		s += value
	}
	return s
}



func WithoutGeneric() {
	ints := map[string]int64{"first": 20, "second": 30}
	floats := map[string]float64{"first": 20.50, "second": 30.05}

	fmt.Printf("Non-Generic Sums: %v and %v \n", sumInts(ints), sumFloats(floats))

}

//end

func sum[T SunConstant | constraints.Ordered](m map[string] T)T{
	var s T
	for _, value := range m {
		s += value
		}
		return s
}


func WithGeneric(){
	ints := map[string]int64{"first": 20, "second": 30}
	floats := map[string]float64{"first": 20.50, "second": 30.05}
	fmt.Printf("Generic Sums: %v and %v \n", sum[int64](ints), sum[float64](floats))
}