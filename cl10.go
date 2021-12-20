package main

import "fmt"

func main() {
	/*var name [3]string
	name[0] = "Nicola"
	name[1] = "Murad"
	name[2] = "Nahid raince"
	fmt.Println(name)
	fmt.Println(name[1])
	*/
	//short way
	//... means joto khusi nite paro
	name := [...]string{"Takla Murad", "Toslima Nasrin", "Murgi kobir", "hasao mahmud"}
	fmt.Println(name)
}
