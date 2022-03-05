// 컴파일이 필요할 경우 main package가 존재해야 한다
package main

import "fmt"

// func add(numbers ...int)int{
// 	total:=0
// 	for _,number := range(numbers){
// 		total+=number
// 	}

// 	return total
// }

// total := add(1,2,3,4,5,6)

// fmt.Println(total)
func main(){
	numbers := [...]int{10,20,30,40,50}

	for i:=0; i<len(numbers); i++{
		fmt.Println(numbers[i])
	}
}