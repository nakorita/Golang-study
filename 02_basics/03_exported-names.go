package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

// 기초 - 패키지와 변수,함수 - Export되는 이름
// Go에서는 대문자로 시작하는 이름이 export된다.
// ex) 위의 코드를 보면 Pi가 math 패키지에서 export된 걸 확인 가능.
// fmt.Println(math.pi) <- 옆에 pi를 대문자로 시작하지 않으면
// export되지 않는다.
// export되지 않은 이름들은 패키지의 밖에서 접근할 수 없다.
