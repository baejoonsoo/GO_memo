// 컴파일이 필요할 경우 main package가 존재해야 한다
package main

import "fmt"

func lala(words ...int) []int{
	return words
}

func main(){
	fmt.Println(lala(1,2,3,4,5))
}