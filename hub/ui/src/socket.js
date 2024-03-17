

const userName = "client";
const password = "password";
const endpoint = "localhost"; 
const port = 8080;

async function listClients(callback) {
    const combo = userName + ":" + password
    const token = btoa(combo);
    let headers = new Headers()
    headers.append("Authorization", "Basic " + token);
    headers.append("Access-Control-Allow-Origin", "*");
    fetch("http://127.0.0.1:8080/users", { headers: headers, method: "GET" })
    .then(response => response.json())
    .then(data => callback(data))
    .catch(e => console.log(e))
}

function getClientName() {
    const urlParams = new URLSearchParams(window.location.search);
    const node = urlParams.get('node');
    return node;
}

function connect(Stomp) {
    let topic = "/topic/log/"
    const urlParams = new URLSearchParams(window.location.search);
    const node = urlParams.get('node');
    topic += node;

    var url = "ws://localhost:8080/ws"

    var client = Stomp.over(new WebSocket(url));

    const combo = userName + ":" + password
    const token = btoa(combo);

    var counter = 0;
    var onMessage = function(message) {
        var body = message.body;
        var child = document.createElement("div"); 
        child.style.color = "black";
        child.innerText = body;
        child.style.fontSize = "10px";
        child.style.padding = "5px";
        if(counter % 2 == 0) {
            child.style.backgroundColor = "lightgray"
        } else {
            child.style.backgroundColor = "white"
        }
        document.getElementById("messages").appendChild(child);
        child.scrollIntoView();
        counter += 1;
    };


    const headers = { "Authorization": "Basic " + token }

    client.connect(headers, function(value) {
        console.log('Connected to STOMP server');
        client.subscribe(topic, onMessage);
    });

}


export {
    connect,
    listClients,
    getClientName
}
