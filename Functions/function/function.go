package function

func ComputeAverage(x []float64)float64{
	total :=0.0

	for _,value :=range x{
		total +=value
	}

	return (total/float64(len(x)))
}

func ComputeAverageAndsum(x []float64)(float64,float64){
	total :=0.0

	for _,value :=range x{
		total +=value
	}

	return (total/float64(len(x))),total
}