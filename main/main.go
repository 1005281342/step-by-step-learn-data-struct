package main

import (
	"fmt"
	"github.com/1005281342/step-by-step-learn-data-struct/HashMap"
)

func main() {
	ln := HashMap.NewLinkNode()
	ln.Value = 10
	elem := HashMap.LinkNode{20, nil}
	ln = HashMap.AppendLeft(*ln, elem)
	fmt.Printf("%+v", elem)
	fmt.Printf("%+v", ln)
}
