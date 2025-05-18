package main

import (
	"fmt"
	"github/lbell91/go-dfs/p2p"
)

func main() {
	transport := BuildTcpTransport(":4000")

	fmt.Println("Good to go for now")
}
