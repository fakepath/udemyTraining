package main
import (
	"fmt"
)
func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // *int var type = a pointer to an int - pointing to a memory address where an int is stored, usefull when working with functions
	fmt.Println(b)
}
