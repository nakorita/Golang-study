package main

import (
	"fmt"
	"math"
)

type vertex4 struct {
	X, Y float64
}

// func (매개변수이름 구조체이름) 메소드이름() 반환형
// 여기서 매개변수이름은 구조체 변수명 -> 메소드 내에서 매개변수처럼 사용함.
// '함수이름'을 입력하지 않고 구조체 이름 뒤에 메소드 이름 입력
// 이 이름은 본문에서 메소드를 이용하기 위해 사용

// (v vertex4)는 어떤 구조체를 전달 받는지를 명시하는 '리시버'
// 구조체 객체 자체를 전달받는 게 아니라,
// ***구조체 객체 정보를 전달받고 메소드의 기능을 수행하는 것.
// ex) triarea := tri.triArea()

// 물론 메소드에서도 값을 복사해서 받는 게 아닌 포인터 리시버도 있음.
// 아래 코드는 '값' 정보를 전달 받아 연산한 후 반환.
func (v vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := vertex4{3, 4}
	fmt.Println(v.Abs())

	r := vertex4{2, 4}
	fmt.Println(r.Abs())
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 메소드
// Go는 클래스를 가지지 않는다.
// 하지만, 그와 같은 타입의 메소드 정의 가능.
// 그 메서드는 receiver 인자가 있는 함수이다.
// receiver는 func키워드와 메서드 이름사이에 위치
// 자체 인수 목록에서 나타난다.
// 위의 예제에서는 Abs 메소드에는 v라는 이름의
// vertex 유형의 리시버가 있다.
