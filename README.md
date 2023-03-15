# array


| Name | Prototype | Description |
| --- | --- | --- |
| NewArray | func NewArray() *Array | Function that initializes the array |
| Add | func (a *Array) Add(value interface{}) | Function that appends a value to the array |
| Get | func (a *Array) Get(index int) interface{} | Function that returns the value at the index |
| Set | func (a *Array) Set(index int, value interface{}) | Function that sets the value at the index |
| Length | func (a *Array) Length() int | Function that returns the length of the array |
| Remove | func (a *Array) Remove(index int) | Function that removes the value at the index |
| Print | func (a *Array) Print() | Function that prints the array |
| Map | func (a *Array) Map(fn func(interface{}) interface{}) *Array | Func that maps over the array and returns a new array with the result of the function fn |
| Filter | func (a *Array) Filter(fn func(interface{}) bool) *Array | Func that filters the array and returns a new array with the filtered values of the function fn |
| Reduce | func (a *Array) Reduce(fn func(interface{}, interface{}) interface{}, initialValue interface{}) interface{} | Func that reduces the array and returns a new array with the reduced values of the function fn |
| Sort | func (a *Array) Sort(comparator func(interface{}, interface{}) bool) | Func that sorts the array, this sort depends on the comparator function |
| BinarySearch | func (a *Array) BinarySearch(value interface{}) int | Function that returns the index of the value in the array using binary search |
| SortAsc | func (a *Array) SortAsc() | Function that sorts the array in ascending order |
| SortDesc | func (a *Array) SortDesc() | Function that sorts the array in descending order |
| Max | func (a *Array) Max() interface{} | Function that returns the maximum value in the array |
| Min | func (a *Array) Min() interface{} | Function that returns the minimum value in the array |
| Slice | func (a *Array) Slice(start, end int) *Array | Function that slices the array from start to end |
