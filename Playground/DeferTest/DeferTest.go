
package main

import "fmt"

var g_num int = 1

func funcTest() {
	fmt.Println("funcTest, g_num", g_num)
}

func main() {
	num := 10

	// will print out num as 10
	defer fmt.Println("world", num)

	num = 15

	fmt.Println("hello")

	// stacked defer
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	// defer function call
	// will print out g_num = -1
	defer funcTest()
	g_num = -1
}
