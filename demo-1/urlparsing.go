package main

import "fmt"
//import "net"
import "net/url"

func main() {
	s := "postgres://user:pass@host:port/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Scheme)

}