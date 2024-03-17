package main

import (
	"fmt"
	"math/rand"
	"node/cmd/node"
	client "node/cmd/client"
	"time"
)




func main() {

    wsURL := "ws://localhost:8080/ws"
    clientPassword := "password"
    clientUsername := "client"

    client := &client.Client{
        Username: clientUsername,
        Password: clientPassword,
        Endpoint: wsURL,
    }

    connection := client.Build()
    defer connection.Disconnect()

    masClient := &node.MASClient{}
    masClient.PodName = "zingg-api-6bc8d96d7-jzcmh"
    masClient.Namespace = "zingg"

    c := make(chan node.LogMessage)

    go masClient.ReadLogs(c)

    for {
        x := <-c
        rand.Seed(time.Now().UnixNano())
        connection.Send("/app/log", "text/plain", []byte(fmt.Sprintf("%v", x)), nil)
    }
}

