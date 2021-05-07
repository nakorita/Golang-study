package main

import "fmt"

func main() {
	m := make(map[string]int)

	// 요소 추가하기
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 58
	fmt.Println("The value: ", m["Answer"])

	// 요소 삭제하기
	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	// 요소 검색하기
	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)

	// 요소 업데이트
	m["Answer"] = 11
	fmt.Println("The value: ", m["Answer"])

	// 두 개의 값 할당 후 키 존재 여부 테스트
	v, ok = m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
}

// 기초 - 더 많은 타입들 - Mutating Maps
// map에 요소 추가 or 업데이트
// ex) 맵 변수이름[key 이름] = 값
// 요소 검색하기
// ex) 맵의 값 = 맵 변수이름[key 이름]
// 요소 삭제하기
// ex) delete(맵 변수이름, key이름)
// key 존재하는지 테스트
// ex) 값 받을 변수 이름, true/false 리턴받을 변수 이름 = 맵이름[key 이름]
// 		-> key가 맵 변수 안에 있다면 true, 아니면 false
// 만약 key가 맵 안에 없다면, 요소는 zero value이다. ("")
// ***만약 요소 또는 true/false를 리턴받을 변수가
// 선언되지 않았다면, 아래와 같은 짧은 정의식 사용
// ex) elem, ok := m[key]
