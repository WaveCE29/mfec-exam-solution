package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{1221, true},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%d) = %t; expected %t", test.input, result, test.expected)
		}
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"IV", 4},
		{"XL", 40},
		{"MMXXIV", 2024},
	}

	for _, test := range tests {
		result := romanToInt(test.input)
		if result != test.expected {
			t.Errorf("romanToInt(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 3, 2}, []int{1, 3, 2}},
		{[]int{1, 3, 1, 1}, []int{1, 3}},
		{[]int{}, []int{}},
		{[]int{5, 5, 5, 5}, []int{5}},
	}

	for _, test := range tests {
		result := removeDuplicates(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("removeDuplicates(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
