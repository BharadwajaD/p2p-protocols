package dht

import (
	"fmt"
	"testing"

	"github.com/BharadwajaD/p2pfs/pkg/rpc_peers"
	"github.com/stretchr/testify/assert"
)

type testGlobal struct {
	network rpc_peers.Network
    ldht_mp map[int]*LocalDHT
}

func new_test_global() testGlobal {

	tg := testGlobal{
		network: rpc_peers.NewNetwork(4),
        ldht_mp: make(map[int]*LocalDHT),
	}

	//TODO: this should be called while creating nodes
    tg.ldht_mp[0] = NewLocalDHT(tg.network.GetNode(0), 0)
	tg.ldht_mp[1] = NewLocalDHT(tg.network.GetNode(1), 1)
	tg.ldht_mp[2] = NewLocalDHT(tg.network.GetNode(2), 2)
	tg.ldht_mp[3] = NewLocalDHT(tg.network.GetNode(3), 3)

	return tg
}

var tg = new_test_global()

func TestLDHTEcho(t *testing.T) {

	reply := ""
	msg := "Hello LDHT Test"
	err := tg.network.GetNode(0).Call(1, "LDHT", "Echo", msg, &reply)

	assert.Nil(t, err)
	assert.Equal(t, msg, reply)
}

func TestFindSuccSingleNode(t *testing.T) {

    l0 := tg.ldht_mp[0]
    key := "test_key1"
    var succ_id int
    err := l0.FindSuccessor(key, &succ_id)

    //since we have disconnected dhts
    assert.Nil(t, err)
    assert.Equal(t, 0, succ_id)
}

func TestFindSucc(t *testing.T){

    // join manually
    h0 := hashKey("0")
    l0 := tg.ldht_mp[0]

    h1 := hashKey("1")
    l1 := tg.ldht_mp[1]

    h2 := hashKey("2")
    l2 := tg.ldht_mp[2]

    key := "test key another"

    fmt.Println(h0, h1, h2, hashKey(key))

    //l1 -> l0 -> l2 -> l1
    //key in l2

    l1.succ_id = 0
    l0.succ_id = 2
    l2.succ_id = 1

    succ_id := -1

    l1.FindSuccessor(key, &succ_id)
    assert.Equal(t, 2, succ_id)

    l0.FindSuccessor(key, &succ_id)
    assert.Equal(t, 2, succ_id)

    l2.FindSuccessor(key, &succ_id)
    assert.Equal(t, 2, succ_id)
}
