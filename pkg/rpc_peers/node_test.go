package rpc_peers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var network = NewNetwork(5)

func TestNetworkCreate(t *testing.T) {
	assert.Equal(t, uint(1234), network.nodes["127.0.0.1:1234"].port)
}

func TestSyncCall(t *testing.T) {
	n0 := network.GetNode(0)

	msg := "Hello from 0"
	reply := ""
	err := n0.Call(1, "Node", "Echo", msg, &reply)

	assert.Nil(t, err)
	assert.Equal(t, msg, reply)
}

func TestAsyncCall(t *testing.T) {

	n0 := network.GetNode(0)

	var msg any = "Hello from 0"
	reply_chan := make(chan string, 1)
	go func() {
		var reply string
		n0.Call(1, "Node", "Echo", msg, &reply)
		reply_chan <- reply
	}()

	reply := <-reply_chan
	assert.Equal(t, msg, reply)
}
