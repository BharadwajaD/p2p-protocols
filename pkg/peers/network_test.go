package peers

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//test SendTo and RecieveFrom functions
func TestNodeSend(t *testing.T) {
	graph := SampleGraph()
	network := GraphToNetwork(&graph)

	msg := Message{
		data: "From 0 To 1",
	}

	n0 := network.nodes["0"]
	n1 := network.nodes["1"]

	go n0.SendTo("1", &msg)
	recv := n1.ReceiveFrom("0")

	assert.Equal(t, msg, *recv)
}

//test Recieve function
func TestNodeRecieve(t *testing.T) {
	graph := SampleGraph()
	network := GraphToNetwork(&graph)

	msg := Message{
		data: "From 0 To 1",
	}

	n0 := network.nodes["0"]
	go n0.SendTo("1", &msg)
	peer, rmsg := network.nodes["1"].Receive()

	assert.Equal(t, peer, "0")
	assert.Equal(t, msg, *rmsg)
}

//test BroadCast function
func TestNodeBroadCast(t *testing.T) {
	graph := SampleGraph()
	network := GraphToNetwork(&graph)

	msg := Message{
		data: "From 0 To 1",
	}

	n0 := network.nodes["0"]
	go n0.Broadcast(&msg)

	for i := 1; i < graph.ncount; i++ {
		rmsg := network.nodes[strconv.Itoa(i)].ReceiveFrom("0")
		//assert.Equal(t, peer, "0")
		assert.Equal(t, msg, *rmsg)
	}
}

//test NoEdge between the nodes case
func TestNoEdge(t *testing.T) {
	graph := SampleGraph()
	network := GraphToNetwork(&graph)

	msg := Message{
		data: "From 0 To 1",
	}

	n1 := network.nodes["1"]
	go n1.Broadcast(&msg)

	peer, rmsg := network.nodes[strconv.Itoa(3)].Receive()
	assert.Equal(t, peer, "")
	assert.Nil(t, rmsg) //nil ptr is different from nil interface
}
