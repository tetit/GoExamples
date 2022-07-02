package main

import (
	"fmt"
	"unsafe"
)

type Empty struct{}

var a struct{}
var x Empty

type NeedInterface interface {
	MethodA()
}

func (Empty) MethodA(){
	fmt.Println("method A")
}
 type myInt int32

 func (myInt) MethodB(){
	fmt.Println("method B")
}

func main() {
	var i int32
	var s string
	var b bool

	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(x))

	c := struct{}{}
	fmt.Println(unsafe.Sizeof(&a))
	fmt.Println(unsafe.Sizeof(&c))
	fmt.Println(unsafe.Sizeof(&b))
	
	var l myInt
	l.MethodB()
	
	var e Empty
	e.MethodA()
	
	fmt.Println(unsafe.Sizeof(l))
	fmt.Println(unsafe.Sizeof(e))

}
