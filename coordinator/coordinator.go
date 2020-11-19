package coordinator

import "log"

//NodeCoordinator defines node coordinator service
type NodeCoordinator struct {
	// keeps node addresses in the
	nodeList []string
}

// AddNode adds a node to list and returns number of nodes
func (nc *NodeCoordinator) AddNode(nodeAddress *string, nodeCount *int) error {

	nc.nodeList = append(nc.nodeList, *nodeAddress)
	*nodeCount = len(nc.nodeList)

	log.Printf("New node  added to list %s count %d \n", *nodeAddress, *nodeCount)

	return nil
}

// GetNodeCount return number of added nodes, nodeAddress is used for logging puposes
func (nc *NodeCoordinator) GetNodeCount(nodeAddress *string, nodeCount *int) error {

	*nodeCount = len(nc.nodeList)
	log.Printf("Node %s requested node count %d \n", *nodeAddress, *nodeCount)

	return nil
}

// GetNodeList returns node list
func (nc *NodeCoordinator) GetNodeList(nodeAddress *string, nodeList *[]string) error {

	*nodeList = nc.nodeList
	log.Printf("Node %s requested node list %v \n", *nodeAddress, *nodeList)

	return nil
}
