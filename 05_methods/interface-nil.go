package main

import (
	"fmt"
	"reflect"
)

type A struct{}

var condition = false

func billo() interface{} {
	var a *A = nil
	if condition {
		a = &A{}
	}
	return a
}

func main() {
	// 밑에 코드 수행 결과 'false'가 출력된다.
	// billo()의 리턴 값 a는 값이 nil이지만, 타입은 *A이다.
	// 그러나 밑의 nil은 타입도 nil, 값도 nil이다.
	fmt.Println(billo() == nil)

	// 밑의 두 코드를 실행해보면 <nil>, *main.A
	// 가 나온다.
	// 이런 값을 typed nil이라고 불린다.
	fmt.Println("value확인: ", reflect.ValueOf(billo()))
	fmt.Println("type확인: ", reflect.TypeOf(billo()))
}

// why?
// interface는 2가지 특징이 있다.
// 타입을 저장하는 T와 값을 저장하는 V 두 가지로 나뉘어 저장되는데,
// 	예를 들면 "hello" 라는 string value를 interface에 저장하면,
//  T에는 string이 V에는 "hello"가 들어간다.
// 인터페이스의 내부에 저장된 타입과 실제 데이터의 타입은
// 모두 같아야만 한다.
// https://blog.billo.io/devposts/typed_nil_go/
