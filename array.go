package array

import (
	"errors"
	"reflect"
	"slices"
)

// Contain test if the value is in the array return True if inside and False if not inside
func Contain(array any, value any) (bool, error) {
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		arr := array.([]any)
		size := len(arr)
		for i := 0; i < size; i++ {
			if arr[i] == value {
				return true, nil
			}
		}
		return false, nil
	} else {
		return false, errors.New("array is not a slice")
	}
}

// Find returns the index of the value in the array
func Find(array any, value any) (int, error) {
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		arr := array.([]any)
		size := len(arr)
		for i := 0; i < size; i++ {
			if arr[i] == value {
				return i, nil
			}
		}
		return -1, nil
	}
	return -1, errors.New("array is not a slice")
}

// IsEqual test two array and return true if same or false if different
func IsEqual(FirstArray, SecondArray any) (bool, error) {
	if reflect.TypeOf(FirstArray).Kind() == reflect.Slice && reflect.TypeOf(SecondArray).Kind() == reflect.Slice {
		firstArr := FirstArray.([]any)
		secondArr := SecondArray.([]any)
		size := len(firstArr)
		if size != len(secondArr) {
			return false, nil
		} else {
			for i := 0; i < size; i++ {
				if firstArr[i] != secondArr[i] {
					return false, nil
				}
			}
			return true, nil
		}
	}
	return false, errors.New("array is not a slice")
}

// Max returns the maximum value in the array
func Max(array any) (any, error) {
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		arr := array.([]any)
		size := len(arr)
		if size == 0 {
			return nil, nil
		}
		max := arr[0]
		for i := 1; i < size; i++ {
			if lessThan(max, arr[i]) {
				max = arr[i]
			}
		}
		return max, nil
	}
	return nil, errors.New("array is not a slice")
}

// Min returns the minimum value in the array
func Min(array any) (any, error) {
	if reflect.TypeOf(array).Kind() == reflect.Slice {
		arr := array.([]any)
		size := len(arr)
		if size == 0 {
			return nil, nil
		}
		min := arr[0]
		for i := 1; i < size; i++ {
			if lessThan(arr[i], min) {
				min = arr[i]
			}
		}
		return min, nil
	}
	return nil, errors.New("array is not a slice")
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

// Remove the value at the index
func Remove(array any, index int) any {
	typeArray := reflect.TypeOf(array)
	if typeArray.Kind() == reflect.Slice {
		switch typeArray.Elem().String() {
		case "int":
			arr := array.([]int)
			if index < 0 || index >= len(arr) {
				return arr
			}
			return slices.Delete(arr, index, index+1)
		case "float64":
			arr := array.([]float64)
			if index < 0 || index >= len(arr) {
				return arr
			}
			arr = append(arr[:index], arr[index+1:]...)
			return arr
		case "string":
			arr := array.([]string)
			if index < 0 || index >= len(arr) {
				return arr
			}
			arr = append(arr[:index], arr[index+1:]...)
			return arr
		case "bool":
			arr := array.([]bool)
			if index < 0 || index >= len(arr) {
				return arr
			}
			arr = append(arr[:index], arr[index+1:]...)
			return arr
		default:
			arr := array.([]any)
			if index < 0 || index >= len(arr) {
				return arr
			}
			arr = append(arr[:index], arr[index+1:]...)
			return arr
		}
	}
	return errors.New("array is not a slice")
}

/*
// Slice the array from start to end included
func Slice(array any, start, end int) (any, error) {
	var arrRet []any

	if start < 0 {
		start = 0
	}
	if end > Length(array)-1 {
		end = Length(array) - 1
	}
	if start >= end {
		return nil, errors.New("start superior or equal to the end")
	}
	for i := start; i <= end; i++ {
		arrRet = append(arrRet, array[i])
	}
	return arrRet, nil
}

// SortAsc sorts the array in ascending order
func SortAsc(array any) {
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
func SortDesc(array any) {
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

*/
