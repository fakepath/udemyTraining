package main
//packages/variables/scope
import (
	"fmt" //file level scope
	"github.com/udemyTraining/stringutil"
	"github.com/udemyTraining/vis"
)

var x int = 42 //package level scope, is enclosing the inner scope of the bellow function
func max(x int) int {
	return 42 + x
}
func main() {
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	for i := 60; i < 70; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	//--------------
	fmt.Println(stringutil.Reverse("!oG, olleH"))
	fmt.Println(stringutil.MyName)
	//----------------
	a := 10       //int
	b := "golang" //string
	c := 4.17     //float64
	d := true     //bool

	//%v the value in default format godoc.org
	//%T type of the value
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	name := "Virgil" //declare and initialize at the same time
	fmt.Println("My name is", name)

	fmt.Println(x)
	foo()
	//----------------------------------------------
	y := 17 //block level scope
	fmt.Println(y)
	//----------------------------------------------
	max := max(7)    //variable shadowing -> never do it
	fmt.Println(max) // the result, not the function
	//max(3) - can't be accesed anymore, bcuz of variable shadowing

	//----------------------------------------------
	fmt.Println(vis.MyName) //variable with upperCase->exported outside the package, lowercase->unexported

	vis.PrintVar() //function exported, using unexported variable
	//	-------------------
	v := 10
	fmt.Println(v)
	{
		fmt.Println(v + 1)
		w := "this is the inner scope variable"
		fmt.Println(w)
	}
	//fmt.Println(w) - outside scope of w
	//----------------------increment func
	fmt.Println(increment())
	fmt.Println(increment())
	//--------------------- OR
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	//---------------------------
	increment1 := wrapper()
	fmt.Println(increment1())
	fmt.Println(increment1())

}
func foo() {
	fmt.Println(x + 1) //block level scope
}
func increment() int {
	x++
	return x
}
func wrapper( /*params*/ ) func( /*the return of the func*/ ) int {
	x := 0
	return func() int {
		x++
		return x
	}
}
