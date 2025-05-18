package main

import (
	"fmt"
	"log"

	"github.com/lbell91/go-dfs/p2p"
)

func main() {
	transport := p2p.BuildTcpTransport(":4000")

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Good to go for now")

	select {

	}
}
