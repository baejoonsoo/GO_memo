// 컴파일이 필요할 경우 main package가 존재해야 한다
package main

import "fmt"

type person struct{
	name string
	age int
	favFood []string
}

func main(){
	favFood := []string{"cake","chicken"}
  jun := person{"junsu",18,favFood}

  fmt.Println(jun)
}