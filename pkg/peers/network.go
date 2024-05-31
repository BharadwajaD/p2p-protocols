package peers

import (
	"fmt"
	"strconv"

	"github.com/BharadwajaD/p2pfs/pkg/utils"
)

type Network struct {
	nodes       map[string]*Node
	nodes_count int
}

//Default: Constructs a star network
func NewNetwork(ncount int, isStar bool) Network {

	nodes := make(map[string]*Node)
	for i := 0; i < ncount; i++ {
		addr := strconv.Itoa(i)
		node := NewNode(addr)
		nodes[addr] = &node
	}

	network := Network{
		nodes:       nodes,
		nodes_count: ncount,
	}

    if isStar == false {
        return network
    }

	//fully connected graph
	for i := 0; i < ncount; i++ {
		for j := i + 1; j < ncount; j++ {
			network.AddChannel(strconv.Itoa(i), strconv.Itoa(j))
		}
	}

	return network
}

/**
 *        n1_inchan
 * NODE2 -------> NODE1
 * NODE2 <------- NODE1
 *        n1_outchan
 **/
func (network *Network) AddChannel(node_id1, node_id2 string) {
	node1 := network.nodes[node_id1]
	node2 := network.nodes[node_id2]

	n1_inchan := make(chan *Message, utils.CHANNEL_SIZE)
	n1_outchan := make(chan *Message, utils.CHANNEL_SIZE)

	node1.inChan[node_id2] = n1_inchan
	node1.outChan[node_id2] = n1_outchan

	node2.inChan[node_id1] = n1_outchan
	node2.outChan[node_id1] = n1_inchan
}

//Prints nodes and channels in adjlist fashon
func (network *Network) Print() {
	for node_id, node := range network.nodes {
		fmt.Printf("\nNode Id: %s\n", node_id)
        str := ""
		for peer_id := range node.outChan {
            str += peer_id + ","
		}
        fmt.Printf("To %s\n", str)

        str = ""
		for peer_id := range node.inChan {
            str += peer_id + ","
		}
        fmt.Printf("From %s\n", str)
	}
}

//Useful for Marshalling and UnMarshalling
type Graph struct{
    ncount int
    edges [][]int //[from, to]
}

func SampleGraph() Graph{
    return Graph{
        ncount: 4,
        edges: [][]int{{1, 2}, {0, 1},{0, 2}, {0, 3}},
    }
}

func GraphToNetwork(graph *Graph) Network {
    network := NewNetwork(graph.ncount, false)
    for _, edge := range graph.edges {
        network.AddChannel(strconv.Itoa(edge[0]), strconv.Itoa(edge[1]))
    }

    return network
}
