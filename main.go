// 컴파일이 필요할 경우 main package가 존재해야 한다
package main

import (
	"fmt"
)

func deferFunc(){
	fmt.Println("defer func!!")
}

func nakedReturn()(name string){
	defer deferFunc()
	name="jun"
	fmt.Println("nakedReturn func")
	return
}

func main() {
	fmt.Println(nakedReturn())
}