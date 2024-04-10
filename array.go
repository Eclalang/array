package array

import (
	"cmp"
	"errors"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"reflect"
	"sort"
)

var Variables = map[string]eclaType.Type{
	"ASC":  eclaType.Int(0),
	"DESC": eclaType.Int(1),
}

// Contains test if the value is in the array return True if inside and False if not inside
func Contains[T comparable](array []T, value T) bool {
	size := len(array)
	for i := 0; i < size; i++ {
		if reflect.DeepEqual(array[i], value) {
			return true
		}
	}
	return false
}

// Find returns the index of the value in the array
func Find[T comparable](array []T, value T) int {
	size := len(array)
	for i := 0; i < size; i++ {
		if reflect.DeepEqual(array[i], value) {
			return i
		}
	}
	return -1
}

// IsEqual test two array and return true if same or false if different
func IsEqual[T comparable](FirstArray, SecondArray []T) bool {
	size := len(FirstArray)
	if size != len(SecondArray) {
		return false
	} else {
		if reflect.DeepEqual(FirstArray, SecondArray) {
			return true
		}
		return false
	}
}

// Max returns the maximum value in the array
func Max[T cmp.Ordered](array []T) (any, error) {
	size := len(array)
	if size == 0 {
		return nil, errors.New("empty array")
	}
	m := array[0]
	for i := 1; i < size; i++ {
		m = max(m, array[i])
	}
	return m, nil
}

// Min returns the minimum value in the array
func Min[T cmp.Ordered](array []T) (any, error) {
	size := len(array)
	if size == 0 {
		return nil, errors.New("empty array")
	}
	m := array[0]
	for i := 1; i < size; i++ {
		m = min(m, array[i])
	}
	return m, nil
}

// Remove the value at the index
func Remove[T comparable](array []T, index int) []T {
	if len(array) <= index || index < 0 {
		return array
	}
	return append(array[:index], array[index+1:]...)
}

// Slice the array from start to end included
func Slice[T comparable](array []T, start, end int) ([]T, error) {
	var arrRet []T

	if start < 0 {
		start = 0
	}
	if end > len(array)-1 {
		end = len(array) - 1
	}
	if start >= end {
		return array, errors.New("start superior or equal to the end")
	}
	for i := start; i <= end; i++ {
		arrRet = append(arrRet, array[i])
	}
	return arrRet, nil
}

// Sort sorts the array in ascending or descending order
func Sort[T cmp.Ordered](array []T, order int) {
	if order == 0 {
		sort.Slice(array, func(i, j int) bool {
			return array[i] < array[j]
		})
	} else {
		sort.Slice(array, func(i, j int) bool {
			return array[i] > array[j]
		})
	}
}
