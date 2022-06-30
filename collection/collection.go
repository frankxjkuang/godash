/**
 * @Time : 2022/6/29 3:37 下午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

// “Collection” Methods
package collection

// CountBy Creates an object composed of keys generated from the results of running each element of collection thru iteratee.
// The corresponding value of each key is the number of times the key was returned by iteratee.
// The iteratee is invoked with one argument: (value).
// The element of arr must be hashable.
func CountBy(arr []any, f func(any) any) (result map[any]int) {
	result = make(map[any]int)
	for i := 0; i < len(arr); i++ {
		key := f(arr[i])
		result[key]++
	}
	return result
}
