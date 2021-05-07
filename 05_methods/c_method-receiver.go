package main

import "fmt"

type Point3 struct {
	X, Y int
}

func (p *Point3) add(a int) {
	p.X += a
	p.Y += a
	fmt.Printf("add() p확인: %p\n", p)
	fmt.Printf("add의 p값 확인: %v\n", p)
}

func (p Point3) mul(a int) {
	p.X *= a
	p.Y *= a
	fmt.Printf("mul의 p확인: %p\n", &p)
	fmt.Printf("mul의 p값 확인: %v\n", p)
}

func main() {
	p := Point3{3, 4}
	fmt.Printf("p의 포인터: %p\n", &p)
	fmt.Printf("p의 값: %v\n", p)

	// 위의 있는 main의 p와 add의 p는 같다.
	// 그러나 mul의 p는 다른데,
	// 이 이유는 add는 main의 p를 대변하기 때문이다.
	// 이는 포인터이기 때문에 add의 p는 main p 그 자체.
	// mul의 p는 값이 복제된 것이기 때문에
	// 동등한 값이 나오지 않는다.
	// 메소드 안에서 값을 바꾸고 싶다면 포인터 리시버 사용,
	// 그게 아니라면 value 리시버 사용
	p.add(10)
	p.mul(100)

	fmt.Printf("p의 값: %v\n", p)
}

// https://kamang-it.tistory.com/entry/Go15%EB%A9%94%EC%86%8C%EB%93%9CMethod%EC%99%80-%EB%A6%AC%EC%8B%9C%EB%B2%84Receiver
