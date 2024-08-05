package test

import "fmt"

/***
Write a program that prints the numbers from 1 to 100. But for multiples of three print
"Fizz" instead of the number and for the multiples of five print "Buzz". For numbers
 which are multiples of both three and five print "FizzBuzz".*/

 var num int =100

 func Solution(){

	for i:=0;i<num;i++{

		if (i % 3 ==0){
			if(i %5 ==0){
				fmt.Println("FizzBuzz",i)
			}
			fmt.Println("Fizz",i)


		}else if(i % 5 ==0){
			fmt.Println("Buzz",i)

		}
	}

 }