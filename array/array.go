/**
 * @Time : 2022/6/28 6:10 下午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

// “Array” Methods
package array

// Chunk Creates an array of elements split into groups the length of `size`.
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

// Concat Creates a new array concatenating array with any additional arrays and/or values.
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

// Difference Creates an array of array values not included in the other given arrays.
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
