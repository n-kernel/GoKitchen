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
var suppliers = make(map[string]*kitchen.Supply)
var storage = kitchen.NewStorage()
var cooks = make(map[string]*kitchen.Cook)

func main() {
	fmt.Println("Hello restaurant!")

	go kitchen.NodeStatusNotify.Run()

	router := httprouter.New()
	router.GET("/layout", webLayout)
	router.POST("/layout/:row", webAdd)
	router.DELETE("/layout/:row/:node", webDel)
	router.GET("/status", webStatus)

	go http.ListenAndServe(":8080", router)

	// Create suppliers
	addSupply("Bread", kitchen.Bread)
	addSupply("Cheese", kitchen.Cheese)
	addSupply("Tomato", kitchen.Tomato)
	addSupply("Lettuce", kitchen.Lettuce)


	for _, supply := range suppliers {
		go supply.Start()
	}

	// Create a cook and start working
	cooks["John"] = kitchen.NewCook("John", storage, time.Second)
	cooks["Bob"] = kitchen.NewCook("Bob", storage, time.Second)

	for _, cook := range cooks {
		go cook.Start()
	}

	// Wait for burgers to be created
	for i := 0; i < 5000; i++ {
		//fmt.Println("Waiting for burger ", i)
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

func addSupply(name string, item kitchen.Item) {
	suppliers[name] = kitchen.NewSupply(name, storage, item, time.Second * 3)
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

	json.NewEncoder(w).Encode(f)
}

func webAdd(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	row := ps.ByName("row")

	switch row {
	case "supply":
		name := util.RandStringBytesMaskImprSrc(4)
		suppliers[name] = kitchen.NewSupply(name, storage, kitchen.Bread, time.Second)
	case "cooks":
		name := util.RandStringBytesMaskImprSrc(4)
		cooks[name] = kitchen.NewCook(name, storage, time.Second)
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

func webStatus(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	listener := make(chan []byte)

	kitchen.NodeStatusNotify.Joining <- listener

	defer func() {
		kitchen.NodeStatusNotify.Leaving <- listener
	}()

	writerClose := w.(http.CloseNotifier).CloseNotify()

	go func() {
		<-writerClose
		kitchen.NodeStatusNotify.Leaving <- listener
	}()

	for {
		fmt.Fprint(w, "data: ", string(<-listener), "\n\n")
		flusher.Flush()
	}
}
