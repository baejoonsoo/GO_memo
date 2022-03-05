package something

import "fmt"

// go언어에서 대문자로 시작하는 함수는 자동으로 export 된다
func SayHello(){
	fmt.Println("hello world!")
}