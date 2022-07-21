/**
 * @Time : 2022/6/28 6:10 下午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

// “Array” Methods
package array

// Chunk creates an array of elements split into groups the length of `size`.
// If `array` can't be split evenly, the final chunk will be the remaining elements.
func Chunk(array []any, size int) (result [][]any) {
	result = make([][]any, 0)
	if size == 0 || len(array) == 0 {
		return
	}
	if len(array) <= size {
		result = append(result, array)
		return
	}
	for i := 0; i < len(array); i += size {
		item := []any{}
		if i+size <= len(array) {
			item = append(item, array[i:i+size]...)
		} else {
			item = append(item, array[i:]...)
		}
		result = append(result, item)
	}
	return
}

// Concat creates a new array concatenating array with any additional arrays and/or values.
func Concat(array []any, args ...any) (result []any) {
	result = make([]any, len(array))
	_ = copy(result, array)
	for _, arg := range args {
		if a, ok := arg.([]any); ok {
			result = append(result, a...)
			continue
		}
		result = append(result, arg)
	}
	return result
}

// Difference creates an array of array values not included in the other given arrays.
// The order and references of result values are determined by the first array.
func Difference(array, values []any) (result []any) {
	result = make([]any, 0)
	keys := make(map[any]struct{})
	for _, v := range values {
		keys[v] = struct{}{}
	}
	for _, v := range array {
		if _, ok := keys[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// Drop creates a slice of array with n elements dropped from the beginning.
func Drop(array []any, num int) []any {
	if num >= len(array) {
		return []any{}
	}
	return array[num:]
}

// DropRight creates a slice of array with n elements dropped from the end.
func DropRight(array []any, num int) []any {
	if num >= len(array) {
		return []any{}
	}
	return array[0 : len(array)-num]
}

// DropRight gets the index at which the first occurrence of value is found in array using SameValueZero for equality comparisons. If fromIndex is negative, it's used as the offset from the end of array.
// Returns the index of the matched value, else -1.
func IndexOf(array []any, value, fromIndex int) int {
	if fromIndex > len(array)-1 || len(array) == 0 {
		return -1
	}
	if fromIndex >= 0 {
		for i, v := range array {
			if i >= fromIndex && v == value {
				return i
			}
		}
		return -1
	}
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// Pull removes all given values from array.
func Pull(array []any, a ...any) []any {
	result := make([]any, 0)
	kv := make(map[any]struct{})
	for _, v := range a {
		kv[v] = struct{}{}
	}
	for _, v := range array {
		if _, ok := kv[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

// PullBy removes all values from array when f return true.
func PullBy(array []any, f func(any) bool) []any {
	result := make([]any, 0)
	for _, v := range array {
		if !f(v) {
			result = append(result, v)
		}
	}
	return result
}

// PullAt removes elements from array corresponding to indexes and returns an array of removed elements.
func PullAt(array []any, idx ...int) []any {
	result := make([]any, 0)
	kv := make(map[int]struct{})
	for _, v := range idx {
		kv[v] = struct{}{}
	}
	for i, v := range array {
		if _, ok := kv[i]; !ok {
			result = append(result, v)
		}
	}
	return result
}
