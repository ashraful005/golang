package main
//global variable
var name string = "Ashraful"

import "fmt"

func main(){

fmt.Println(name)

}
func f(){
fmt.Println(name)
}


//It will give an error because f() doesn't have access to "name"