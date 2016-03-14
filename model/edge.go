package model

type edge struct {
	id      int
	node1Id int
	node2Id int
	length  float64
}

func (edge *edge) otherEndNodeId(nodeId int) int {
	if edge == nil || edge.node1Id != nodeId && edge.node2Id != nodeId {
		return -1
	}
	if edge.node1Id == nodeId {
		return edge.node2Id
	} else {
		return edge.node1Id
	}
}
