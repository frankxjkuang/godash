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

func ExampleDifference() {
	fmt.Println(Difference([]any{3, 2, 1}, []any{4, 2}))
	fmt.Println(Difference([]any{3, 2, 1, "a", false, "c"}, []any{2, "c"}))
	// Output:
	// [3 1]
	// [3 1 a false]
}

func ExampleDrop() {
	fmt.Println(Drop([]any{1, 2, 3}, 0))
	fmt.Println(Drop([]any{1, 2, 3}, 2))
	fmt.Println(Drop([]any{1, 2, 3}, 3))
	fmt.Println(Drop([]any{1, 2, 3}, 5))
	// Output:
	// [1 2 3]
	// [3]
	// []
	// []
}

func ExampleDropRight() {
	fmt.Println(DropRight([]any{1, 2, 3}, 0))
	fmt.Println(DropRight([]any{1, 2, 3}, 2))
	fmt.Println(DropRight([]any{1, 2, 3}, 3))
	fmt.Println(DropRight([]any{1, 2, 3}, 5))
	// Output:
	// [1 2 3]
	// [1]
	// []
	// []
}

func ExampleIndexOf() {
	fmt.Println(IndexOf([]any{1, 2, 1, 2}, 2, 0))
	fmt.Println(IndexOf([]any{1, 2, 1, 2}, 2, 2))
	fmt.Println(IndexOf([]any{1, 2, 1, 2}, 2, -1))
	fmt.Println(IndexOf([]any{1, 2, 1, 2}, 0, -2))
	fmt.Println(IndexOf([]any{1, 2, 1, 2}, 0, 5))
	// Output:
	// 1
	// 3
	// 3
	// -1
	// -1
}
