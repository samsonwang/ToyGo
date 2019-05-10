
package main

import "fmt"

// 接收变长参数列表
func f (strs ...string) {
	fmt.Println(len(strs))
}

func main() {
	s1 := []string{"a", "b", "c"}

	// unpack operator
	// expand operator
	// 用于将slice中的元素释放出来
	f(s1...)

//	var s1_unpack = (s1...)

	// 以下写法编译不通过
	// fmt.Println(s1...)
}
