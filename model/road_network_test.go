package model

import (
	"testing"
)

func TestRoadNetwork_Distance(t *testing.T) {
	nodeNetwork, _ := CreateRoadNetworkFromFiles("../data/simple.node", "../data/simple.edge")
	if nodeNetwork.Distance(0, 5) != 11 {
		t.Errorf("BUG!")
	}
}
