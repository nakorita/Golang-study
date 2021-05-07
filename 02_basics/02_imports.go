package ny

import (
	"fmt"
	"math"
)

func test() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	test()
}

// 기초 - 패키지와 변수,함수 - import
// 괄호로 import를 그룹 짓는다.
// 이를 "factored" import문이라고 한다.
// 위의 방식 말고 import "fmt"
// 				 import "math"
// 이렇게 표현할 수 있다.
// 하지만 위의 방식보다는 위의 import문 사용이 good
