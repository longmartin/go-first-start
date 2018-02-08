package main

import "fmt"
import "flag"

func main() {

	wordPtr := flag.String("word", "foo", "bar")

	numbPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("wordPtr:", *wordPtr)
	fmt.Println("numbPtr:", *numbPtr)
	fmt.Println("boolPtr:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}