package cluster

import "github.com/CloudGoSight/cloudgosight_api/biz/model"

type Node interface {
	// Init a node from database model
	Init(node *model.Node)

	// Check if given feature is enabled
	IsFeatureEnabled(feature string) bool

	// Subscribe node status change to a callback function
	SubscribeStatusChange(callback func(isActive bool, id uint))
}
