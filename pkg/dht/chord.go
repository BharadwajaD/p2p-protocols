package dht

import (
	"crypto/sha1"
	"strconv"

	"github.com/BharadwajaD/p2pfs/pkg/rpc_peers"
)

func hashKey(key string) uint32 {
	h := sha1.New()
	h.Write([]byte(key))
	bs := h.Sum(nil)

	// Use first 4 bytes for a 32-bit hash
	return (uint32(bs[0]) << 24) |
		(uint32(bs[1]) << 16) |
		(uint32(bs[2]) << 8) |
		uint32(bs[3])
}

// 1. consistent hashing
// 2. successor node_id
// 3. finger table
type LocalDHT struct {
	node *rpc_peers.Node //for rpc calls

	node_id int
	succ_id int //node ids

	HashTable map[string]string //local hash map

	//TODO: finger table: {node_id + 2^(i-1), peer_id}
}

func NewLocalDHT(node *rpc_peers.Node, node_id int) *LocalDHT {
    ldht := LocalDHT{
		node:      node,
		node_id:   node_id,
		succ_id:   -1,
		HashTable: make(map[string]string),
	}

    node.RegisterName("LDHT", &ldht) //LDHT.(i) service is exported for rpc calls
    return &ldht
}

// rpc exported function
// will return node_id of node containing the given key
func (ldht *LocalDHT) FindSuccessor(key string, succ_id *int) error {


	if ldht.succ_id == -1 {
		//single node
		*succ_id = ldht.node_id
		return nil
	}

	key_id := hashKey(key)
	node_hash := hashKey(strconv.Itoa(ldht.node_id))
	succ_hash := hashKey(strconv.Itoa(ldht.succ_id))


	if key_id > node_hash && key_id <= succ_hash {
		*succ_id = ldht.succ_id
		return nil
	}

	return ldht.node.Call(ldht.succ_id, "LDHT", "FindSuccessor", key, succ_id) //rpc call
}

// TODO: Join function
// func (ldht *LocalDHT) Join(new_node_id int) error {}

func (ldht *LocalDHT) Echo(msg string, reply *string) error {
    *reply = msg
    return nil
}
