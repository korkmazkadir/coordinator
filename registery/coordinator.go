package registery

import (
	"log"
	"sync"
)

//NodeRegistery defines node coordinator service
type NodeRegistery struct {
	// keeps node addresses in the
	nodeList []string
	mutex    *sync.Mutex
}

//NewNodeRegistery creates a node registery
func NewNodeRegistery() *NodeRegistery {
	registery := new(NodeRegistery)
	registery.mutex = &sync.Mutex{}
	return registery
}

// AddNode adds a node to list and returns number of nodes
func (nc *NodeRegistery) AddNode(nodeAddress *string, nodeCount *int) error {

	nc.mutex.Lock()
	defer nc.mutex.Unlock()

	nc.nodeList = append(nc.nodeList, *nodeAddress)
	*nodeCount = len(nc.nodeList)

	log.Printf("New node  added to list %s count %d \n", *nodeAddress, *nodeCount)

	return nil
}

// GetNodeCount return number of added nodes, nodeAddress is used for logging puposes
func (nc *NodeRegistery) GetNodeCount(nodeAddress *string, nodeCount *int) error {

	nc.mutex.Lock()
	defer nc.mutex.Unlock()

	*nodeCount = len(nc.nodeList)
	//log.Printf("Node %s requested node count %d \n", *nodeAddress, *nodeCount)

	return nil
}

// GetNodeList returns node list
func (nc *NodeRegistery) GetNodeList(nodeAddress *string, nodeList *[]string) error {

	nc.mutex.Lock()
	defer nc.mutex.Unlock()

	*nodeList = nc.nodeList
	//log.Printf("Node %s requested node list %v \n", *nodeAddress, *nodeList)

	return nil
}
