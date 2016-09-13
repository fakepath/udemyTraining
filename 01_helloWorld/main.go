package main

import "fmt"

func main() {
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
