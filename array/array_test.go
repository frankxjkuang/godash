/**
 * @Time : 2022/6/28 6:43 下午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

package array

import (
	"fmt"
)

func ExampleChunk() {
	fmt.Println(Chunk([]any{"a", "b", "c", "d"}, 2))
	fmt.Println(Chunk([]any{"a", "b", "c", "d"}, 3))
	// Output:
	// [[a b] [c d]]
	// [[a b c] [d]]
}

func ExampleConcat() {
	arr := []any{1}
	fmt.Println(Concat(arr, 2, []any{3}, [][]any{{4}}))
	fmt.Println(arr)
	// Output:
	// [1 2 3 [[4]]]
	// [1]
}
