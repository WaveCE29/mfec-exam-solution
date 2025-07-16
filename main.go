package main

import (
	"fmt"
	"strconv"
)

// ตรวจสอบว่าเลขเป็นพาลินโดรมหรือไม่
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	n := len(str)

	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

// แปลงเลขโรมันเป็นเลขฐาน 10
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	sum := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if i < n-1 && roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}
	return sum
}

// ลบค่าซ้ำจาก slice โดยเก็บค่าที่เจอครั้งแรกไว้
func removeDuplicates(input []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range input {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// ทดสอบฟังก์ชัน isPalindrome
	fmt.Println("== Check Palindrome ==")
	palindromeInputs := []int{121, -121, 10}
	for _, input := range palindromeInputs {
		output := isPalindrome(input)
		fmt.Printf("Input: %d\nOutput: %t\n================\n", input, output)
	}

	// ทดสอบฟังก์ชัน romanToInt
	fmt.Println("== Convert Roman to Integer ==")
	romanInputs := []string{"III", "LVIII", "MCMXCIV"}
	for _, input := range romanInputs {
		output := romanToInt(input)
		fmt.Printf("Input: \"%s\"\nOutput: %d\n================\n", input, output)
	}

	// ทดสอบฟังก์ชัน removeDuplicates
	fmt.Println("== Remove Duplicates from Slice ==")
	duplicateTests := [][]int{
		{1, 1, 2},
		{1, 3, 2},
		{1, 3, 1, 1},
	}

	for _, input := range duplicateTests {
		output := removeDuplicates(input)
		fmt.Printf("Input: %v\nOutput: %v\n================\n", input, output)
	}
}
