// 컴파일이 필요할 경우 main package가 존재해야 한다
package main

import "fmt"

func main(){
	myMap := map[string]int{"class":4, "grade":2, "number":5}

	for key,value :=range myMap{
		fmt.Printf("key : %s, value : %d\n",key,value)
	}
}