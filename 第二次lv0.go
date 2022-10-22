package main

import (
	"fmt"
	"reflect"
)

type wxjj int

type wxgg = int

func main() {
	var a wxjj

	var b wxgg

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	rfTypeOf(a)
	rfTypeOf(b)

	TypeOf(a)
	TypeOf(b)
}

func rfTypeOf(data interface{}) {
	of := reflect.TypeOf(data)
	fmt.Println(of)
}

func TypeOf(data interface{}) {
	switch data.(type) {
	case wxgg:
		fmt.Println("Type is int")
	case nil:
		fmt.Println("Type is nil")
	default:
		fmt.Println("Type Not Found")
	}
}


type WxGG struct {
	Name string
	Age  int
}

func main() {

	//最常见的方式
	a := WxGG{
		Name: "wxgg1",
		Age:  18,
	}

	var b WxGG
	b.Name = "wxgg2"
	b.Age = 18


	c:=NewWxGG("wxgg tql",18)


	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", WxJJ)
	fmt.Printf("%#v\n", c)

}
func NewWxGG(name string, age int) *WxGG {
	return &WxGG{
		Name: name,
		Age:  age,
	}
}
Wxjj := struct {
Name string
Age int
}{
"wxjj2",
  18,
}

fmt.Println(wxjj)
fmt.Println(wxjj.Name)
fmt.Println(wxjj.Age)
type WxGG struct {
	Name  string
	Age   int
	Books []Book
}

type Book struct {
	Name string
}

func (w WxGG) PrintName() {
	fmt.Println(w.Name)
}
func (w WxGG) PrintAge() {
	fmt.Println(w.Age)
}

func (w WxGG) PrintBook() {
	fmt.Println(w.Books)
}

func (b Book) PrintBookName() {
	fmt.Println(b.Name)
}
func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)
}
func main() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}
type Sayer interface {
	Say()
}

type dog struct {
}

type cat struct {
}

func (d dog) Say() {
	fmt.Println("汪汪汪")
}

func (c cat) Say() {
	fmt.Println("喵喵喵")
}
func main() {
	var x Sayer
	a := cat{}
	b := dog{}
	x = a
	x.Say()
	x = b
	x.Say()
}
func TypeOf(data interface{}) {
	switch data.(type) {
	case wxgg:
		fmt.Println("Type is int")
	case nil:
		fmt.Println("Type is nil")
	default:
		fmt.Println("Type Not Found")
	}
}

func main() {
	var x interface{}
	s := "YuanShen"
	x = s
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}