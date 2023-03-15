package arraylibtester

import (
	"fmt"
	"testing"

	list "github.com/Eclalang/array"
)

func TestArray(t *testing.T) {
	arr := list.NewArray()
	if arr == nil {
		t.Error("Array is nil")
	}
	arr.Add(1)
	arr.Add(28)
	arr.Add(0)
	arr.Add(8)
	arr.Add(5)
	arr.Add(21)
	arr.Add(7)
	arr.Add(10)
	fmt.Printf("arr : ")
	arr.Print()
	if arr.Length() != 8 {
		t.Error("Array length is not 3")
	} else {
		fmt.Println("Array length is set successfully")
	}
	if arr.Get(0) != 1 {
		t.Error("Array[0] is not 1")
	} else {
		fmt.Println("Array[0] is set successfully")
	}
	if arr.Get(1) != 28 {
		t.Error("Array[1] is not 2")
	} else {
		fmt.Println("Array[1] is set successfully")
	}
	if arr.Get(2) != 0 {
		t.Error("Array[2] is not 3")
	} else {
		fmt.Println("Array[2] is set successfully")
	}
	arr.Set(1, 4)
	if arr.Get(1) != 4 {
		t.Error("Array[1] is not 4")
	} else {
		fmt.Println("Array[1] is set successfully")
	}
	arr.Remove(2)
	if arr.Get(2) == 28 {
		t.Error("Array[2] is not removed")
	} else {
		fmt.Println("Array[2] is removed successfully")
	}
	fmt.Printf("arr : ")
	arr.Print()
	fmt.Println("arr length :", arr.Length())
	if arr.Length() != 7 {
		t.Error("Array length is not 7")
	} else {
		fmt.Println("Array length is set successfully")
	}
	arr2 := arr.Map(func(v interface{}) interface{} {
		return v.(int) * 5
	})
	for i := 0; i < arr2.Length(); i++ {
		if arr2.Get(i) != arr.Get(i).(int)*5 {
			t.Error("Function Map is not working properly")
		}
	}
	fmt.Println("Function Map is working successfully")
	arr3 := arr.Filter(func(v interface{}) bool {
		return v.(int) > 2
	})
	fmt.Printf("arr3 : ")
	arr3.Print()
	for i := 0; i < arr3.Length(); i++ {
		if arr3.Get(i).(int) < 2 {
			t.Error("Function Filter is not working properly")
		}
	}
	fmt.Println("Function Filter is working successfully")
	arr4 := arr.Reduce(func(v interface{}, acc interface{}) interface{} {
		return v.(int) + acc.(int)
	}, 0)
	sum := 0
	for i := 0; i < arr.Length(); i++ {
		sum += arr.Get(i).(int)
	}
	if arr4.(int) != sum {
		t.Error("Array4 is not the sum of Array")
	} else {
		fmt.Println("Function Reduce is working successfully")
	}
	arr.Sort(func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	})
	for i := 0; i < arr.Length()-1; i++ {
		if arr.Get(i).(int) > arr.Get(i+1).(int) {
			t.Error("Array is not sorted")
		}
	}
	fmt.Println("Function Sort is working successfully")
	fmt.Printf("arr : ")
	arr.Print()
}
