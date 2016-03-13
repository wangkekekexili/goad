package main

import (
	"github.com/wangkekekexili/goad/model"
	"fmt"
)

func main() {
	nodeNetwork, _ := model.CreateRoadNetworkFromFiles("data/simple.node", "data/simple.edge")
}
