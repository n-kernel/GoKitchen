package kitchen

import (
	"github.com/Jeroenimoo/GoKitchen/util/notify"
)

type Item int
const (
	Bread   Item = iota
	Cheese
	Tomato
	Lettuce

	Burger
)

var Items = []Item{
	Bread,
	Cheese,
	Tomato,
	Lettuce,

	Burger,
}

type Status string
const (
	Working = "WORKING"
	Waiting = "WAITING"
	Finished = "FINISHED"
)

type NodeType string
const (
	NodeTypeSupply = "SUPPLY"
	NodeTypeStorage = "STORAGE"
	NodeTypeCook = "COOK"
)

type Node struct {
	Type NodeType
	Name string
}

var NodeStatusNotify = notify.NewNotifier()

func (s *Node) updateStatus(status Status) {
	go func() {
		data := map[string]interface{} {
			"type": s.Type,
			"name": s.Name,
			"status": status,
		}

		NodeStatusNotify.Message(data)
	}()
}