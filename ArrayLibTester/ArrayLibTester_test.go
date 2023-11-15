package arraylibtester

import (
	"fmt"
	"testing"

	lib "github.com/Eclalang/array"
)

// TestAppend test the function Append with type int, float(64), string and boolean
func TestAppend(t *testing.T) {
	arr := make([]any, 0)

	arr = lib.Append(arr, 1)
	arr = lib.Append(arr, 2)
	arr = lib.Append(arr, 1.1)
	arr = lib.Append(arr, 2.2)
	arr = lib.Append(arr, "one")
	arr = lib.Append(arr, "two")
	arr = lib.Append(arr, true)
	arr = lib.Append(arr, false)

	if arr[0] == 1 && arr[1] == 2 {
		fmt.Println("Yes | Array Append is set successfully for int")
	} else {
		t.Error("No | Array Append isn't set successfully for int")
	}
	if arr[2] == 1.1 && arr[3] == 2.2 {
		fmt.Println("Yes | Array Append is set successfully for float(64)")
	} else {
		t.Error("No | Array Append isn't set successfully for float(64)")
	}
	if arr[4] == "one" && arr[5] == "two" {
		fmt.Println("Yes | Array Append is set successfully for string")
	} else {
		t.Error("No | Array Append isn't set successfully for string")
	}
	if arr[6] == true && arr[7] == false {
		fmt.Println("Yes | Array Append is set successfully for boolean")
	} else {
		t.Error("No | Array Append isn't set successfully for boolean")
	}
}

// TestLength test the function Length with array of type int, float(64), string and boolean
func TestLength(t *testing.T) {
	arrInt := []any{1, 2, 3, 4, 5}
	arrFloat := []any{1.1, 2.2, 3.3, 4.4, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}
	arrBool := []any{true, false, true, false, true}

	if lib.Length(arrInt) == 5 {
		fmt.Println("Yes | Array Length is set successfully for int")
	} else {
		t.Error("No | Array Length isn't set successfully for int")
	}
	if lib.Length(arrFloat) == 5 {
		fmt.Println("Yes | Array Length is set successfully for float(64)")
	} else {
		t.Error("No | Array Length isn't set successfully for float(64)")
	}
	if lib.Length(arrString) == 5 {
		fmt.Println("Yes | Array Length is set successfully for string")
	} else {
		t.Error("No | Array Length isn't set successfully for string")
	}
	if lib.Length(arrBool) == 5 {
		fmt.Println("Yes | Array Length is set successfully for boolean")
	} else {
		t.Error("No | Array Length isn't set successfully for boolean")
	}
}

// TestRemove test the function Remove with array of type int, float(64), string and boolean
func TestRemove(t *testing.T) {
	arrInt := []any{1, 2, 3, 4, 5}
	arrFloat := []any{1.1, 2.2, 3.3, 4.4, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}
	arrBool := []any{true, false, true, false, true}

	arrInt = lib.Remove(arrInt, 2)
	arrFloat = lib.Remove(arrFloat, 2)
	arrString = lib.Remove(arrString, 2)
	arrBool = lib.Remove(arrBool, 2)

	if arrInt[2] == 4 {
		fmt.Println("Yes | Array Remove is set successfully for int")
	} else {
		t.Error("No | Array Remove isn't set successfully for int")
	}
	if arrFloat[2] == 4.4 {
		fmt.Println("Yes | Array Remove is set successfully for float(64)")
	} else {
		t.Error("No | Array Remove isn't set successfully for float(64)")
	}
	if arrString[2] == "four" {
		fmt.Println("Yes | Array Remove is set successfully for string")
	} else {
		t.Error("No | Array Remove isn't set successfully for string")
	}
	if arrBool[2] == false {
		fmt.Println("Yes | Array Remove is set successfully for boolean")
	} else {
		t.Error("No | Array Remove isn't set successfully for boolean")
	}
}

// TestFind test the function Find with array of type int, float, string and boolean
func TestFind(t *testing.T) {
	arr := []any{1, 2, 3.3, 4.4, "five", "six", true, false}

	if lib.Find(arr, 1) == 0 && lib.Find(arr, 2) == 1 {
		fmt.Println("Yes | Array Find is set successfully for int")
	} else {
		t.Error("No | Array Find isn't set successfully for int")
	}
	if lib.Find(arr, 3.3) == 2 && lib.Find(arr, 4.4) == 3 {
		fmt.Println("Yes | Array Find is set successfully for float(64)")
	} else {
		t.Error("No | Array Find isn't set successfully for float(64)")
	}
	if lib.Find(arr, "five") == 4 && lib.Find(arr, "six") == 5 {
		fmt.Println("Yes | Array Find is set successfully for string")
	} else {
		t.Error("No | Array Find isn't set successfully for string")
	}
	if lib.Find(arr, true) == 6 && lib.Find(arr, false) == 7 {
		fmt.Println("Yes | Array Find is set successfully for boolean")
	} else {
		t.Error("No | Array Find isn't set successfully for boolean")
	}
}

// TestContain test the function Contain with array of type int, float, string and boolean
func TestContain(t *testing.T) {
	arr := []any{1, 2, 3.3, 4.4, "five", "six", true, false}

	if lib.Contain(arr, 1) == true && lib.Contain(arr, 2) == true && lib.Contain(arr, 3) == false {
		fmt.Println("Yes | Array Find is set successfully for int")
	} else {
		t.Error("No | Array Find isn't set successfully for int")
	}
	if lib.Contain(arr, 3.3) == true && lib.Contain(arr, 4.4) == true && lib.Contain(arr, 5.5) == false {
		fmt.Println("Yes | Array Find is set successfully for float(64)")
	} else {
		t.Error("No | Array Find isn't set successfully for float(64)")
	}
	if lib.Contain(arr, "five") == true && lib.Contain(arr, "six") == true && lib.Contain(arr, "seven") == false {
		fmt.Println("Yes | Array Find is set successfully for string")
	} else {
		t.Error("No | Array Find isn't set successfully for string")
	}
	if lib.Contain(arr, true) == true && lib.Contain(arr, false) == true {
		fmt.Println("Yes | Array Find is set successfully for boolean")
	} else {
		t.Error("No | Array Find isn't set successfully for boolean")
	}
}

// TestIsEqual test the function IsEqual with array of type int, float(64), string and boolean
func TestIsEqual(t *testing.T) {
	arrOne := []any{1, 2, 3.3, 4.4, "five", "six", true, false}
	arrTwo := []any{1, 2, 3.3, 4.4, "five", "six", true, false}
	arrThree := []any{1, 2, 3, 4, 5, 6}
	arrFour := []any{2, 1, 3.3, 4.4, "five", "six", false, true}

	if lib.IsEqual(arrOne, arrTwo) {
		fmt.Println("Yes | Array IsEqual is set successfully when same")
	} else {
		t.Error("No | Array IsEqual isn't set successfully when same")
	}
	if !lib.IsEqual(arrOne, arrThree) && !lib.IsEqual(arrOne, arrFour) {
		fmt.Println("Yes | Array IsEqual is set successfully when different")
	} else {
		t.Error("No | Array IsEqual isn't set successfully when different")
	}
}

// TestSortAsc test the function SortAsc with array of type int, float(64) and string
func TestSortAsc(t *testing.T) {
	arrInt := []any{3, 6, 9, 2, 5, 8, 0, 1, 4, 7}
	arrIntExpected := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrFloat := []any{3.3, 6.6, 9.9, 2.2, 5.5, 8.8, 0.0, 1.1, 4.4, 7.7}
	arrFloatExpected := []any{0.0, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	arrString := []any{"a", "ccc", "bb", "eeeee", "dddd"}
	arrStringExpected := []any{"a", "bb", "ccc", "dddd", "eeeee"}

	lib.SortAsc(arrInt)
	lib.SortAsc(arrFloat)
	lib.SortAsc(arrString)

	if lib.IsEqual(arrInt, arrIntExpected) {
		fmt.Println("Yes | Array SortAsc is set successfully for int")
	} else {
		t.Error("No | Array SortAsc isn't set successfully for int")
	}
	if lib.IsEqual(arrFloat, arrFloatExpected) {
		fmt.Println("Yes | Array SortAsc is set successfully for float(64)")
	} else {
		t.Error("No | Array SortAsc isn't set successfully for float(64)")
	}
	if lib.IsEqual(arrString, arrStringExpected) {
		fmt.Println("Yes | Array SortAsc is set successfully for string")
	} else {
		t.Error("No | Array SortAsc isn't set successfully for string")
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

	lib.SortDesc(arrInt)
	lib.SortDesc(arrFloat)
	lib.SortDesc(arrString)

	if lib.IsEqual(arrInt, arrIntExpected) {
		fmt.Println("Yes | Array SortDesc is set successfully for int")
	} else {
		t.Error("No | Array SortDesc isn't set successfully for int")
	}
	if lib.IsEqual(arrFloat, arrFloatExpected) {
		fmt.Println("Yes | Array SortDesc is set successfully for float(64)")
	} else {
		t.Error("No | Array SortDesc isn't set successfully for float(64)")
	}
	if lib.IsEqual(arrString, arrStringExpected) {
		fmt.Println("Yes | Array SortDesc is set successfully for string")
	} else {
		t.Error("No | Array SortDesc isn't set successfully for string")
	}
}

// TestMax test the function Max with array of type int, float(64) and string
func TestMax(t *testing.T) {
	arrInt := []any{1, 20, 300, 44, 56}
	arrFloat := []any{100.1, 20.2, 3.346598, 458.9, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}

	if lib.Max(arrInt) == 300 {
		fmt.Println("Yes | Array Max is set successfully for int")
	} else {
		t.Error("No | Array Max isn't set successfully for int")
	}
	if lib.Max(arrFloat) == 458.9 {
		fmt.Println("Yes | Array Max is set successfully for float(64)")
	} else {
		t.Error("No | Array Max isn't set successfully for float(64)")
	}
	if lib.Max(arrString) == "three" {
		fmt.Println("Yes | Array Max is set successfully for string")
	} else {
		t.Error("No | Array Max isn't set successfully for string")
	}
}

// TestMin test the function Min with array of type int, float(64) and string
func TestMin(t *testing.T) {
	arrInt := []any{1, 20, 300, 44, 56}
	arrFloat := []any{100.1, 20.2, 3.346598, 458.9, 5.5}
	arrString := []any{"one", "two", "three", "four", "five"}

	if lib.Min(arrInt) == 1 {
		fmt.Println("Yes | Array Min is set successfully for int")
	} else {
		t.Error("No | Array Min isn't set successfully for int")
	}
	if lib.Min(arrFloat) == 3.346598 {
		fmt.Println("Yes | Array Min is set successfully for float(64)")
	} else {
		t.Error("No | Array Min isn't set successfully for float(64)")
	}
	if lib.Min(arrString) == "one" || lib.Min(arrString) == "two" {
		fmt.Println("Yes | Array Min is set successfully for string")
	} else {
		t.Error("No | Array Min isn't set successfully for string")
	}
}

// todo
func TestSlice(t *testing.T) {

}
