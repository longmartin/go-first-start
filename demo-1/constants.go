package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000

	const d = 3e20/50000

	fmt.Println(d)

	fmt.Println(math.Sin(n))

}
