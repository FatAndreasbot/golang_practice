```go
package main

import "fmt"

type MyStruct struct {
	MyInt int
}

func func1() MyStruct {
	return MyStruct{MyInt: 1}
}

func func2() *MyStruct {
	return &MyStruct{}
}

func func3(s *MyStruct) {
	s.MyInt = 333
}

func func4(s MyStruct) {
	s.MyInt = 923
}

func func5() *MyStruct {
	return nil
}

func main() {
    ms1 := func1()
    fmt.Println(ms1.MyInt) // 1 - просто возврат поля структуры

    ms2 := func2()
    fmt.Println(ms2.MyInt) // 0 - поле пустой структуры равно 0

    func3(ms2)
    fmt.Println(ms2.MyInt) // 333 - передаем структуру по ссылке. меняем значение

    func4(ms1)
    fmt.Println(ms1.MyInt) // 333 - передаем по значению. оригинальный инстанс не мутируется

    ms5 := func5()
    fmt.Println(ms5.MyInt) // panic - null pointer dereference
}
```
