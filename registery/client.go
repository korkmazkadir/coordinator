package registery

import (
	"fmt"
	"log"
	"net/rpc"
)

//Client deifnes a node coordinator client
type Client struct {
	IPAddress  string
	PortNumber int
	client     *rpc.Client
}

//Connect rpc dials, if error occures panics
func (c *Client) Connect() {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", c.IPAddress, c.PortNumber))
	if err != nil {
		panic(err)
	}
	c.client = client
}

//Close closes rpc client, if error occures logs
func (c *Client) Close() {
	err := c.client.Close()
	if err != nil {
		log.Printf("Error occured during cordinator client close: %s\n", err)
	}
}

//AddNode adds a node to node list and returns node counts, if error occures panics
func (c *Client) AddNode(nodeAddress string) int {
	var count int
	err := c.client.Call("NodeRegistery.AddNode", nodeAddress, &count)
	if err != nil {
		panic(err)
	}
	return count
}

//GetNodeCount gets the count of added nodes to the coordinator, if error occures panics
func (c *Client) GetNodeCount(nodeAddress string) int {
	var count int
	err := c.client.Call("NodeRegistery.GetNodeCount", nodeAddress, &count)
	if err != nil {
		panic(err)
	}
	return count
}

//GetNodeList gets node list from coordinator service, if error occures panics
func (c *Client) GetNodeList(nodeAddress string) []string {
	var nodeList []string
	err := c.client.Call("NodeRegistery.GetNodeList", nodeAddress, &nodeList)
	if err != nil {
		panic(err)
	}
	return nodeList
}
