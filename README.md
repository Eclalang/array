# Array library

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/array)](https://goreportcard.com/report/github.com/Eclalang/array)
[![codecov](https://codecov.io/gh/Eclalang/array/graph/badge.svg?token=YNCIYERVBO)](https://codecov.io/gh/Eclalang/array)

## Candidate functions :

| Func Name |                         Prototype                          |                                          Description                                          | Comments |
|:---------:|:----------------------------------------------------------:|:---------------------------------------------------------------------------------------------:|:--------:|
|  Contain  |    func Contain(array any, value any) (bool, error) {}     |                               Test if the value is in the array                               |   N/A    |
|   Find    |      func Find(array any, value any) (int, error) {}       | Returns the index of the value in the array <br/> (return -1 if the value isn't in the array) |   N/A    |
|  IsEqual  | func IsEqual(FirstArray, SecondArray any) (bool, error) {} |                          Test two array and check if they are equals                          |   N/A    |
|    Max    |            func Max(array any) (any, error) {}             |                            Returns the maximum value in the array                             |   N/A    |
|    Min    |            func Min(array any) (any, error) {}             |                            Returns the minimum value in the array                             |   N/A    |
|  Remove   |          func Remove(array any, index int) any {}          |                          Removes the value of the array at the index                          |   N/A    |
|   Slice   |   func Slice(array any, start, end int) (any, error) {}    |                              Slices the array from start to end                               |   N/A    |
|  SortAsc  |                 func SortAsc(array any) {}                 |                              Sorts the array in ascending order                               |   N/A    |
| SortDesc  |                func SortDesc(array any) {}                 |                              Sorts the array in descending order                              |   N/A    |
