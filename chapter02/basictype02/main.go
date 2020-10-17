package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b=%v  type=%T\n", b, b)

	var str2 = "131455"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 0, 64)
	fmt.Printf("n1=%v  type=%T", n1, n1)
}
