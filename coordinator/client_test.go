package coordinator

import "testing"

func TestClient(t *testing.T) {

	client := Client{IPAddress: "localhost", PortNumber: 1234}
	client.Connect()

	nodes := []string{"node-1", "node-2", "node-3"}

	count := client.AddNode(nodes[0])
	if count != 1 {
		t.Errorf("could not retreive correct count, expected 1 returned %d", count)
	}

	count = client.AddNode(nodes[1])
	if count != 2 {
		t.Errorf("could not retreive correct count, expected 2 returned %d", count)
	}

	count = client.GetNodeCount(nodes[0])
	if count != 2 {
		t.Errorf("could not retreive correct count, expected 2 returned %d", count)
	}

	nodeListFromService := client.GetNodeList(nodes[1])
	if nodeListFromService[0] != nodes[0] || nodeListFromService[1] != nodes[1] {
		t.Errorf("could not retreive correct node list %v", nodeListFromService)
	}

}
