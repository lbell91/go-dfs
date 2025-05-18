package p2p

// HandshakeHandler function... ?
type HandshakeHandler func(Peer) error

func NOPHandshakeHandler (any) error {return nil}

