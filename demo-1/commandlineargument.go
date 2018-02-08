package main

import "os"
import "fmt"

func main() {
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(arg)
	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
}