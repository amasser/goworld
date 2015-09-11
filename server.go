package main

import (
    "fmt"
    "time"
    "github.com/johanhenriksson/goworld/network"
)


func main() {
    fmt.Println("Hello world")

    s, err := network.Listen("127.0.0.1:1025")
    if err != nil {
        panic(err)
    }
    defer s.Stop()

    go s.Worker()

    updateRate := 1
    target_time := time.Second / time.Duration(updateRate)
    frame := time.Now()

    for {
        delta := time.Since(frame)
        frame = time.Now()

        s.Update(delta.Seconds())

        wait := target_time - time.Since(frame) - time.Millisecond
        time.Sleep(wait)
    }
}
