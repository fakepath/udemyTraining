package main

import "fmt"

const p string = "death and taxes"
const (
	Pi       = 3.14
	Language = "Go"
)
const (
	A = iota //0
	B = iota //1
	C = iota //2
)

//resets the iota value
//const (
//	D = iota //0
//	E = iota //1
//	F = iota //2
//)

const (
	_  = iota             // throw away the zero
	KB = 1 << (iota * 10) // 1*10
	MB = 1 << (iota * 10) //2*10
)

func main() {
	const q int = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(Pi)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
