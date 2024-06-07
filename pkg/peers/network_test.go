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

	n0 := network.Nodes["0"]
	n1 := network.Nodes["1"]

	go n0.SendConnected("1", &msg)
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

	n0 := network.Nodes["0"]
	go n0.SendConnected("1", &msg)
	peer, rmsg := network.Nodes["1"].Receive()

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

	n0 := network.Nodes["0"]
	go n0.Broadcast(&msg)

	for i := 1; i < graph.ncount; i++ {
		rmsg := network.Nodes[strconv.Itoa(i)].ReceiveFrom("0")
		//assert.Equal(t, peer, "0")
		assert.Equal(t, msg, *rmsg)
	}
}

//test NoEdge between the nodes case
func TestNoEdge(t *testing.T) {
	graph := SampleGraph()
	network := GraphToNetwork(&graph)

	msg := Message{
		data: "Test data",
	}

	n1 := network.Nodes["1"]
	go n1.Broadcast(&msg)

	peer, rmsg := network.Nodes[strconv.Itoa(3)].Receive()
	assert.Equal(t, peer, "")
	assert.Nil(t, rmsg) //nil ptr is different from nil interface
}

func TestUnConnectedChan(t *testing.T) {
	graph := SampleGraph()
	network := GraphToNetwork(&graph)
	msg := Message{
		data: "Test data",
	}

    n3 := network.Nodes["3"]
    go n3.SendUnConnected("2", &msg)

	rmsg := network.Nodes["2"].RecieveUnConnected()
	assert.Equal(t, msg, *rmsg) //nil ptr is different from nil interface
}
