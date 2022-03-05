# GO_Scraper

## function(함수)

keyword: func

- go언어는 main package를 우선으로 실행

- 함수의 이름가 대문자로 시작할 시 해당 함수를 export된다

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
  - 함수 내부에서만 축약형으로 선언 가능하디
  - <b style="color:pink">var의 축약형</b>으로 수정이 가능한 변수가 선언된다
