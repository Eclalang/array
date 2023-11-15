## ARRAY LIBRARY FOR ECLA

# Candidate functions :

| Func Name |                         Prototype                         |                 Description                 | Comments |
|:---------:|:---------------------------------------------------------:|:-------------------------------------------:|:--------:|
|  Append   |       func Append(array []any, value any) []any {}        |        Appends a value to the array         |   N/A    |
|  Length   |              func Length(array []any) int {}              |       Returns the length of the array       |   N/A    |
|  Remove   |       func Remove(array []any, index int) []any {}        | Removes the value of the array at the index |   N/A    |
|   Find    |         func Find(array []any, value any) int {}          | Returns the index of the value in the array |   N/A    |
|  Contain  |       func Contain(array []any, value any) bool {}        |      Test if the value is in the array      |          |
|  IsEqual  | func IsEqual(FirstArray []any, SecondArray []any) bool {} | Test two array and check if they are equals |   N/A    |
|  SortAsc  |               func SortAsc(array []any) {}                |     Sorts the array in ascending order      |   N/A    |
| SortDesc  |               func SortDesc(array []any) {}               |     Sorts the array in descending order     |   N/A    |
|    Max    |               func Max(array []any) any {}                |   Returns the maximum value in the array    |   N/A    |
|    Min    |               func Min(array []any) any {}                |   Returns the minimum value in the array    |   N/A    |
|   Slice   |      func Slice(array []any, start, end int) any {}       |     Slices the array from start to end      |   N/A    |
