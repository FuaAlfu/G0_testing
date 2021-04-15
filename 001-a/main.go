package main

import "fmt"

func main() {
	fmt.Println("result: ", sum(2, 2))

}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
