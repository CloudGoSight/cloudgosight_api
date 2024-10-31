package cluster

import (
	"github.com/CloudGoSight/cloudgosight_api/biz/balance"
	"sync"
)

var Default *NodePool

type Pool interface {
	BalanceNodeByFeature(feature string, lb balance.Balancer)
}

// NodePool 通用节点池
type NodePool struct {
	active   map[uint]Node
	inactive map[uint]Node

	featureMap map[string][]Node

	lock sync.RWMutex
}
