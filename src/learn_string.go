package main

import "fmt"

func main() {
	fmt.Println("Simple String")
	fmt.Println(`
This is a multiline. \n
String. That can also contain "quotes"
	`)
	fmt.Println(":) ")
	fmt.Println("\u2272")
}
