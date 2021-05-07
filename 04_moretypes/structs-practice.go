package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000000000"
	fmt.Println(p1)

	// 구조체 선언 및 초기화 방법1
	p2 := person{"nam", 31, "01022220000"} // 필드 이름을 생력할 시 순서대로 저장함
	fmt.Println(p2)

	// 구조체 선언 및 초기화 방법2
	p3 := person{contact: "01011110000", name: "park", age: 23} // 필드 이름을 명시할 시 순서와 상관 없이 저장할 수 있음
	fmt.Println(p3)

	p3.name = "ryu" //필드에 저장된 값을 수정할 수 있음
	fmt.Println(p3)

	fmt.Println(p3.contact) //필드 값의 개별 접근도 가능함
}

// 구조체를 선언과 동시에 초기화하는 방법은 총 두 가지.
// 1) 초기화할 때 값을 나열하면 구조체에 선언한 필드 순서대로 저장.
// 2) 필드 이름에 값을 지정한다면 순서에 상관없이 해당 필드에 값이 저장.
// Go언어의 구조체는 기본적으로 'mutable' 개체.
// 필드값이 변화할 경우 별도로 새 개체를 만들지 않고
// 해당 개체 메모리에서 직접 변경된다.
