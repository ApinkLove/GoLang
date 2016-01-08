package main

import (
	"bufio"
	//	. "fmt" : 이런식으로 하면 fmt를 생략할수 있으나 나중에 중복될수 있으니 비추
	//	l "fmt" : 이런 식으로 하면 fmt대신 l로 쓰면 됩니다. 개인적으로 이걸 추천
	"fmt"
	_ "math" // 사용하지 않을때는 이런 식으로 해주세요
	"os"
	"runtime"
)

func main() {
	fmt.Println("CPU Count :", runtime.NumCPU()) // 자바처럼 + 안쓰네
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
