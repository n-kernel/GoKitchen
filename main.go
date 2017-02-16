package main

import (
    "fmt"
    "time"
    "github.com/Jeroenimoo/GoKitchen/kitchen"
)

// Initializes the storage we will use
var storage = kitchen.Storage {
    make(chan bool, 10),
    make(chan bool, 10),
    make(chan bool, 10),
    make(chan bool, 10),

    make(chan bool, 10),
}

func main() {
    fmt.Println("Hello restaurant!")

    // Create a cook and start working
    cook1 := kitchen.NewCook(storage, time.Second)
    go cook1.Start()

    // Create suppliers
    supplierBread := kitchen.NewSupply(storage.Bread, time.Second)
    supplierCheese := kitchen.NewSupply(storage.Cheese, time.Second)
    supplierTomato := kitchen.NewSupply(storage.Tomato, time.Second)
    supplierLettuce := kitchen.NewSupply(storage.Lettuce, time.Second)

    go supplierBread.Start()
    go supplierCheese.Start()
    go supplierTomato.Start()
    go supplierLettuce.Start()

    // Wait for burgers to be created
    for i := 0; i < 4; i++ {
        fmt.Println("Waiting for burger ", i)
        <-storage.Burger
    }

    supplierBread.Stop()
    supplierCheese.Stop()
    supplierTomato.Stop()
    supplierLettuce.Stop()

    cook1.Stop()

    fmt.Println("Nom nom nom!")
}