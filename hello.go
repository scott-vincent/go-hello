package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/scott-vincent/go-hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseString("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
