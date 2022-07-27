/**
 * @Time : 2022/7/27 6:08 下午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

package string

import (
	"fmt"
)

func ExampleSplitByRegexp() {
	result, _ := SplitByRegexp("hello1world", "[0-9]+")
	fmt.Println(result)
	result1, _ := SplitByRegexp("111111hello222222world333333", "\\d{6}")
	fmt.Println(result1)
	// Output:
	// [hello world]
	// [ hello world ]
}

func ExampleTruncate()  {
	fmt.Println(Truncate("hello world!", "", 11))
	fmt.Println(Truncate("hello world!", "", 6))
	// Output:
	// hello world...
	// hello ...
}