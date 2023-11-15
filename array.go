package array

import (
	"fmt"
	"sort"
)

// todo Check the comments

// Append a value to the array
func Append(array []any, value any) []any {
	array = append(array, value)
	return array
}

// Length returns the length of the array
func Length(array []any) int {
	return len(array)
}

// Remove the value at the index
// todo to upgrade with error
func Remove(array []any, index int) []any {
	if index < 0 || index >= len(array) {
		fmt.Println("Error | The index isn't in the array")

	}
	array = append(array[:index], array[index+1:]...)
	return array
}

// Find returns the index of the value in the array
// todo to test
func Find(array []any, value any) int {
	size := Length(array)
	for i := 0; i < size; i++ {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// Contain test if the value is in the array return True if inside and False if not inside
// todo to test
func Contain(array []any, value any) bool {
	size := Length(array)
	for i := 0; i < size; i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}

// IsEqual test two array and return true if same or false if different
func IsEqual(FirstArray []any, SecondArray []any) bool {
	size := Length(FirstArray)
	if Length(FirstArray) != Length(SecondArray) {
		return false
	} else {
		for i := 0; i < size; i++ {
			if FirstArray[i] != SecondArray[i] {
				return false
			}
		}
		return true
	}
}

// SortAsc sorts the array in ascending order
// todo to modify and upgrade
func SortAsc(array []any) {
	sort.Slice(array, func(i, j int) bool {
		switch array[i].(type) {
		case int:
			return array[i].(int) < array[j].(int)
		case float64:
			return array[i].(float64) < array[j].(float64)
		case string:
			return array[i].(string) < array[j].(string)
		default:
			return false
		}
	})
}

// SortDesc sorts the array in descending order
// todo to modify and upgrade
func SortDesc(array []any) {
	sort.Slice(array, func(i, j int) bool {
		switch array[i].(type) {
		case int:
			return array[i].(int) > array[j].(int)
		case float64:
			return array[i].(float64) > array[j].(float64)
		case string:
			return array[i].(string) > array[j].(string)
		default:
			return false
		}
	})
}

// Max returns the maximum value in the array
func Max(array []any) any {
	size := Length(array)
	if size == 0 {
		return nil
	}

	max := array[0]
	for i := 1; i < size; i++ {
		if lessThan(max, array[i]) {
			max = array[i]
		}
	}
	return max
}

// Min returns the minimum value in the array
func Min(array []any) any {
	size := Length(array)
	if size == 0 {
		return nil
	}
	min := array[0]
	for i := 1; i < size; i++ {
		if lessThan(array[i], min) {
			min = array[i]
		}
	}
	return min
}

func lessThan(a any, b any) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case float64:
		return a.(float64) < b.(float64)
	case string:
		return len(a.(string)) < len(b.(string))
	default:
		return false
	}
}

// Slice the array from start to end
// todo to test
func Slice(array []any, start, end int) []any {
	if start < 0 {
		start = 0
	}
	if end > Length(array) {
		end = Length(array)
	}
	if start >= end {
		return make([]interface{}, 0)
	}
	return array[start : end+1]
}
