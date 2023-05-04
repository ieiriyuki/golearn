package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println(quote.Go())

    names := []string{"Gradys", "Samantha", "Darrin"}

    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
