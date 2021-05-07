package main

import "fmt"

type triangle struct {
	width, height float32
}

// value receiver는 전달받은 객체의 필드 값 변경해도
// 메소드 빠져나가면 값이 소멸되어 바뀌지 않음.
func (s triangle) triAreaVal() float32 { //Value Receiver
	s.width += 10
	return s.width * s.height / 2
}

// 포인터 정보를 전달해서 구조체 필드 값에 직접 접근
// 리시버 부분에 주솟값을 참조하는 연산자인 '*'를 구조체 이름 앞에 붙임.
// pointer receiver는 구조체 객체의 포인터를 전달받아 연산해서
// 객체의 실제 필드 값이 바뀜.
func (s *triangle) triAreaRef() float32 { //Pointer Reciver
	s.width += 10
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea_val := tri1.triAreaVal()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea_val)

	triarea_ref := tri1.triAreaRef()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea_ref)
}
