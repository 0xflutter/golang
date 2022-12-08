package main

import "fmt"

func main() {
	fmt.Println("hello Go")
	fmt.Println("سلام به همه")
	fmt.Println("Bonjour, tout le monde")
	sayHello("ali")
}

func sayHello(name string){
	fmt.Println("hello", name)
}
