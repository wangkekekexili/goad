package model

type node struct {
	id    int
	x     float64
	y     float64
	edges map[int]*edge

	// For Dijkstra algorithm.
	current  float64
	shortest float64
	inHeap   bool
}
