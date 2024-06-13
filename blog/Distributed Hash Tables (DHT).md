Many p2p based applications store metadata using DHT. uTorrent uses a system known as Distributed Hash Table (DHT) and The size of the DHT is approximately **8.4 million nodes** [Reference](http://bittorrent.org/beps/bep_0042.html). Apart from torrent, DHT is used in GNUnet, IPFS, Oracle Coherence.
Seeing its importance, I started implementing DHT. A DHT service should provide: 

````Go
Insert(key, value)
Lookup(key) -> ID of node

// Node related
Join(node) -> Joins the node into the network
Remove(node)
````

## Chord

[Paper](https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf)

The implementation involves:

1. Consistent Hashing
1. Finger tables

## Kademlia

[Paper](https://pdos.csail.mit.edu/~petar/papers/maymounkov-kademlia-lncs.pdf)
