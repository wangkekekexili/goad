package model

import (
	"bufio"
	"container/heap"
	"errors"
	"os"
	"strconv"
	"strings"
)

type RoadNetwork struct {
	nodes map[int]*node
	edges map[int]*edge
}

func CreateRoadNetworkFromFiles(nodeFileName, edgeFileName string) (*RoadNetwork, error) {

	roadNetwork := &RoadNetwork{
		nodes: make(map[int]*node),
		edges: make(map[int]*edge),
	}

	nodeFile, err := os.Open(nodeFileName)
	if err != nil {
		return nil, err
	}
	defer nodeFile.Close()

	edgeFile, err := os.Open(edgeFileName)
	if err != nil {
		return nil, err
	}
	defer edgeFile.Close()

	reader := bufio.NewScanner(nodeFile)
	for reader.Scan() {
		line := reader.Text()
		items := strings.Split(line, " ")
		if len(items) != 3 {
			return nil, errors.New("Format error.")
		}
		id, err := strconv.Atoi(items[0])
		if err != nil {
			return nil, err
		}
		x, err := strconv.ParseFloat(items[1], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(items[2], 64)
		if err != nil {
			return nil, err
		}
		newNode := &node{
			id:    id,
			x:     x,
			y:     y,
			edges: make(map[int]*edge),
		}
		roadNetwork.nodes[id] = newNode
	}

	reader = bufio.NewScanner(edgeFile) // 0 0 5 40.444225
	for reader.Scan() {
		line := reader.Text()
		items := strings.Split(line, " ")
		if len(items) != 4 {
			return nil, errors.New("Format error.")
		}
		id, err := strconv.Atoi(items[0])
		if err != nil {
			return nil, err
		}
		node1Id, err := strconv.Atoi(items[1])
		if err != nil {
			return nil, err
		}
		node2Id, err := strconv.Atoi(items[2])
		if err != nil {
			return nil, err
		}
		length, err := strconv.ParseFloat(items[3], 64)
		newEdge := &edge{
			id:      id,
			node1Id: node1Id,
			node2Id: node2Id,
			length:  length,
		}
		roadNetwork.edges[id] = newEdge
		roadNetwork.nodes[node1Id].edges[id] = newEdge
		roadNetwork.nodes[node2Id].edges[id] = newEdge
	}

	return roadNetwork, nil
}

func (roadNetwork *RoadNetwork) NumberOfNodes() int {
	if roadNetwork == nil {
		return 0
	}
	return len(roadNetwork.nodes)
}

func (roadNetwork *RoadNetwork) NumberOfEdges() int {
	if roadNetwork == nil {
		return 0
	}
	return len(roadNetwork.edges)
}

// nodeHeap is used for Dijkstra algorithm.
type nodeHeap []*node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h nodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].indexInHeap, h[j].indexInHeap = h[j].indexInHeap, h[i].indexInHeap
}
func (h *nodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	node := x.(*node)
	node.indexInHeap = len(*h)
	*h = append(*h, node)
}
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Distance returns the min distance between the nodes with ID startId and destinationId.
// It returns -1 if ids don't exist or there is no path between the two nodes.
func (roadNetwork *RoadNetwork) Distance(startId, destinationId int) float64 {
	if roadNetwork == nil || roadNetwork.nodes[startId] == nil || roadNetwork.nodes[destinationId] == nil {
		return -1
	}
	roadNetwork.initDijkstra()
	startNode := roadNetwork.nodes[startId]
	startNode.distance = 0
	startNode.nodeStatus = reached
	nodeHeap := &nodeHeap{startNode}
	heap.Init(nodeHeap)

	for nodeHeap.Len() != 0 {
		currentNode := heap.Pop(nodeHeap).(*node)
		currentNode.nodeStatus = finished
		if currentNode.id == destinationId {
			return currentNode.distance
		}
		for _, edge := range currentNode.edges {
			nextNodeId := edge.otherEndNodeId(currentNode.id)
			nextNode := roadNetwork.nodes[nextNodeId]
			if nextNode.nodeStatus == finished {
				continue
			} else if nextNode.nodeStatus == reached {
				if currentNode.distance+edge.length < nextNode.distance {
					nextNode.distance = currentNode.distance + edge.length
				}
			} else {
				nextNode.distance = currentNode.distance + edge.length
				nextNode.nodeStatus = reached
				heap.Push(nodeHeap, nextNode)
			}
		}
	}
	return -1
}

func (roadNetwork *RoadNetwork) initDijkstra() {
	for _, node := range roadNetwork.nodes {
		node.distance = 0
		node.indexInHeap = -1
		node.nodeStatus = unvisited
	}
}
