package main

import (
	"fmt"
	"math/rand"
)

func Main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

// 기초 - 패키지와 변수,함수 - 패키지
// 프로그램은 main 패키지에서 실행을 시작.
// 패키지의 이름은 import 경로의 마지막 요소와 같다.
// ex) "math/rand" 패키지는 package rand문으로 시작하는
// 파일들로 구성되어 있다.
// cf) 위의 예제 프로그램 실행할 때 마다 같은 숫자 반환.
// 		이 프로그램이 실행되는 환경은 결정되어있으므로.
