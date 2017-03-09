package main

import (
	"fmt"
	"github.com/Jeroenimoo/GoKitchen/kitchen"
	"net/http"
	"encoding/json"
	"time"
	"github.com/Jeroenimoo/GoKitchen/util"
	"github.com/julienschmidt/httprouter"
	"github.com/Jeroenimoo/GoKitchen/comm"
	"math/rand"
)

// Initializes the storage we will use
var suppliers = make(map[string]*kitchen.Supply)
var storage = kitchen.NewStorage()
var cooks = make(map[string]*kitchen.Cook)
var customers = make(map[string]*kitchen.Customer)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Hello restaurant!")

	go kitchen.EventBus.Run()

	router := httprouter.New()
	router.GET("/layout", webLayout)
	router.POST("/layout/:row", webAdd)
	router.DELETE("/layout/:row/:node", webDel)
	router.GET("/status", webStatus)

	router.NotFound = http.FileServer(http.Dir("static"))

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

	addCustomer("Harry", kitchen.Burger)
	addCustomer("Bilbo", kitchen.Burger)
	addCustomer("BEAKFAST", kitchen.Burger)

	for _, customer := range customers {
		go customer.Run()
	}

	for {
		continue
	}

	for _, supply := range suppliers {
		go supply.Stop()
	}

	for _, cook := range cooks {
		go cook.Stop()
	}

	for _, customer := range customers {
		go customer.Stop()
	}

	fmt.Println("Nom nom nom!")
}

func addSupply(name string, item kitchen.Item) {
	refill := time.Duration(float32(time.Second) * (2.0 + 2.0 * rand.Float32()))
	suppliers[name] = kitchen.NewSupply(name, storage, item, refill)
}

func addCustomer(name string, item kitchen.Item) {
	customer := kitchen.NewCustomer(name, storage, item, time.Second * 10, time.Second)
	customers[name] = customer

	go func() {
		customer.Run()
		fmt.Println(customer.Name, " left the building saying: ", customer.ExitMessage)
		// Not thread safe, im wild
		delete(customers, name)
	}()
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

	customerIds := make([]string, len(customers))

	counter = 0
	for k := range customers {
		customerIds[counter] = k
		counter++
	}

	var f interface{}
	f = map[string]interface{}{
		"supply":  itemIds,
		"storage": [] string{"storage"},
		"cooks":   cookIds,
		"customers":   customerIds,
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
	case "customers":
		addCustomer(util.RandStringBytesMaskImprSrc(4), kitchen.Burger)
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

	listener := make(chan *comm.Event)

	kitchen.EventBus.Joining <- listener

	defer func() {
		kitchen.EventBus.Leaving <- listener
	}()

	writerClose := w.(http.CloseNotifier).CloseNotify()

	go func() {
		<-writerClose
		kitchen.EventBus.Leaving <- listener
	}()

	for {
		event := <-listener
		data := event.Data
		jsonData, err := json.Marshal(data)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Fprint(w, "event: test\n")//, event.Name)
		fmt.Fprint(w, "data: ", string(jsonData), "\n\n")
		flusher.Flush()
	}
}
