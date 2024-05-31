package peers

import (
	"reflect"
	"time"

	"github.com/BharadwajaD/p2pfs/pkg/utils"
)

type Message struct {
	data string
}

type Node struct {
	address string
	inChan  map[string]<-chan *Message //read only chan
	outChan map[string]chan<- *Message //write only chan
}

func NewNode(address string) Node {
	return Node{
		address: address,
		inChan:  make(map[string]<-chan *Message),
		outChan: make(map[string]chan<- *Message),
	}
}

//Non Blocking send to peer_addr
func (node *Node) SendTo(peer_addr string, msg *Message) {
	node.outChan[peer_addr] <- msg
}

//Blocking recieve from peer_addr
func (node *Node) ReceiveFrom(peer_addr string) *Message {
	return <-node.inChan[peer_addr]
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

	pid, value, ok := reflect.Select(cases)
	if ok == false  || pid == chan_count{
		//channel closed case or timeout
		return "", nil
	}


	return pids[pid], value.Interface().(*Message)
}

//Broadcasts message to all connected nodes
func (node *Node) Broadcast(msg *Message) {
	//will block if there exist a full outchan
	for _, outchan := range node.outChan {
		outchan <- msg
	}
}
