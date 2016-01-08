package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := 1 // 변수를 선언하면서 값까지 대입합니다.
	// 문자열은 더할순 있으나 빼기는 불가능합니다.
	/*
		리얼 c와 동일
	*/
	_ = a	
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
