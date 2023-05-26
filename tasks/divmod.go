package main

import "fmt"

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

/*This function will divide the int a and b.
The result of this division will be stored in the int pointed by div.
The remainder of this division will be stored in the int pointed by mod.*/
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b

}
