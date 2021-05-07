package main

import "fmt"

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
	bit1, mask1                          // bit1 == 2, mask1 == 1
	_, _                                 // iota == 2를 생략하여 bit2와 mask2 생략
	bit3, mask3                          // bit3 == 8, mask3 == 7
)

const (
	zero = iota // 0
	one  = iota // 1
)
const (
	two   = iota // 0
	three = iota // 1
)

const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(bit0, mask0, bit1, mask1, bit3, mask3)
	fmt.Println(zero, one, two, three)
	fmt.Println(KB, MB, GB, TB)
}

// 상수는 보통 연속되는 값을 가진다.
// const 키워드와 ()를 사용하면 const키워드 반복할 필요 X
// 상수에 값을 일일이 대입하지 않고,
// 순서대로 생성하려면 iota를 사용하면 된다.
// iota는 데이터 타입이 없는 정수를 연속으로 생성해준다.
// iota에 연산자와 특정 값을 조합하면
// 상수값을 0, 1, 2, 3처럼 하지 않고 특정한 순서로 생성 가능.
// ***특정 iota값을 생략하려면 _을 사용하면 된다.
// iota는 위에 상수 선언부를 보면 iota를 생략해서도 사용 가능.
// 한번 사용하면 그 다음 상수에도 자동으로 적용된다.
// iota는 상수마다 별도로 적용된다.
// 상수를 괄호를 써서 여러개 선언할 때엔
// 첫 번째 값은 무조건 초기화해야 한다.
