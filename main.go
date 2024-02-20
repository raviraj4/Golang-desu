package main

import "fmt"

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

func main() {
	fibo(10)
	
	// declaring (variable, int): 
	// var name = "arnab"
	// var number int = 69

	// printing in golang
	// fmt.Println("number is , ", number)

	// fmt.Println("good morning, " + name)
}
