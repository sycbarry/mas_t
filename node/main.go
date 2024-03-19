package main

import (
	"fmt"
	"node/cmd/node"
	client "node/cmd/client"
    "os"

    "github.com/joho/godotenv"
)



func main() {

    err := godotenv.Load()
    if err != nil {
        panic(err)
    }


    wsURL := os.Getenv("WSHUB")
    clientPassword := os.Getenv("WSCLIENTPASSWORD")
    clientUsername := os.Getenv("WSCLIENTUSERNAME")

    masClient := &node.MASClient{}
    masClient.Username = os.Getenv("MASCLIENTUSERNAME")
    masClient.Password = os.Getenv("MASCLIENTPASSWORD")
    masClient.Host = os.Getenv("MASCLIENTHOST")
    masClient.PodName = os.Getenv("MASCLIENTPODNAME")
    masClient.ContainerName = os.Getenv("MASCLIENTCONTAINERNAME")
    masClient.Namespace = os.Getenv("MASCLIENTNAMESPACE")

    client := &client.Client{
        Username: clientUsername,
        Password: clientPassword,
        Endpoint: wsURL,
    }

    connection := client.Build()
    defer connection.Disconnect()

    c := make(chan node.LogMessage)
    go masClient.ReadLogs(c)

    for {
        x := <-c
        connection.Send("/app/log", "text/plain", []byte(fmt.Sprintf("%v", x)), nil)
    }

}

