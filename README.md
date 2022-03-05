# GO_Scraper

## function(함수)

keyword: func

- go언어는 main package를 우선으로 실행

- 함수의 이름가 대문자로 시작할 시 해당 함수를 export된다

- 함수에 값을 받거나 return 할때는 그 값을 명시해야 한다

  ```go
  func hello(a <a type>, b <b type>) <return type>{
    return "go lang"
  }

  ex:
  func multiply(a int,b int) int{
  	return a*b
  }
  ```

  - array를 return 하는 경우 []를 type 앞에 작성한다

    <br>

- 함수의 parameter의 타입이 동일할 시 아래와 같이 작성할 수 있다

  ```go
  // (a int, b int) -> (a, b int)

   func multiply(a,b int) int{
  	return a*b
  }
  ```

- 여러개의 값을 return 할 수 있다

  ```go
  func lenAndSting(number string) (int,string){
  return len(name), name
  }
  ```

- naked return<br>
  go언어가 값을 return하는 또 다른 방식으로 return 값 타입 지정 시 해당 타입으로 return 될 변수를 만든 뒤 return 할 수 있다

  ```go
  func nakedReturn1()(name string){
    // 선언한 변수에 return 할 값을 저장한다
    name="jun"

    // return문 뒤에 변수를 붙히지 않은 채 return한다
    // return문 뒤에 생성한 변수명을 전부 입력한다면 에러는 나지 않음
    return
  }
  nakedReturn1 output: "jun"

  func nakedReturn2() (name string, age int){
    name = "junsu"
    age = 18

    return
  }
  nakedReturn2 output: "junsu 18"
  ```

- 무한한 갯수의 parameter들을 받을 수 있다

  ```go
  func lala(words ...int){
    fmt.Println(words)
  }

  func main(){
    lala(1,2,3,4,5)
  }

  output: [1 2 3 4 5]
  ```

  - <b style="color: pink">...</b> 키워드를 통해 무한한 갯수의 parameter를 받을 수 있으며 array로 저장된다

- defer
  함수가 return 됐을 때 defer 키워드 뒤에 입력한 함수를 실행한다

  ```go
  func deferFunc(){
    fmt.Println("defer func!!")
  }

  func goFunc()(text string){
    defer deferFunc()
    fmt.Println("goFunc")

    text="return string"
    return
  }

  func main() {
    fmt.Println(goFunc())
  }

  // |output|
  // goFunc
  // defer func!!
  // return string
  ```

<br>

## Variable & constants(변수와 상수)

keyword: var(변수), const(상수)

- 선언 방법

  ```go
  <var or const> <name> <type> = <value>

  ex:
  const name string = "junsu"
  var age int = 18
  ```

  <br>

- 축약형
  ```go
  name := "junsu"
  ```
  - 초기값의 타입에 맞춰 변수의 타입이 지정되며 타입이 변경되지 않는다
  - 함수 내부에서만 축약형으로 선언 가능하다
  - <b style="color:pink">var의 축약형</b>으로 수정이 가능한 변수가 선언된다

## 반복문

- forEach, map 등의 배열 함수 없이 for만 존재
- for문

  ```go
  func main(){
    numbers := [...]int{10,20,30,40,50}

    for i := 0; i < len(numbers); i++{
      fmt.Println(numbers[i])
    }
  }

  // output
  // 10
  // 20
  // 30
  // 40
  // 50
  ```

- array를 도는 for문

  range를 이용하여 배열의 요소들을 순회할 수 있다

  ```go
  func main() {
  numbers := [...]int{10,20,30,40,50}

  	for index, number := range(numbers){
  		fmt.Println("index: ",index, "number: ",number)
  	}
  }

  // output
  // index:  0 number:  10
  // index:  1 number:  20
  // index:  2 number:  30
  // index:  3 number:  40
  // index:  4 number:  50
  ```

## if else

- 기본 문법은 c언어와 유사

  ```go
  if a>10{
    // ...code
  }else{
    // ...code
  }
  ```

<br>

- if else문 내에서만 사용할 변수를 선언 할 수 있다<br>
  <b>variable expression</b>

  ```go
  number:=10
  if newNumber := number+10; newNumber < 30 {
    fmt.Println("if", newNumber)
  } else{
    fmt.Println("else", newNumber)
  }
  ```

## switch

- variable expression 사용 가능
- 소괄호 제거
- 이하생략
