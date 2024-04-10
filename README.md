# Array library

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/array)](https://goreportcard.com/report/github.com/Eclalang/array)
[![codecov](https://codecov.io/gh/Eclalang/array/graph/badge.svg?token=YNCIYERVBO)](https://codecov.io/gh/Eclalang/array)

## Candidate functions :

| Func Name |                              Prototype                              |                                          Description                                          | Comments |
|:---------:|:-------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------:|:--------:|
| Contains  |       func Contain[T comparable](array []T, value T) bool {}        |                               Test if the value is in the array                               |   N/A    |
|   Find    |         func Find[T comparable](array []T, value T) int {}          | Returns the index of the value in the array <br/> (return -1 if the value isn't in the array) |   N/A    |
|  IsEqual  |   func IsEqual[T comparable](FirstArray, SecondArray []T) bool {}   |                          Test two array and check if they are equals                          |   N/A    |
|    Max    |         func Max[T cmp.Ordered](array []T) (any, error) {}          |                            Returns the maximum value in the array                             |   N/A    |
|    Min    |         func Min[T cmp.Ordered](array []T) (any, error) {}          |                            Returns the minimum value in the array                             |   N/A    |
|  Remove   |       func Remove[T comparable](array []T, index int) []T {}        |                          Removes the value of the array at the index                          |   N/A    |
|   Slice   | func Slice[T comparable](array []T, start, end int) ([]T, error) {} |                              Slices the array from start to end                               |   N/A    |
|   Sort    |          func Sort[T cmp.Ordered](array []T, order int) {}          |                       Sorts the array in ascending or descending order                        |   N/A    |

## Constants :
| Name | Type | Value |
|:----:|:----:|:-----:|
| ASC  | int  |   0   |
| DESC | int  |   1   |