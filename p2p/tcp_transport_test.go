package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	listenAddr := ":4443"
	transport := BuildTcpTransport(listenAddr)

	assert.Equal(t, transport.listenAddress, listenAddr)

	transport.ListenAndAccept()

	assert.Nil(t, transport.ListenAndAccept())

}
