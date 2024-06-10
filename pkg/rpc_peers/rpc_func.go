package rpc_peers

import "log"

/**
* rpc exported functions
 */
func (n *Node) Echo(msg string, reply *string) error {
	*reply = msg
    log.Printf("Called echo of node %d with msg: %+v and reply : %+v\n", n.port, msg, *reply)
	return nil
}

