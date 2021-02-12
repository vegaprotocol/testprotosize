package main

import (
	"fmt"

	"github.com/vegaprotocol/testprotosize/proto"
)

func main() {
	x := proto.ThingWithSize{
		Size: 1,
	}
	fmt.Printf("x=%s\n", x.String())
	// Expected: x=size:1
}
