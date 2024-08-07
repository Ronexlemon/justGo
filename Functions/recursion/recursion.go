package recursion

func Factorial(x uint)uint{
if x ==0{
	return 1
}

return x * Factorial(x-1)

}

//eg fact 5*4*3*2*1
