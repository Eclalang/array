package array

import (
	"fmt"
	"sort"
)

// Define The Array Type that's an list of interface{} values
type Array struct {
	values []interface{}
}

// Functions that initialize the array
func NewArray() *Array {
	return &Array{values: make([]interface{}, 0)}
}

// Function that appends a value to the array
func (a *Array) Add(value interface{}) {
	a.values = append(a.values, value)
}

// Function that returns the value at the index
func (a *Array) Get(index int) interface{} {
	return a.values[index]
}

// Function that sets the value at the index
func (a *Array) Set(index int, value interface{}) {
	a.values[index] = value
}

// Function that returns the length of the array
func (a *Array) Length() int {
	return len(a.values)
}

// Function that removes the value at the index
func (a *Array) Remove(index int) {
	if index < 0 || index >= len(a.values) {
		return
	}

	a.values = append(a.values[:index], a.values[index+1:]...)
}

// Function that prints the array
func (a *Array) Print() {
	fmt.Println(a.values)
}

// Func that maps over the array and returns a new array with the result of the function fn
func (a *Array) Map(fn func(interface{}) interface{}) *Array {
	newArray := NewArray()
	for _, v := range a.values {
		newArray.Add(fn(v))
	}
	return newArray
}

// Func that filters the array and returns a new array with the filtered values of the function fn
func (a *Array) Filter(fn func(interface{}) bool) *Array {
	newArray := NewArray()
	for _, v := range a.values {
		if fn(v) {
			newArray.Add(v)
		}
	}
	return newArray
}

// Func that reduces the array and returns a new array with the reduced values of the function fn
func (a *Array) Reduce(fn func(interface{}, interface{}) interface{}, initialValue interface{}) interface{} {
	accumulator := initialValue
	for _, v := range a.values {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

// Func that sorts the array, this sort depends on the comparator function
func (a *Array) Sort(comparator func(interface{}, interface{}) bool) {
	sort.Slice(a.values, func(i, j int) bool {
		return comparator(a.values[i], a.values[j])
	})
}

// Function that returns the index of the value in the array using binary search
func (a *Array) BinarySearch(value interface{}) int {
	return binarySearch(a.values, value, 0, len(a.values)-1)
}

func binarySearch(values []interface{}, value interface{}, left int, right int) int {
	if right < left {
		fmt.Println("Value not found")
		return -1
	}

	mid := (left + right) / 2

	if values[mid] == value {
		return mid
	} else if lessThan(value, values[mid]) {
		return binarySearch(values, value, left, mid-1)
	} else {
		return binarySearch(values, value, mid+1, right)
	}
}

func lessThan(a interface{}, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case string:
		return a.(string) < b.(string)
	default:
		return false
	}
}

// Function that sorts the array in ascending order
func (a *Array) SortAsc() {
	sort.Slice(a.values, func(i, j int) bool {
		switch a.values[i].(type) {
		case int:
			return a.values[i].(int) < a.values[j].(int)
		case string:
			return a.values[i].(string) < a.values[j].(string)
		default:
			return false
		}
	})
}

// Function that sorts the array in descending order
func (a *Array) SortDesc() {
	sort.Slice(a.values, func(i, j int) bool {
		switch a.values[i].(type) {
		case int:
			return a.values[i].(int) > a.values[j].(int)
		case string:
			return a.values[i].(string) > a.values[j].(string)
		default:
			return false
		}
	})
}

// Function that returns the maximum value in the array
func (a *Array) Max() interface{} {
	if len(a.values) == 0 {
		return nil
	}

	max := a.values[0]
	for i := 1; i < len(a.values); i++ {
		if lessThan(max, a.values[i]) {
			max = a.values[i]
		}
	}
	return max
}

// Function that returns the minimum value in the array
func (a *Array) Min() interface{} {
	if len(a.values) == 0 {
		return nil
	}

	min := a.values[0]
	for i := 1; i < len(a.values); i++ {
		if lessThan(a.values[i], min) {
			min = a.values[i]
		}
	}
	return min
}

// Function that slices the array from start to end
func (a *Array) Slice(start, end int) *Array {
	if start < 0 {
		start = 0
	}
	if end > a.Length() {
		end = a.Length()
	}
	if start >= end {
		return NewArray()
	}
	return &Array{values: a.values[start : end+1]}
}
