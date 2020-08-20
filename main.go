package main

import (
	"fmt"
	"strings"

	"github.com/richardlehane/crock32"
)

func main() {
	for i := 10; i < 32; i++ {
		fmt.Printf("%2v:%v\n", i, strings.ToUpper(crock32.Encode(uint64(i))))
	}
}
