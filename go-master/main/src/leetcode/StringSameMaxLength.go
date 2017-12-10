package main

import "fmt"
import "strings"

/**
* Created by wangqing on 17/12/9.
*
* 输入两个字符串a和b，输出最长公共字串的长度
* eg：a="abcd",b="cbce" result=2
*/
func main() {
	a := "abcde"
	b := "bfbcde"
	result := searchSameStrLength(a, b)
	fmt.Println("result=", result)

}

func searchSameStrLength(a string, b string) int {
	max := 0
	length := len(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			part := a[i:j]
			if (strings.Contains(b, part)) {
				partLen := len(part)
				if (partLen > max) {
					max = partLen
				}
			}
		}
	}
	return max
}
