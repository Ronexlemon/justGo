package mapping

import "golang.org/x/exp/constraints"


func MapValues[T constraints.Ordered](values []T,mapfunc func (T) T)[]T{
	var newValues []T

	for _,v := range values{
		newValue := mapfunc(v)
		newValues = append(newValues,newValue)
	}
	return newValues
}