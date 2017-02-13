package main

import (
    "fmt"
    "./kitchen"
    "time"
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

    // Pass burger ingredients to storage
    supply(storage.Bread, 4)
    supply(storage.Cheese, 4)
    supply(storage.Tomato, 4)
    supply(storage.Lettuce, 4)

    // Wait for burgers to be created
    for i := 0; i < 4; i++ {
        fmt.Println("Waiting for burger ", i)
        <-storage.Burger
    }

    fmt.Println("Nom nom nom!")
}

func supply(ingredient chan bool, amount int) {
    for i := 0; i < amount; i++ {
        ingredient <- true
    }
}