package main

import (
    "fmt"
    "time"
)

// Holds buffers of ingredients
type Storage struct {
    bread chan bool
    cheese chan bool
    tomato chan bool
    lettuce chan bool

    burger chan bool
}

// Initializes the storage we will use
var storage = Storage {
    make(chan bool, 10),
    make(chan bool, 10),
    make(chan bool, 10),
    make(chan bool, 10),

    make(chan bool, 10),
}

func main() {
    fmt.Println("Hello restaurant!")

    // Hires a cook to make burgers
    go hireCook()

    // Pass burger ingredients to storage
    supply(storage.bread, 4)
    supply(storage.cheese, 4)
    supply(storage.tomato, 4)
    supply(storage.lettuce, 4)

    // Wait for burgers to be created
    for i := 0; i < 4; i++ {
        fmt.Println("Making burger ", i)
        <-storage.burger
    }

    fmt.Println("Nom nom nom!")
}

// A cook which makes burgers infinitely
func hireCook() {
    for {
        assembleBurger()
    }
}

func assembleBurger() {
    fmt.Println("Started making a burger!")

    // Wait for burger items to be available
    <-storage.bread
    <-storage.cheese
    <-storage.tomato
    <-storage.lettuce

    // Wait for burger preparation time
    time.Sleep(time.Second)

    storage.burger <- true
    fmt.Println("Created a burger!")
}

func supply(ingredient chan bool, amount int) {
    for i := 0; i < amount; i++ {
        ingredient <- true
    }
}