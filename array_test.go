package array

import (
	"fmt"
	"testing"
)

// TestContain test the function Contain with array of type int, float, string and boolean
func TestContain(t *testing.T) {
	arr := []any{1, 2, 3.3, 4.4, "five", "six", true, false}

	// Test int
	check1, err := Contain(arr, 1)
	if err != nil {
		t.Error("No (1/5) | Array Contain isn't set successfully for int")
	}
	check2, err := Contain(arr, 2)
	if err != nil {
		t.Error("No (1/5) | Array Contain isn't set successfully for int")
	}
	check3, err := Contain(arr, 3)
	if err != nil {
		t.Error("No (1/5) | Array Contain isn't set successfully for int")
	}
	if check1 == true && check2 == true && check3 == false {
		fmt.Println("Yes (1/5) | Array Find is set successfully for int")
	} else {
		t.Error("No (1/5) | Array Find isn't set successfully for int")
	}

	// Test float
	check1, err = Contain(arr, 3.3)
	if err != nil {
		t.Error("No (2/5) | Array Contain isn't set successfully for float(64)")
	}
	check2, err = Contain(arr, 4.4)
	if err != nil {
		t.Error("No (2/5) | Array Contain isn't set successfully for float(64)")
	}
	check3, err = Contain(arr, 5.5)
	if err != nil {
		t.Error("No (2/5) | Array Contain isn't set successfully for float(64)")
	}
	if check1 == true && check2 == true && check3 == false {
		fmt.Println("Yes (2/5) | Array Find is set successfully for float(64)")
	} else {
		t.Error("No (2/5) | Array Find isn't set successfully for float(64)")
	}

	// Test string
	check1, err = Contain(arr, "five")
	if err != nil {
		t.Error("No (3/5) | Array Contain isn't set successfully for string")
	}
	check2, err = Contain(arr, "six")
	if err != nil {
		t.Error("No (3/5) | Array Contain isn't set successfully for string")
	}
	check3, err = Contain(arr, "seven")
	if err != nil {
		t.Error("No (3/5) | Array Contain isn't set successfully for string")
	}

	if check1 == true && check2 == true && check3 == false {
		fmt.Println("Yes (3/5) | Array Find is set successfully for string")
	} else {
		t.Error("No (3/5) | Array Find isn't set successfully for string")
	}

	// Test boolean
	check1, err = Contain(arr, true)
	if err != nil {
		t.Error("No (4/5) | Array Contain isn't set successfully for boolean")
	}
	check2, err = Contain(arr, false)
	if err != nil {
		t.Error("No (4/5) | Array Contain isn't set successfully for boolean")
	}
	if check1 == true && check2 == true {
		fmt.Println("Yes (4/5) | Array Find is set successfully for boolean")
	} else {
		t.Error("No (4/5) | Array Find isn't set successfully for boolean")
	}

	// Test error
	check1, err = Contain(1, 1)
	if err != nil {
		fmt.Println("Yes (5/5) | Array Contain isn't set successfully for wrong type")
	} else {
		t.Error("No (5/5) | Array Contain is set successfully for wrong type")
	}
}

// TestFind test the function Find with array of type int, float, string and boolean
func TestFind(t *testing.T) {
	arr := []any{1, 2, 3.3, 4.4, "five", "six", true, false}

	// Test int
	index1, err := Find(arr, 1)
	if err != nil {
		t.Error("No (1/6) | Array Find isn't set successfully for int")
	}
	index2, err := Find(arr, 2)
	if err != nil {
		t.Error("No (1/6) | Array Find isn't set successfully for int")
	}
	if index1 == 0 && index2 == 1 {
		fmt.Println("Yes (1/6) | Array Find is set successfully for int")
	} else {
		t.Error("No (1/6) | Array Find isn't set successfully for int")
	}

	// Test float
	index1, err = Find(arr, 3.3)
	if err != nil {
		t.Error("No (2/6) | Array Find isn't set successfully for float(64)")
	}
	index2, err = Find(arr, 4.4)
	if err != nil {
		t.Error("No (2/6) | Array Find isn't set successfully for float(64)")
	}
	if index1 == 2 && index2 == 3 {
		fmt.Println("Yes (2/6) | Array Find is set successfully for float(64)")
	} else {
		t.Error("No (2/6) | Array Find isn't set successfully for float(64)")
	}

	// Test string
	index1, err = Find(arr, "five")
	if err != nil {
		t.Error("No (3/6) | Array Find isn't set successfully for string")
	}
	index2, err = Find(arr, "six")
	if err != nil {
		t.Error("No (3/6) | Array Find isn't set successfully for string")
	}
	if index1 == 4 && index2 == 5 {
		fmt.Println("Yes (3/6) | Array Find is set successfully for string")
	} else {
		t.Error("No (3/6) | Array Find isn't set successfully for string")
	}

	// Test boolean
	index1, err = Find(arr, true)
	if err != nil {
		t.Error("No (4/6) | Array Find isn't set successfully for boolean")
	}
	index2, err = Find(arr, false)
	if err != nil {
		t.Error("No (4/6) | Array Find isn't set successfully for boolean")
	}
	if index1 == 6 && index2 == 7 {
		fmt.Println("Yes (4/6) | Array Find is set successfully for boolean")
	} else {
		t.Error("No (4/6) | Array Find isn't set successfully for boolean")
	}

	// Test index out of range
	index1, err = Find(arr, 8)
	if err != nil {
		t.Error("No (5/6) | Array Find isn't set successfully for index out of range")
	}
	if index1 == -1 {
		fmt.Println("Yes (5/6) | Array Find is set successfully for index out of range")
	} else {
		t.Error("No (5/6) | Array Find isn't set successfully for index out of range")
	}

	// Test error
	index1, err = Find(1, 1)
	if err != nil {
		fmt.Println("Yes (6/6) | Array Find isn't set successfully for wrong type")
	} else {
		t.Error("No (6/6) | Array Find is set successfully for wrong type")
	}
}

// TestIsEqual test the function IsEqual with array of type int, float(64), string and boolean
func TestIsEqual(t *testing.T) {
	arrOne := []any{1, 2, 3.3, 4.4, "five", "six", true, false}
	arrTwo := []any{1, 2, 3.3, 4.4, "five", "six", true, false}
	arrThree := []any{1, 2, 3, 4, 5, 6}
	arrFour := []any{2, 1, 3.3, 4.4, "five", "six", false, true}

	equal, err := IsEqual(arrOne, arrTwo)
	if err != nil {
		t.Error("No (1/3) | Array IsEqual isn't set successfully when same")
	}
	if equal {
		fmt.Println("Yes (1/3) | Array IsEqual is set successfully when same")
	} else {
		t.Error("No (1/3) | Array IsEqual isn't set successfully when same")
	}

	equal1, err := IsEqual(arrOne, arrThree)
	if err != nil {
		t.Error("No (2/3) | Array IsEqual isn't set successfully when different")
	}
	equal2, err := IsEqual(arrOne, arrFour)
	if err != nil {
		t.Error("No (2/3) | Array IsEqual isn't set successfully when different")
	}
	if !equal1 && !equal2 {
		fmt.Println("Yes (2/3) | Array IsEqual is set successfully when different")
	} else {
		t.Error("No (2/3) | Array IsEqual isn't set successfully when different")
	}

	// Test error
	equal, err = IsEqual(1, 1)
	if err != nil {
		fmt.Println("Yes (3/3) | Array IsEqual isn't set successfully for wrong type")
	} else {
		t.Error("No (3/3) | Array IsEqual is set successfully for wrong type")
	}
}

// TestMax test the function Max with array of type int, float(64) and string (number of char in the string)
func TestMax(t *testing.T) {
	arrInt := []any{1, 20, 300, 44, 56}
	arrFloat := []any{100.1, 20.2, 3.346598, 458.9, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}
	var arrEmpty []any

	// Test int
	max, err := Max(arrInt)
	if err != nil {
		t.Error("No (1/5) | Array Max isn't set successfully for int")
	}
	if max == 300 {
		fmt.Println("Yes (1/5) | Array Max is set successfully for int")
	} else {
		t.Error("No (1/5) | Array Max isn't set successfully for int")
	}

	// Test float
	max, err = Max(arrFloat)
	if err != nil {
		t.Error("No (2/5) | Array Max isn't set successfully for float(64)")
	}
	if max == 458.9 {
		fmt.Println("Yes (2/5) | Array Max is set successfully for float(64)")
	} else {
		t.Error("No (2/5) | Array Max isn't set successfully for float(64)")
	}

	// Test string
	max, err = Max(arrString)
	if err != nil {
		t.Error("No (3/5) | Array Max isn't set successfully for string")
	}
	if max == "three" {
		fmt.Println("Yes (3/5) | Array Max is set successfully for string")
	} else {
		t.Error("No (3/5) | Array Max isn't set successfully for string")
	}

	// Test empty array
	max, err = Max(arrEmpty)
	if err != nil {
		t.Error("No (4/5) | Array Max isn't set successfully for nil")
	}
	if max == nil {
		fmt.Println("Yes (4/5) | Array Max is set successfully for nil")
	} else {
		t.Error("No (4/5) | Array Max isn't set successfully for nil")
	}

	// Test error
	max, err = Max(1)
	if err != nil {
		fmt.Println("Yes (5/5) | Array Max isn't set successfully for wrong type")
	} else {
		t.Error("No (5/5) | Array Max is set successfully for wrong type")
	}
}

// TestMin test the function Min with array of type int, float(64) and string (number of char in the string)
func TestMin(t *testing.T) {
	arrInt := []any{1, 20, 300, 44, 56}
	arrFloat := []any{100.1, 20.2, 3.346598, 458.9, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}
	var arrEmpty []any

	// Test int
	min, err := Min(arrInt)
	if err != nil {
		t.Error("No (1/5) | Array Min isn't set successfully for int")
	}
	if min == 1 {
		fmt.Println("Yes (1/5) | Array Min is set successfully for int")
	} else {
		t.Error("No (1/5) | Array Min isn't set successfully for int")
	}

	// Test float
	min, err = Min(arrFloat)
	if err != nil {
		t.Error("No (2/5) | Array Min isn't set successfully for float(64)")
	}
	if min == 3.346598 {
		fmt.Println("Yes (2/5) | Array Min is set successfully for float(64)")
	} else {
		t.Error("No (2/5) | Array Min isn't set successfully for float(64)")
	}

	// Test string
	min, err = Min(arrString)
	if err != nil {
		t.Error("No (3/5) | Array Min isn't set successfully for string")
	}
	if min == "one" || min == "two" {
		fmt.Println("Yes (3/5) | Array Min is set successfully for string")
	} else {
		t.Error("No (3/5) | Array Min isn't set successfully for string")
	}

	// Test empty array
	min, err = Min(arrEmpty)
	if err != nil {
		t.Error("No (4/5) | Array Min isn't set successfully for nil")
	}
	if min == nil {
		fmt.Println("Yes (4/5) | Array Min is set successfully for nil")
	} else {
		t.Error("No (4/5) | Array Min isn't set successfully for nil")
	}

	// Test error
	min, err = Min(1)
	if err != nil {
		fmt.Println("Yes (5/5) | Array Min isn't set successfully for wrong type")
	} else {
		t.Error("No (5/5) | Array Min is set successfully for wrong type")
	}
}

/*
// TestRemove test the function Remove with array of type int, float(64), string and boolean
func TestRemove(t *testing.T) {
	arrInt := []int{1, 2, 3, 4, 5}
	arrFloat := []any{1.1, 2.2, 3.3, 4.4, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}
	arrBool := []any{true, false, true, false, true}

	arrInt = Remove(arrInt, 2)
	arrFloat = Remove(arrFloat, 2)
	arrString = Remove(arrString, 2)
	arrBool = Remove(arrBool, 2)

	if arrInt[2] == 4 {
		fmt.Println("Yes (1/4) | Array Remove is set successfully for int")
	} else {
		t.Error("No (1/4) | Array Remove isn't set successfully for int")
	}
	if arrFloat[2] == 4.4 {
		fmt.Println("Yes (2/4) | Array Remove is set successfully for float(64)")
	} else {
		t.Error("No (2/4) | Array Remove isn't set successfully for float(64)")
	}
	if arrString[2] == "four" {
		fmt.Println("Yes (3/4) | Array Remove is set successfully for string")
	} else {
		t.Error("No (3/4) | Array Remove isn't set successfully for string")
	}
	if arrBool[2] == false {
		fmt.Println("Yes (4/4) | Array Remove is set successfully for boolean")
	} else {
		t.Error("No (4/4) | Array Remove isn't set successfully for boolean")
	}
}

// TestSlice test the function Slice with array of type int, float(64), string and boolean
func TestSlice(t *testing.T) {
	arrInt := []any{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arrIntExpected := []any{8, 7, 6, 5, 4}
	arrFloat := []any{9.9, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1, 0.0}
	arrFloatExpected := []any{9.9, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3, 2.2}
	arrString := []any{"eeeee", "dddd", "ccc", "bb", "a"}
	arrStringExpected := []any{"dddd", "ccc", "bb"}
	arrBool := []any{true, false, true, true, false, false, true}
	arrBoolExpected := []any{false, false, true}

	arrInt, _ = Slice(arrInt, 1, 5)
	arrFloat, _ = Slice(arrFloat, -10, 7)
	arrString, _ = Slice(arrString, 1, 3)
	arrBool, _ = Slice(arrBool, 4, 50)
	_, errSlice := Slice(arrInt, 4, 1)

	if IsEqual(arrInt, arrIntExpected) {
		fmt.Println("Yes (1/5) | Array Slice is set successfully for int")
	} else {
		t.Error("No (1/5) | Array Slice isn't set successfully for int")
	}
	if IsEqual(arrFloat, arrFloatExpected) {
		fmt.Println("Yes (2/5) | Array Slice is set successfully for float(64)")
	} else {
		t.Error("No (2/5) | Array Slice isn't set successfully for float(64)")
	}
	if IsEqual(arrString, arrStringExpected) {
		fmt.Println("Yes (3/5) | Array Slice is set successfully for string")
	} else {
		t.Error("No (3/5) | Array Slice isn't set successfully for string")
	}
	if IsEqual(arrBool, arrBoolExpected) {
		fmt.Println("Yes (4/5) | Array Slice is set successfully for boolean")
	} else {
		t.Error("No (4/5) | Array Slice isn't set successfully for boolean")
	}
	if errSlice != nil {
		fmt.Println("Yes (5/5) | Array Slice is set successfully for wrong start and end values")
	} else {
		t.Error("No (5/5) | Array Slice isn't set successfully for wrong start and end values")
	}
}

// TestSortAsc test the function SortAsc with array of type int, float(64) and string
func TestSortAsc(t *testing.T) {
	arrInt := []int{3, 6, 9, 2, 5, 8, 0, 1, 4, 7}
	arrIntExpected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrFloat := []any{3.3, 6.6, 9.9, 2.2, 5.5, 8.8, 0.0, 1.1, 4.4, 7.7}
	arrFloatExpected := []any{0.0, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	arrString := []any{"a", "ccc", "bb", "eeeee", "dddd"}
	arrStringExpected := []any{"a", "bb", "ccc", "dddd", "eeeee"}
	arrBool := []any{true, false, true, true, false, false, true}
	//arrMixed := []any{1, 2, "str", "a"}

	SortAsc(arrInt)
	SortAsc(arrFloat)
	SortAsc(arrString)
	SortAsc(arrBool)
	//SortAsc(arrMixed)

	if IsEqual(arrInt, arrIntExpected) {
		fmt.Println("Yes (1/5) | Array SortAsc is set successfully for int")
	} else {
		t.Error("No (1/5) | Array SortAsc isn't set successfully for int")
	}
	if IsEqual(arrFloat, arrFloatExpected) {
		fmt.Println("Yes (2/5) | Array SortAsc is set successfully for float(64)")
	} else {
		t.Error("No (2/5) | Array SortAsc isn't set successfully for float(64)")
	}
	if IsEqual(arrString, arrStringExpected) {
		fmt.Println("Yes (3/5) | Array SortAsc is set successfully for string")
	} else {
		t.Error("No (3/5) | Array SortAsc isn't set successfully for string")
	}
	if IsEqual(arrBool, []any{false, false, true, true, true}) {
		fmt.Println("Yes (4/5) | Array SortAsc is set successfully for boolean")
	} else {
		t.Error("No (4/5) | Array SortAsc isn't set successfully for boolean")
	}
}

// TestSortDesc test the function SortAsc with array of type int, float(64) and string
func TestSortDesc(t *testing.T) {
	arrInt := []any{3, 6, 9, 2, 5, 8, 0, 1, 4, 7}
	arrIntExpected := []any{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arrFloat := []any{3.3, 6.6, 9.9, 2.2, 5.5, 8.8, 0.0, 1.1, 4.4, 7.7}
	arrFloatExpected := []any{9.9, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1, 0.0}
	arrString := []any{"a", "ccc", "bb", "eeeee", "dddd"}
	arrStringExpected := []any{"eeeee", "dddd", "ccc", "bb", "a"}

	SortDesc(arrInt)
	SortDesc(arrFloat)
	SortDesc(arrString)

	if IsEqual(arrInt, arrIntExpected) {
		fmt.Println("Yes (1/3) | Array SortDesc is set successfully for int")
	} else {
		t.Error("No (1/3) | Array SortDesc isn't set successfully for int")
	}
	if IsEqual(arrFloat, arrFloatExpected) {
		fmt.Println("Yes (2/3) | Array SortDesc is set successfully for float(64)")
	} else {
		t.Error("No (2/3) | Array SortDesc isn't set successfully for float(64)")
	}
	if IsEqual(arrString, arrStringExpected) {
		fmt.Println("Yes (3/3) | Array SortDesc is set successfully for string")
	} else {
		t.Error("No (3/3) | Array SortDesc isn't set successfully for string")
	}
}
*/
