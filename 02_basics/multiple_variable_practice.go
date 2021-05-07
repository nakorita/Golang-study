package main

import "fmt"

// 방법 2
var (
	a, b int = 3, 4
	c, d     = 5, 6
	// cdd  string = "Hi"
	// 위의 변수 cdd 이전에는 c로 선언했었음.
	// 노란줄 떠서 cdd로 변경해준건데,
	// 이는 10_short_variable_declarations.go에
	// 내가 이미 main() 내에 같은 이름으로 변수 c를 선언해주었기 때문에
	// 노란 줄 떴었음.
	// 이름이 main으로 같은 package내에는 같은 이름의 변수 선언 X
	// d int = 30
)

func main() {
	// 방법 1
	var x, y int = 10, 20
	// 방법 3
	e, f, g := 2.5, true, `vivian`
	fmt.Println(x, y)
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g)
}
