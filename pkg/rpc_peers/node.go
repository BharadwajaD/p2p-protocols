package rpc_peers

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//export Node for rpc
type Node struct {
	node_id    uint8
	ipaddr     string
	port       uint
	rpc_server *rpc.Server
}

// create node and register it to rpc server
func NewNode(ipaddr string, port uint) Node {

	host := fmt.Sprintf("%s:%d", ipaddr, port)
	node_id := uint8(port - 1234)

	rpc_server := rpc.NewServer()
	node := Node{
		ipaddr:     ipaddr,
		port:       port,
		node_id:    node_id,
		rpc_server: rpc_server,
	}

	err := rpc_server.RegisterName(fmt.Sprintf("Node%d", node.node_id), &node)
	if err != nil {
		log.Fatalf("node: %+v registering error %+v\n", node, err)
	}

	rpc_server.HandleHTTP(fmt.Sprintf("/rpc%d", node.node_id), fmt.Sprintf("/debug%d", node.node_id))

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal("rpc listner error\n")
	}

	go http.Serve(listener, nil)
	return node
}

func (n *Node) RegisterName(name string, _type any) error {
	return n.rpc_server.RegisterName(fmt.Sprintf("%s%d", name, n.node_id), _type)
}

// args any
// reply *any
func (n *Node) Call(peer_node_id int, service_type, service_method string, args, reply any) error {

	host := fmt.Sprintf("127.0.0.1:%d", 1234+peer_node_id)
	path := fmt.Sprintf("/rpc%d", peer_node_id)
	service_method = fmt.Sprintf("%s%d.%s", service_type, peer_node_id, service_method)

	rpc_client, err := rpc.DialHTTPPath("tcp", host, path)
	if err != nil {
		return err
	}

	err = rpc_client.Call(service_method, args, reply)
	return err
}
