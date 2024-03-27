package array

import (
	"fmt"
	"reflect"
	"testing"
)

// TestContain test the function Contain with array of type int, float, string and boolean
func TestContain(t *testing.T) {
	arr := []any{1, 2, 3.3, 4.4, "five", "six", true, false}

	// Test int
	check1 := Contain(arr, 1)
	check2 := Contain(arr, 2)
	check3 := Contain(arr, 3)
	if check1 == true && check2 == true && check3 == false {
		fmt.Println("Yes (1/4) | Array Find is set successfully for int")
	} else {
		t.Error("No (1/4) | Array Find isn't set successfully for int")
	}

	// Test float
	check1 = Contain(arr, 3.3)
	check2 = Contain(arr, 4.4)
	check3 = Contain(arr, 5.5)
	if check1 == true && check2 == true && check3 == false {
		fmt.Println("Yes (2/4) | Array Find is set successfully for float64")
	} else {
		t.Error("No (2/4) | Array Find isn't set successfully for float64")
	}

	// Test string
	check1 = Contain(arr, "five")
	check2 = Contain(arr, "six")
	check3 = Contain(arr, "seven")

	if check1 == true && check2 == true && check3 == false {
		fmt.Println("Yes (3/4) | Array Find is set successfully for string")
	} else {
		t.Error("No (3/4) | Array Find isn't set successfully for string")
	}

	// Test boolean
	check1 = Contain(arr, true)
	check2 = Contain(arr, false)
	if check1 == true && check2 == true {
		fmt.Println("Yes (4/4) | Array Find is set successfully for boolean")
	} else {
		t.Error("No (4/4) | Array Find isn't set successfully for boolean")
	}
}

// TestFind test the function Find with array of type int, float, string and boolean
func TestFind(t *testing.T) {
	arr := []any{1, 2, 3.3, 4.4, "five", "six", true, false}

	// Test int
	index1 := Find(arr, 1)
	index2 := Find(arr, 2)
	if index1 == 0 && index2 == 1 {
		fmt.Println("Yes (1/5) | Array Find is set successfully for int")
	} else {
		t.Error("No (1/5) | Array Find isn't set successfully for int")
	}

	// Test float
	index1 = Find(arr, 3.3)
	index2 = Find(arr, 4.4)
	if index1 == 2 && index2 == 3 {
		fmt.Println("Yes (2/5) | Array Find is set successfully for float64")
	} else {
		t.Error("No (2/5) | Array Find isn't set successfully for float64")
	}

	// Test string
	index1 = Find(arr, "five")
	index2 = Find(arr, "six")
	if index1 == 4 && index2 == 5 {
		fmt.Println("Yes (3/5) | Array Find is set successfully for string")
	} else {
		t.Error("No (3/5) | Array Find isn't set successfully for string")
	}

	// Test boolean
	index1 = Find(arr, true)
	index2 = Find(arr, false)
	if index1 == 6 && index2 == 7 {
		fmt.Println("Yes (4/5) | Array Find is set successfully for boolean")
	} else {
		t.Error("No (4/5) | Array Find isn't set successfully for boolean")
	}

	// Test index out of range
	index1 = Find(arr, 8)
	if index1 == -1 {
		fmt.Println("Yes (5/5) | Array Find is set successfully for index out of range")
	} else {
		t.Error("No (5/5) | Array Find isn't set successfully for index out of range")
	}
}

// TestIsEqual test the function IsEqual with array of type int, float64, string and boolean
func TestIsEqual(t *testing.T) {
	arrOne := []any{1, 2, 3.3, 4.4, "five", "six", true, false}
	arrTwo := []any{1, 2, 3.3, 4.4, "five", "six", true, false}
	arrThree := []any{1, 2, 3, 4, 5, 6}
	arrFour := []any{2, 1, 3.3, 4.4, "five", "six", false, true}

	equal := IsEqual(arrOne, arrTwo)
	if equal {
		fmt.Println("Yes (1/2) | Array IsEqual is set successfully when same")
	} else {
		t.Error("No (1/2) | Array IsEqual isn't set successfully when same")
	}

	equal1 := IsEqual(arrOne, arrThree)
	equal2 := IsEqual(arrOne, arrFour)
	if !equal1 && !equal2 {
		fmt.Println("Yes (2/2) | Array IsEqual is set successfully when different")
	} else {
		t.Error("No (2/2) | Array IsEqual isn't set successfully when different")
	}
}

// TestMax test the function Max with array of type int, float64 and string (number of char in the string)
func TestMax(t *testing.T) {
	arrInt := []int{1, 20, 300, 44, 56}
	arrFloat := []float64{100.1, 20.2, 3.346598, 458.9, 5.5}
	arrString := []string{"one", "two", "three", "four", "five"}
	var arrEmpty []int

	// Test int
	intMax, _ := Max(arrInt)
	if intMax == 300 {
		fmt.Println("Yes (1/4) | Array Max is set successfully for int")
	} else {
		t.Error("No (1/4) | Array Max isn't set successfully for int")
	}

	// Test float
	float64Max, _ := Max(arrFloat)
	if float64Max == 458.9 {
		fmt.Println("Yes (2/4) | Array Max is set successfully for float64")
	} else {
		t.Error("No (2/4) | Array Max isn't set successfully for float64")
	}

	// Test string
	stringMax, _ := Max(arrString)
	if stringMax == "two" {
		fmt.Println("Yes (3/4) | Array Max is set successfully for string")
	} else {
		t.Error("No (3/4) | Array Max isn't set successfully for string")
	}

	// Test empty array
	_, err := Max(arrEmpty)
	if err != nil {
		fmt.Println("Yes (4/4) | Array Max is set successfully for NaN")
	} else {
		t.Error("No (4/4) | Array Max isn't set successfully for NaN")
	}
}

// TestMin test the function Min with array of type int, float64 and string (number of char in the string)
func TestMin(t *testing.T) {
	arrInt := []int{1, 20, 300, 44, 56}
	arrFloat := []float64{100.1, 20.2, 3.346598, 458.9, 5.5}
	arrString := []string{"one", "two", "three", "four", "five"}
	var arrEmpty []int

	// Test int
	intMin, _ := Min(arrInt)
	if intMin == 1 {
		fmt.Println("Yes (1/4) | Array Min is set successfully for int")
	} else {
		t.Error("No (1/4) | Array Min isn't set successfully for int")
	}

	// Test float
	float64Min, _ := Min(arrFloat)
	if float64Min == 3.346598 {
		fmt.Println("Yes (2/4) | Array Min is set successfully for float64")
	} else {
		t.Error("No (2/4) | Array Min isn't set successfully for float64")
	}

	// Test string
	stringMin, _ := Min(arrString)
	if stringMin == "five" {
		fmt.Println("Yes (3/4) | Array Min is set successfully for string")
	} else {
		t.Error("No (3/4) | Array Min isn't set successfully for string")
	}

	// Test empty array
	_, err := Min(arrEmpty)
	if err != nil {
		fmt.Println("Yes (4/4) | Array Min is set successfully for NaN")
	} else {
		t.Error("No (4/4) | Array Min isn't set successfully for NaN")
	}
}

// TestRemove test the function Remove with array of type int, float64, string and boolean
func TestRemove(t *testing.T) {
	arrInt := []int{1, 2, 3, 4, 5}
	arrFloat := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	arrString := []string{"one", "two", "three", "four", "five"}
	arrBool := []bool{true, false, true, false, true}
	arrAny := []any{1, 2, "a", true}

	arrInt = Remove(arrInt, 2)
	arrFloat = Remove(arrFloat, 2)
	arrString = Remove(arrString, 2)
	arrBool = Remove(arrBool, 2)
	arrAny = Remove(arrAny, 1)

	newArrInt := Remove(arrInt, -1)
	newArrFloat := Remove(arrFloat, -1)
	newArrString := Remove(arrString, -1)
	newArrBool := Remove(arrBool, -1)
	newArrAny := Remove(arrAny, -1)

	// Test array int
	if arrInt[2] == 4 {
		fmt.Println("Yes (1/10) | Array Remove is set successfully for int")
	} else {
		t.Error("No (1/10) | Array Remove isn't set successfully for int")
	}

	if reflect.DeepEqual(newArrInt, arrInt) {
		fmt.Println("Yes (2/10) | Array Remove is set successfully for wrong index")
	} else {
		t.Error("No (2/10) | Array Remove isn't set successfully for wrong index")
	}
	// Test array float32
	if arrFloat[2] == 4.4 {
		fmt.Println("Yes (3/10) | Array Remove is set successfully for float64")
	} else {
		t.Error("No (3/10) | Array Remove isn't set successfully for float64")
	}
	if reflect.DeepEqual(newArrFloat, arrFloat) {
		fmt.Println("Yes (4/10) | Array Remove is set successfully for wrong index")
	} else {
		t.Error("No (4/10) | Array Remove isn't set successfully for wrong index")
	}

	// Test array string
	if arrString[2] == "four" {
		fmt.Println("Yes (5/10) | Array Remove is set successfully for string")
	} else {
		t.Error("No (5/10) | Array Remove isn't set successfully for string")
	}
	if reflect.DeepEqual(newArrString, arrString) {
		fmt.Println("Yes (6/10) | Array Remove is set successfully for wrong index")
	} else {
		t.Error("No (6/10) | Array Remove isn't set successfully for wrong index")
	}

	// Test array bool
	if arrBool[2] == false {
		fmt.Println("Yes (7/10) | Array Remove is set successfully for boolean")
	} else {
		t.Error("No (7/10) | Array Remove isn't set successfully for boolean")
	}
	if reflect.DeepEqual(newArrBool, arrBool) {
		fmt.Println("Yes (8/10) | Array Remove is set successfully for wrong index")
	} else {
		t.Error("No (8/10) | Array Remove isn't set successfully for wrong index")
	}

	// Test array any
	if arrAny[1] == "a" {
		fmt.Println("Yes (9/10) | Array Remove is set successfully for any")
	} else {
		t.Error("No (9/10) | Array Remove isn't set successfully for any")
	}
	if reflect.DeepEqual(newArrAny, arrAny) {
		fmt.Println("Yes (10/10) | Array Remove is set successfully for wrong index")
	} else {
		t.Error("No (10/10) | Array Remove isn't set successfully for wrong index")
	}
}

// TestSlice test the function Slice with array of type int, float64, string and boolean
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

	equal := IsEqual(arrInt, arrIntExpected)
	if equal {
		fmt.Println("Yes (1/5) | Array Slice is set successfully for int")
	} else {
		t.Error("No (1/5) | Array Slice isn't set successfully for int")
	}

	equal = IsEqual(arrFloat, arrFloatExpected)
	if equal {
		fmt.Println("Yes (2/5) | Array Slice is set successfully for float64")
	} else {
		t.Error("No (2/5) | Array Slice isn't set successfully for float64")
	}

	equal = IsEqual(arrString, arrStringExpected)
	if equal {
		fmt.Println("Yes (3/5) | Array Slice is set successfully for string")
	} else {
		t.Error("No (3/5) | Array Slice isn't set successfully for string")
	}

	equal = IsEqual(arrBool, arrBoolExpected)
	if equal {
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

// TestSortAsc test the function SortAsc with array of type int, float64 and string
func TestSortAsc(t *testing.T) {
	arrInt := []int{3, 6, 9, 2, 5, 8, 0, 1, 4, 7}
	arrIntExpected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrFloat := []float64{3.3, 6.6, 9.9, 2.2, 5.5, 8.8, 0.0, 1.1, 4.4, 7.7}
	arrFloatExpected := []float64{0.0, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	arrString := []string{"a", "ccc", "bb", "eeeee", "dddd"}
	arrStringExpected := []string{"a", "bb", "ccc", "dddd", "eeeee"}

	SortAsc(arrInt)
	SortAsc(arrFloat)
	SortAsc(arrString)

	// Test array int
	equal := IsEqual(arrInt, arrIntExpected)
	if equal {
		fmt.Println("Yes (1/3) | Array SortAsc is set successfully for int")
	} else {
		t.Error("No (1/3) | Array SortAsc isn't set successfully for int")
	}

	// Test array float64
	equal = IsEqual(arrFloat, arrFloatExpected)
	if equal {
		fmt.Println("Yes (2/3) | Array SortAsc is set successfully for float64")
	} else {
		t.Error("No (2/3) | Array SortAsc isn't set successfully for float64")
	}

	// Test array string
	equal = IsEqual(arrString, arrStringExpected)
	if equal {
		fmt.Println("Yes (3/3) | Array SortAsc is set successfully for string")
	} else {
		t.Error("No (3/3) | Array SortAsc isn't set successfully for string")
	}
}

// TestSortDesc test the function SortAsc with array of type int, float64 and string
func TestSortDesc(t *testing.T) {
	arrInt := []int{3, 6, 9, 2, 5, 8, 0, 1, 4, 7}
	arrIntExpected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arrFloat := []float64{3.3, 6.6, 9.9, 2.2, 5.5, 8.8, 0.0, 1.1, 4.4, 7.7}
	arrFloatExpected := []float64{9.9, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1, 0.0}
	arrString := []string{"a", "ccc", "bb", "eeeee", "dddd"}
	arrStringExpected := []string{"eeeee", "dddd", "ccc", "bb", "a"}

	SortDesc(arrInt)
	SortDesc(arrFloat)
	SortDesc(arrString)

	if IsEqual(arrInt, arrIntExpected) {
		fmt.Println("Yes (1/3) | Array SortDesc is set successfully for int")
	} else {
		t.Error("No (1/3) | Array SortDesc isn't set successfully for int")
	}
	if IsEqual(arrFloat, arrFloatExpected) {
		fmt.Println("Yes (2/3) | Array SortDesc is set successfully for float64")
	} else {
		t.Error("No (2/3) | Array SortDesc isn't set successfully for float64")
	}
	if IsEqual(arrString, arrStringExpected) {
		fmt.Println("Yes (3/3) | Array SortDesc is set successfully for string")
	} else {
		t.Error("No (3/3) | Array SortDesc isn't set successfully for string")
	}
}
