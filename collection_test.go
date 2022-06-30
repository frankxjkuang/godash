/**
 * @Time : 2022/6/29 3:51 下午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

package godash

import (
	"fmt"
	"math"
)

func ExampleCountBy() {
	fmt.Println(CountBy([]any{6.1, 4.2, 6.3}, func(a any) any {
		return math.Floor(a.(float64))
	}))
	fmt.Println(CountBy([]any{"one", "two", "three"}, func(a any) any {
		return len(a.(string))
	}))
	// Output:
	// map[4:1 6:2]
	// map[3:2 5:1]
}
