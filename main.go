package main

import (
	"fmt"
	"github.com/Jeroenimoo/GoKitchen/kitchen"
	"net/http"
	"encoding/json"
	"time"
	"github.com/Jeroenimoo/GoKitchen/util"
)

// Initializes the storage we will use
var suppliers = make(map[string]kitchen.Supply)
var storage = kitchen.NewStorage()
var cooks = make(map[string]kitchen.Cook)

type WebAdd struct {
	Row string `json:"row"`
}

type WebDel struct {
	Row  string `json:"row"`
	Node string `json:"node"`
}

func main() {
	fmt.Println("Hello restaurant!")

	http.HandleFunc("/layout", webLayout)
	http.HandleFunc("/layout/add", webAdd)
	http.HandleFunc("/layout/del", webDel)

	go http.ListenAndServe(":8080", nil)

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

func webLayout(w http.ResponseWriter, _ *http.Request) {
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

func webAdd(w http.ResponseWriter, r *http.Request) {
	var data WebAdd
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	switch data.Row {
	case "supply":
		suppliers[util.RandStringBytesMaskImprSrc(4)] = kitchen.NewSupply(storage, kitchen.Bread, time.Second)
	case "cooks":
		cooks[util.RandStringBytesMaskImprSrc(4)] = kitchen.NewCook(storage, time.Second)
	default:
		http.Error(w, "Unkown row "+data.Row, 400)
	}
}

func webDel(w http.ResponseWriter, r *http.Request) {
	var data WebDel
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	switch data.Row {
	case "supply":
		delete(suppliers, data.Node)
	case "cooks":
		delete(cooks, data.Node)
	default:
		http.Error(w, "Unkown row "+data.Row, 400)
	}
}
