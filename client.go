package main

import (
    "time"
    "strconv"
    "github.com/johanhenriksson/goworld/network"
)

func main() {
    cl, err := network.ConnectTo("localhost:1025")
    if err != nil { panic(err) }
    defer cl.Close()

    go cl.Worker()

    i := 0
    frame := time.Now()
    for {
        delta := time.Since(frame).Seconds()
        frame = time.Now()

        msg := "msg" + strconv.Itoa(i)
        i++

        bf := []byte(msg)
        cl.Send(bf)
        cl.Update(delta)

        time.Sleep(2 * time.Second)
    }
}
