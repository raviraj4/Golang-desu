package main

import (
	"fmt"
	"math"
)
// function to solve the fibonacci
func fibo(num int) {
	if num < 1 {
		fmt.Println(num)
	}
	a, b := 0, 1
	fmt.Println(a)
	fmt.Println(b)
	for i := 2; i <= num-1; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
}
// function using the math module
func trig(arg float64) (float64, float64, float64) {
	SinA := func(angle float64) float64 {
		return math.Sin(angle)
	}

	CosA := func(angle float64) float64 {
		return math.Cos(angle)
	}
	TanA := func(angle float64) float64 {
		return math.Tan(angle)
	}
	
	res1 := SinA(arg)
	res2 := CosA(arg)
	res3 := TanA(arg)

	return res1, res2, res3
}

func main() {
	 
	 fmt.Println(trig(45))
	// printing in golang
	// fmt.Println("number is , ", number)

	// array
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 3
	fmt.Println(arr1)

	// DATATYPES IN GOLANG
	// var name = "arnab"
	// var number int = 69
	// var isMale bool = true
	// var number float64 = 69.3

	// FIBONAACCI
	// fibo(10)

	// fmt.Println("good morning, " + name)
}
