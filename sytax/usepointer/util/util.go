package util

import (
	// "bytes"
	"fmt"
	"unsafe"
)

type student struct {
	name string
	age  int
}

func Test1(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Test2(u uint64) float64 {
	return *(*float64)(unsafe.Pointer(&u))
}

func Test3() {
	var i = 1
	fmt.Println(unsafe.Pointer(&i))
	fmt.Printf("0x%x", uintptr(unsafe.Pointer(&i)))
}

//	func Test4() {
//		var n int64 = 1
//		address := uintptr(unsafe.Pointer(&n))
//		np := *(*uint64)(unsafe.Pointer(address))
//		fmt.Println(*np)
//	}
func Test4() {
	p := &student{
		name: "tom",
		age:  18,
	}
	pp := unsafe.Pointer(p)
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(p.name))))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(p.age))))
	var slice = []int{1, 2, 3, 4, 5, 6, 7}
	s := unsafe.Pointer(&slice[0])
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(s) + 8)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(s) + 16)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(s) + 24)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(s) + 32)))
	fmt.Println(*(*int)(unsafe.Pointer(uintptr(s) + 40)))
	fmt.Println(*(*int)(unsafe.Add(s, 48)))
}

func Test5() {
	var slice = []byte("hello,world!")
	str := unsafe.String(unsafe.SliceData(slice), len(slice))
	fmt.Println(str)
}
