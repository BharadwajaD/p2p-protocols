package peers

import (
	"reflect"
	"time"

	"github.com/BharadwajaD/p2pfs/pkg/utils"
)

type Mtype int

const  (
    Command Mtype = iota
)

type Message struct {
    mtype Mtype
	data string
}

type Node struct {
	network *Network
	address string

	State           State
	inChan          map[string]<-chan *Message //read only chan
	outChan         map[string]chan<- *Message //write only chan
	UnConnectedChan chan *Message
}

func NewNode(address string, network *Network) Node {
	return Node{
		network:         network,
		address:         address,
		State:           NewState(),
		inChan:          make(map[string]<-chan *Message),
		outChan:         make(map[string]chan<- *Message),
		UnConnectedChan: make(chan *Message, utils.CHANNEL_SIZE),
	}
}

// send without connection
func (node *Node) SendUnConnected(peer_addr string, msg *Message) {
	node.network.Nodes[peer_addr].UnConnectedChan <- msg
}

// Non Blocking send to peer_addr
func (node *Node) SendConnected(peer_addr string, msg *Message) {
	node.outChan[peer_addr] <- msg
}

// Blocking recieve from peer_addr
func (node *Node) ReceiveFrom(peer_addr string) *Message {
	return <-node.inChan[peer_addr]
}

func (node *Node) RecieveUnConnected() *Message {
	return <-node.UnConnectedChan
}

// Message from some channel
// returns nil if it doesn't Receive msg from any channel after utils.CHANNEL_TIMEOUT seconds
func (node *Node) Receive() (string, *Message) {

	chan_count := len(node.inChan)
	cases := make([]reflect.SelectCase, chan_count)
	pids := make([]string, len(node.inChan))
	i := 0
	for peer_id, inchan := range node.inChan {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(inchan)}
		pids[i] = peer_id
		i++
	}

	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv,
		Chan: reflect.ValueOf(time.After(utils.CHANNEL_TIMEOUT))}) //timeout

	idx, value, ok := reflect.Select(cases)
	if ok == false || idx == chan_count {
		return "", nil //channel closed case or timeout
	}

	return pids[idx], value.Interface().(*Message)
}

// Broadcasts message to all connected nodes
func (node *Node) Broadcast(msg *Message) {
	for _, outchan := range node.outChan {
		outchan <- msg
	}
}

func (node *Node) StartNode() {
    //ever running loop
    for {
    }
}
