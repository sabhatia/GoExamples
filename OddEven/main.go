package main

import "fmt"

func main() {
	var numList [11]int
	fmt.Println("Number Checker Code")

	const rangeStart = 0
	const rangeEnd = len(numList)

	for i := rangeStart; i < rangeEnd; i++ {
		numList[i] = i
	}

	for _, num := range numList {
		if isEven(num) {
			fmt.Println(num, "is Even")
		} else {
			fmt.Println(num, "is Odd")
		}
	}
}

func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}
