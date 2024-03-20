package main

import (
	"fmt"
	"node/cmd/node"
    "sync"
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

    ocClient := &node.MASClient{}
    ocClient.Username = os.Getenv("MASCLIENTUSERNAME")
    ocClient.Password = os.Getenv("MASCLIENTPASSWORD")
    ocClient.Host = os.Getenv("MASCLIENTHOST")
    ocClient.Namespace = os.Getenv("MASCLIENTNAMESPACE")


    ocClient.BuildConfig()
    podClients, err := ocClient.ListBindablePods()
    if err != nil {
        panic(err.Error())
    }

    client := &client.Client{
        Username: clientUsername,
        Password: clientPassword,
        Endpoint: wsURL,
    }

    var wg sync.WaitGroup;

    for _, podName := range(podClients) {

        wg.Add(1)

        go func(pod string) {


            masClient := &node.MASClient{}
            masClient.Username = os.Getenv("MASCLIENTUSERNAME")
            masClient.Password = os.Getenv("MASCLIENTPASSWORD")
            masClient.Host = os.Getenv("MASCLIENTHOST")
            masClient.Namespace = os.Getenv("MASCLIENTNAMESPACE")
            masClient.PodName = pod
            masClient.BuildConfig()
            connection := client.Build()
            defer connection.Disconnect()
            c := make(chan node.LogMessage)
            go masClient.ReadLogs(c)
            for {
                x := <-c
                connection.Send("/app/log/" + podName, "text/plain", []byte(fmt.Sprintf("%v", x)), nil)
            }

            wg.Done()

        }(podName)
    }

    wg.Wait()

}

