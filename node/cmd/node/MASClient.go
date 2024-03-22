package node

import (
	"context"
    "net/http"
    "crypto/tls"

    "strings"
    "io/ioutil"
    "io"
    "fmt"


	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MASClient struct {
    ConfigPath string
    PodName string
    Namespace string
    Username string 
    Password string 
    Host string
    Config *rest.Config
}

func containsSubstring(s string, substrings []string) bool {
    for _, substr := range substrings {
        if strings.Contains(s, substr) && ! strings.Contains(s, "build") {
            return true;
        }
    }
    return false;
}

func (client *MASClient) ListBindablePods() ([]string, error) {
    clientset, err := kubernetes.NewForConfig(client.Config)
    if err != nil {
        panic(err.Error())
    }
    podList, err := clientset.CoreV1().Pods(client.Namespace).List(context.Background(), metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }

    substrings := []string{"mea", "jms", "ui", "mxinst", "cron", "rpt"}
    var podClients []string

    for _, pod := range podList.Items {
        if containsSubstring(pod.GetName(), substrings) {
            podClients = append(podClients, pod.GetName())
        }
    }

    return podClients, nil
}

func (client *MASClient) BuildConfig() {
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
    client.Config = config;
}

func (client *MASClient) ReadLogs(c chan LogMessage) {

    clientset, err := kubernetes.NewForConfig(client.Config)
    if err != nil {
        panic(err.Error())
    }

    pod, err := clientset.CoreV1().Pods(client.Namespace).Get(context.Background(), client.PodName, metav1.GetOptions{})
    if err != nil {
        panic(err.Error())
    }

    var container string;
    var podLogs io.ReadCloser
    var seconds int64 = 10;

    if len(pod.Spec.Containers) > 0{
        container = pod.Spec.Containers[0].Name
        podLogs, err = clientset.CoreV1().Pods(client.Namespace).GetLogs(client.PodName, &v1.PodLogOptions{ Follow: true, Container: container, Timestamps: true, SinceSeconds: &seconds }).Stream(context.Background())
    } else {
        podLogs, err = clientset.CoreV1().Pods(client.Namespace).GetLogs(client.PodName, &v1.PodLogOptions{ Follow: true, Timestamps: true, SinceSeconds: &seconds }).Stream(context.Background())
    }

    fmt.Println(client.PodName);


    if err != nil {
        panic(err.Error())
    }
    defer podLogs.Close()

    buf := make([]byte, 2048)
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
