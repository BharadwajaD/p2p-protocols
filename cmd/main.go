package main

import (
	//"github.com/BharadwajaD/p2pfs/pkg/dht"
	"github.com/BharadwajaD/p2pfs/pkg/peers"
)


func InitLocalDHT(network *peers.Network) {
    for node_id, node := range network.Nodes {
    }
}

func main(){
    ncount := 15
    network := peers.NewNetwork(ncount, false)

    InitLocalDHT(&network)

}
