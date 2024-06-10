package rpc_peers

import "fmt"

type Network struct {
	nodes map[string]*Node
}

// create ncount nodes
func NewNetwork(ncount uint) Network {

	nodes := make(map[string]*Node)
	var port uint = 1234
	ipaddr := "127.0.0.1"

	for i := 0; i < int(ncount); i++ {
		node := NewNode(ipaddr, port)
		nodes[fmt.Sprintf("%s:%d", ipaddr, port)] = &node
		port++
	}

	return Network{
		nodes: nodes,
	}
}

func (network *Network) GetNode(node_id int) *Node {

	ipaddr := "127.0.0.1"
	port := 1234
	return network.nodes[fmt.Sprintf("%s:%d", ipaddr, port+node_id)]
}
