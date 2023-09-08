package main

import "fmt"

// See more details from https://www.liwenzhou.com/posts/Go/functional-options-pattern/
func main() {
	var opts []Option

	opts = append(opts, WithName("test"))
	opts = append(opts, WithStatus("ACTIVE"))
	var token = [16]byte{}
	opts = append(opts, WithToken(token))

	m := NewMap(opts...)
	fmt.Printf("%v", m)
}