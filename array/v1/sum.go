package main

import "fmt"

func Sum(a [5]int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Sum(myArray))
}
