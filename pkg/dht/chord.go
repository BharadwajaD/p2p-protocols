package dht

import (
	"crypto/sha1"
	"math"

	"github.com/BharadwajaD/p2pfs/pkg/peers"
)

type Hash struct {
	nbits uint8
}

func NewHash() Hash {
	return Hash{
		nbits: 32,
	}
}

func (hash *Hash) hashKey(key string) uint32 {
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
	node    *peers.Node //for channel
	hash    Hash
	node_id string
	SuccId  string
	PredId  string //key hashes

	HashMap map[string]string

	//TODO: finger table: {node_id + 2^(i-1), peer_id}
}

func NewLocalDHT(network *peers.Network, node_id string) LocalDHT {
	return LocalDHT{
		node:    network.Nodes[node_id],
		node_id: node_id,
		PredId:  "",
		SuccId:  "",
		HashMap: make(map[string]string),
		hash:    NewHash(),
	}
}

func (ldht *LocalDHT) FindSuccessor(key string) string {
	key_id := ldht.hash.hashKey(key)
	node_hash := ldht.hash.hashKey(ldht.node_id)
	succ_hash := ldht.hash.hashKey(ldht.SuccId)

	if key_id > node_hash && key_id <= succ_hash {
		return ldht.SuccId
	}

	return ldht.FindSuccessor(key)
}
