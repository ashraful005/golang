package main

import "fmt"

func main(){
//both gives same output
fmt.Println("1 + 1 = ",1+1)      //integer
fmt.Println("1 + 1 = ",1.0+1.0)  //floating number

//strings
fmt.Println("Hello World")
fmt.Println(len("Hello World"))
fmt.Println("Hello World"[1])
fmt.Println("Hello " + "World")

fmt.Println("Hello" == "World")
fmt.Println("Hello" == "Hello")


//variables vs constants
var a int = 22
const b int =22
var name, city string = "Ashraful", "Cumilla"
fmt.Println(a)
fmt.Println(b)
}