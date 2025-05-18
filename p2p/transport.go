package p2p

// Peer is an interface representing remote node
type Peer interface {
}

// Transport is a communication handler (eg tcp/udp/websockets/grpc/etc...)
type Transport interface {
	ListenAndAccept() error

}
