package model

type node struct {
	id    int
	x     float64
	y     float64
	edges map[int]*edge

	// For Dijkstra algorithm.
	distance    float64
	indexInHeap int
	nodeStatus  nodeStatus
}

// For Dijkstra algorithm
type nodeStatus int32

const (
	unvisited nodeStatus = iota
	reached   nodeStatus = iota
	finished  nodeStatus = iota
)
