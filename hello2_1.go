package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
)

func main() {
	var num1 complex64 = 1 + 2i
	var r1 float32 = real(num1) // 실수부분
	var i1 float32 = imag(num1) // 허수부분
	fmt.Println(r1)
	fmt.Println(i1)
	var num2 byte = 'a'
	//	var num3 byte = "a" : 큰 따옴표를 사용하시면 아니되옵니다.
	fmt.Println(num2) // 문자의 숫자값을 출력한다.
	var r4 rune = '너'
	_ = r4 // 사용안할땐 이래야 된다는 점~
	var r2 rune = '\U0000e55c'
	//	var r3 rune = '슬기' : 한글자여야만 , 작은 따옴표 여야만 한다!
	fmt.Println(r2)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
