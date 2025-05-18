package main

import (
	"fmt"
	"github.com/lbell91/go-dfs/p2p"
)

func main() {
	transport := p2p.BuildTcpTransport(":4000")

	fmt.Println("Good to go for now")
}
