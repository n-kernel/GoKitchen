package kitchen

import (
	"github.com/Jeroenimoo/GoKitchen/comm"
)

type Item int
const (
	Bread   Item = iota
	Cheese
	Tomato
	Lettuce

	Burger
)

type enumItem struct {
	Item
	name string
}

var Items = []Item {
	Bread,
	Cheese,
	Tomato,
	Lettuce,
	Burger,
}

var Ingredients = []Item {
	Bread,
	Cheese,
	Tomato,
	Lettuce,
}

var EnumItems = []enumItem{
	{Bread, "BREAD"},
	{Cheese, "CHEESE"},
	{Tomato, "TOMATO"},
	{Lettuce, "LETTUCE"},

	{Burger, "BURGER"},
}

func (i Item) GetName() string {
	return EnumItems[i].name
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
	NodeTypeCustomer = "CUSTOMER"
)

type Node struct {
	Type NodeType
	Name string
}

var EventBus = comm.NewEventBus()

func (s *Node) updateStatus(status Status, message string) {
	data := map[string]interface{} {
		"type": s.Type,
		"name": s.Name,
		"status": status,
		"message": message,
	}

	EventBus.Publish <- &comm.Event{"nodeStatus", data}
}