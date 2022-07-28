package main

import "fmt"

func main() {
	//mt.Println(fibonacci(4))
	//var fact = factorial(20)
	//fmt.Println(fact)
	allPrimes(17)
}

func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func factorial(x int) int {
	y := 1
	if x == 0 {
		return y
	}
	y = y * (x)
	return factorial(x-1) * y
}

func allPrimes(x int) {
	var isPrime bool
	for j := x; j > 1; j-- {
		isPrime = true
		for i := 2; i < j; i++ {
			if j%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(j, " ")
		}
	}
}
