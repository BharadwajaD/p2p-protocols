package main

import "github.com/BharadwajaD/p2pfs/pkg/peers"


func main(){
    ncount := 5
    network := peers.NewNetwork(ncount, true)

    network.Print()
}
