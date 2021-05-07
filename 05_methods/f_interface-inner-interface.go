package main

import "fmt"

type GetName interface {
	hisName()
}

type GetAge interface {
	hisAge()
}

// InfoPerson 인터페이스는
// GetName과 GetAge를 모두 가져온다.
// 이건 곧 InfoPerson은 GetName과 GetAge의
// 메소드를 모두 포함하게 된다.
// 즉, interface 내부에 interface를 선언하면
// 그 선언된 interface의 메소드를 가져오게 된다.
type InfoPerson interface {
	GetName
	GetAge
}

type Person struct {
	name string
	age  int
}

func (p Person) hisName() {
	fmt.Println(p.name)
	// fmt.Println("확인2: ", p)
}

func (p Person) hisAge() {
	fmt.Println(p.age)
}

func main() {
	var i InfoPerson
	i = Person{"Kukaro", 27}
	// fmt.Println("확인1: ", i)
	i.hisName()
	i.hisAge()

	var g GetName
	g = Person{"James", 15}
	g.hisName()
	// GetName형 인터페이스 객체이므로
	// GetName 인터페이스에는
	// hisAge()라는 메소드가 없어서
	// 구문 오류 발생
	// g.hisAge()
}
