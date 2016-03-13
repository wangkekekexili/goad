package model

import (
	"bufio"
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
			id: id,
			x:  x,
			y:  y,
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
	}

	return roadNetwork, nil
}

func (roadNetwork *RoadNetwork) NumberOfNodes() int {
	if roadNetwork == nil {
		return 0
	}
	return len(roadNetwork.nodes)
}

func (RoadNetwork *RoadNetwork) NumberOfEdges() int {
	if RoadNetwork == nil {
		return 0
	}
	return len(RoadNetwork.edges)
}
