/**
 * @Time : 2022/7/22 9:54 上午
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

package string

import (
	"regexp"
)

// SplitByRegexp creates an array split by regexp.
func SplitByRegexp(str, expr string) (result []string, err error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return
	}
	return r.Split(str, -1), nil
}

// Truncate string if it's longer than the given maximum string length. The last characters of the truncated string are replaced with the omission string which defaults to "...".
func Truncate(str, omission string, max int) string {
	if len(str) <= max {
		return str
	}
	if omission == "" {
		omission = "..."
	}
	return str[0:max] + omission
}