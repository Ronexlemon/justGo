package mapping


func MapValues(values []int,mapfunc func (int) int)[]int{
	var newValues []int

	for _,v := range values{
		newValue := mapfunc(v)
		newValues = append(newValues,newValue)
	}
	return newValues
}