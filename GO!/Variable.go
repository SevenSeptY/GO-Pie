package main

import "fmt"

func main() {
	var student1 string = "Qi Yao"
	student2 := "Felix Qi"
	x := 2
	var y int = 4

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(y)
}
//-----------------------------------

package main

import "fmt"

func main() {
	var a string
	var b bool
	var c int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
/*These variables are declared but they have not been assigned initial values.
By running the code, we can see that they already have the default values of their respective types:

a is ""
b is 0
c is false
*/
//---------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	var student1 string
	student1 = "John"

	var student2 string = "Snow"
	fmt.Print(student1 + " " + student2)
}

