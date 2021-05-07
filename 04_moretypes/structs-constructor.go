package main

import "fmt"

type mapStruct struct {
	data map[int]string
}

//맵 형태의 data필드를 가지는 "mapStruct"를 정의함
func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}          //구조체 객체를 생성하고 초기화함
	d.data = map[int]string{} //data 필드의 맵을 초기화함
	return &d                 //초기화 한 포인터 구조체를 반환함
}

func main() {
	s1 := newStruct() // 생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}
	s2.data = map[int]string{}
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)
}

// 구조체의 필드 자체가 사용 전에 초기화되어야 하는 경우가 있다.
//
