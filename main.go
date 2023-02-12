package main

import "fmt"

func main(){
//local variable
var name string = "Ashraful"

fmt.Println(name)

}
func f(){
fmt.Println(name)
}


//It will give an error because f() doesn't have access to "name"