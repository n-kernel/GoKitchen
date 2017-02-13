package main

import (
    "fmt"
    "time"
)

type Storage struct {
    bread chan bool
    cheese chan bool
    tomato chan bool
    lettuce chan bool

    burger chan bool
}

var storage = Storage {
    make(chan bool),
    make(chan bool),
    make(chan bool),
    make(chan bool),

    make(chan bool),
}

func main() {
    fmt.Println("Hello restaurant!")

    // Start assembling of burger
    go assembleBurger()

    // Pass burger ingredients to storage
    storage.bread <- true
    storage.cheese <- true
    storage.tomato <- true
    storage.lettuce <- true

    // Wait for burger to be created
    <-storage.burger
    fmt.Println("Nom nom nom!")
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