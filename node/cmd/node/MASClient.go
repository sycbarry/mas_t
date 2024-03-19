package node

import (
	"context"
    "net/http"
    "crypto/tls"

    "strings"
    "io/ioutil"
    "fmt"


	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type MASClient struct {
    ConfigPath string
    PodName string
    Namespace string
    Username string 
    Password string 
    Host string
    ContainerName string
}

func (client *MASClient) ReadLogs(c chan LogMessage) {
    config, err := rest.InClusterConfig()
    if err != nil {
        if(client.ConfigPath != "" ) {
            filePath := client.ConfigPath
            config, err = clientcmd.BuildConfigFromFlags("", filePath)
            if err != nil {
                panic(err.Error())
            }
        }
        if (client.Username != "" && client.Password != "" && client.Host != "") {
            config = &rest.Config {
                Host: client.Host,
                Username: client.Username,
                Password: client.Password,
                BearerToken: "",
                TLSClientConfig: rest.TLSClientConfig{Insecure: true},
            }
        }
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }


    /*
    _, err := getToken(clientset, client.Username, client.Password, client.Host)
    if err != nil {
        panic(err.Error())
    }
    */

    // config.BearerToken = token


    podLogs, err := clientset.CoreV1().Pods(client.Namespace).GetLogs(client.PodName, &v1.PodLogOptions{Follow: true, Container: client.ContainerName}).Stream(context.Background())
    if err != nil {
        panic(err.Error())
    }
    defer podLogs.Close()

    buf := make([]byte, 1024)
    counter := 0; 
    for {
        n, _ := podLogs.Read(buf)
        if n == 0 {
            continue
        }
        log := string(buf[:n]);
        counter += 1
        message := &LogMessage{line: counter, message: log}
        c <- *message
    }


}


func getToken(clientset *kubernetes.Clientset, username, password, endpoint string) (string, error) {
    // Create a custom HTTP client with transport to bypass TLS certificate verification
    httpClient := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }

    // Authenticate with the Kubernetes API server using basic authentication
    // and retrieve the token
    fmt.Println(username, password, endpoint)
    tokenResp, err := httpClient.Post(endpoint, "application/x-www-form-urlencoded", strings.NewReader(fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)))
    if err != nil {
        return "", err
    }
    defer tokenResp.Body.Close()

    body, err := ioutil.ReadAll(tokenResp.Body)
    if err != nil {
        return "", err
    }

    token := strings.TrimSpace(string(body))
    fmt.Println(token)

    return token, nil
}
