package main

import (
	"fmt"
	"github.com/Jeroenimoo/GoKitchen/kitchen"
	"net/http"
	"encoding/json"
	"time"
	"github.com/Jeroenimoo/GoKitchen/util"
	"github.com/julienschmidt/httprouter"
)

// Initializes the storage we will use
var suppliers = make(map[string]kitchen.Supply)
var storage = kitchen.NewStorage()
var cooks = make(map[string]kitchen.Cook)

func main() {
	fmt.Println("Hello restaurant!")

	router := httprouter.New()
	router.GET("/layout", webLayout)
	router.POST("/layout/:row", webAdd)
	router.DELETE("/layout/:row/:node", webDel)

	go http.ListenAndServe(":8080", router)

	// Create suppliers
	suppliers["bread"] = kitchen.NewSupply(storage, kitchen.Bread, time.Second)
	suppliers["Cheese"] = kitchen.NewSupply(storage, kitchen.Cheese, time.Second)
	suppliers["Tomato"] = kitchen.NewSupply(storage, kitchen.Tomato, time.Second)
	suppliers["Lettuce"] = kitchen.NewSupply(storage, kitchen.Lettuce, time.Second)

	for _, supply := range suppliers {
		go supply.Start()
	}

	// Create a cook and start working
	cooks["John"] = kitchen.NewCook(storage, time.Second)
	cooks["Bob"] = kitchen.NewCook(storage, time.Second)

	for _, cook := range cooks {
		go cook.Start()
	}

	// Wait for burgers to be created
	for i := 0; i < 4; i++ {
		fmt.Println("Waiting for burger ", i)
		<-storage.Burger
	}

	for _, supply := range suppliers {
		go supply.Stop()
	}

	for _, cook := range cooks {
		go cook.Stop()
	}

	for {
		continue
	}

	fmt.Println("Nom nom nom!")
}

func webLayout(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	itemIds := make([]string, len(suppliers))

	var counter int
	for k := range suppliers {
		itemIds[counter] = k
		counter++
	}

	cookIds := make([]string, len(cooks))

	counter = 0
	for k := range cooks {
		cookIds[counter] = k
		counter++
	}

	var f interface{}
	f = map[string]interface{}{
		"supply":  itemIds,
		"storage": [] string{"storage"},
		"cooks":   cookIds,
	}

	fmt.Println("Hello request!")

	json.NewEncoder(w).Encode(f)
}

func webAdd(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	row := ps.ByName("row")

	switch row {
	case "supply":
		suppliers[util.RandStringBytesMaskImprSrc(4)] = kitchen.NewSupply(storage, kitchen.Bread, time.Second)
	case "cooks":
		cooks[util.RandStringBytesMaskImprSrc(4)] = kitchen.NewCook(storage, time.Second)
	default:
		http.Error(w, "Unkown row "+row, 404)
	}
}

func webDel(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	row := ps.ByName("row")
	node := ps.ByName("node")

	switch row {
	case "supply":
		if _, ok := suppliers[node]; !ok {
			http.Error(w, "Unkown node "+node, 404)
			return
		}

		delete(suppliers, node)
	case "cooks":
		if _, ok := suppliers[node]; !ok {
			http.Error(w, "Unkown node "+node, 404)
			return
		}

		delete(cooks, node)
	default:
		http.Error(w, "Unkown row "+row, 404)
	}
}
