package main

import "fmt"

func reverseInt(n int) int {
	reversedInt := 0
	for n > 0 {
		reversedInt = reversedInt*10 + n%10
		n /= 10
	}
	return reversedInt
}

func isPalindrome(n int) bool {
	if n == reverseInt(n) {
		return true
	} else {
		return false
	}
}

func main() {

	testNums := []int{121, 1221, 123, 452345, 9876789, 76544567}

	for _, value := range testNums {
		fmt.Printf("Value is %d\n", value)
		fmt.Printf("Reversed value is %d\n", reverseInt(value))
		fmt.Printf("Is it a palindrome? %t\n\n", isPalindrome(value))
	}
}
