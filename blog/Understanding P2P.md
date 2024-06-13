
I started with a question, "How does one store data in decentralised applications ?", the answer was using [IPFS](https://ipfs.tech/) -  A modular suite of **protocols** and standards for organizing and moving data, designed from the ground up with the principles of content addressing and **peer-to-peer** networking. 
Now I moved my focus towards learning p2p protocols. Learning by implementing is an efficient way of learning. But unlike single computer based algorithms, implementing p2p protocols need a setup of nodes(peers) and communication among them.

1. [[RPC, Nodes and Network]]: Creating network of nodes using [net/rpc](https://pkg.go.dev/net/rpc) package.
2. [[Distributed Hash Tables (DHT)]]: Distributed look up service similar to hash table. 