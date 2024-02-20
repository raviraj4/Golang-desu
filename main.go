package main

import "fmt"

func fibo(num int) {

	if num < 1 {
		fmt.Println(num)
	}
	// fmt.Println(a)
	// fmt.Println(b)
	a, b := 0, 1
	for i := 2; i <= num; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
}

func main() {
	fibo(10)
	// ----
	// var name = "arnab"
	// var number int = 69
	// fmt.Println("number is , ", number)

	// fmt.Println("good morning, " + name)
}
