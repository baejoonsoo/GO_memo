// 컴파일이 필요할 경우 main package가 존재해야 한다
package main

import (
	"fmt"
	"strings"
)


func lenAndUpper(name string) (lenght int,uppercase string){
	lenght= len(name)
	uppercase = strings.ToUpper(name)

	return
}

func nakedReturn()(name string){
	name="jun"
	return
}


func main() {
	totalLenght,upperName:=lenAndUpper("goLang")

	fmt.Println(totalLenght,upperName)
	fmt.Println(lenAndUpper("go"))

	fmt.Println(nakedReturn())
}