package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TcpPeer represents the remote node over established TCP Connection
type TcpPeer struct {
	// conn underlying connection of the peer node
	conn         net.Conn

	// call and retrieve connection => outbound true
	// accept and consume connection => outbound false
	outboundPeer bool
}

type TcpTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func BuildTcpPeer(conn net.Conn, outboundPeer bool) *TcpPeer {
	return &TcpPeer{
		conn: conn,
		outboundPeer: outboundPeer,
	}
}

func BuildTcpTransport(listenAddr string) *TcpTransport {
	return &TcpTransport{
		listenAddress: listenAddr,
	}
}

func (t *TcpTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TcpTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		
		go t.handleConn(conn)
	}
}



func (t *TcpTransport) handleConn(conn net.Conn) {
	peer := BuildTcpPeer(conn, true)
	
	fmt.Printf("New incoming connection %+v\n", peer)
}
