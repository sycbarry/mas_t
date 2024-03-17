package node

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type MASClient struct {
    ConfigPath string
    PodName string
    Namespace string
}

func (client *MASClient) ReadLogs(c chan LogMessage) {
    config, err := rest.InClusterConfig()
    if err != nil {
        filePath := "/Users/sebastianbarry/.kube/zingg-prod-kubeconfig.yaml"
        config, err = clientcmd.BuildConfigFromFlags("", filePath)
        if err != nil {
            panic(err.Error())
        }
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
    
    podLogs, err := clientset.CoreV1().Pods(client.Namespace).GetLogs(client.PodName, &v1.PodLogOptions{Follow: true}).Stream(context.Background())
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

