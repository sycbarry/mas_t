package client

import (
	b64 "encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/go-stomp/stomp"
	"github.com/gorilla/websocket"
)


type Client struct {
    Username string
    Password string 
    Endpoint string
}



func (c *Client) Build() (conn stomp.Conn) {

    basicHeader := BasicHeader{Username: c.Username, Password: c.Password}
    basicHeader.Encode()
    wsHeaders := http.Header{}
    wsHeaders.Add("Authorization", basicHeader.GetHeader() )

    wsConn, _, err := websocket.DefaultDialer.Dial(c.Endpoint, wsHeaders)
    if err != nil {
        log.Fatalf("Error while connecting to WebSocket endpoint %s: %v\n", c.Endpoint, err)
    }
    //defer wsConn.Close()

    wsReadWriteCloser := &WebSocketReadWriteCloser{conn: wsConn}
    stompConn, err := stomp.Connect(wsReadWriteCloser)
    if err != nil {
        log.Fatalf("Error while connecting to STOMP server: %v\n", err)
    }
    //defer stompConn.Disconnect();

    return *stompConn;

}



type BasicHeader struct {
    Username string 
    Password string
    B64EncodedPassword string
}


func (b *BasicHeader) Encode() {
    value := b64.StdEncoding.EncodeToString([]byte(b.Username + ":" + b.Password))
    b.B64EncodedPassword = value
}


func (b* BasicHeader) GetHeader() (value string) {
    return "Basic " + b.B64EncodedPassword
}


// WebSocketReadWriteCloser wraps a websocket connection to implement io.ReadWriteCloser.
type WebSocketReadWriteCloser struct {
    conn *websocket.Conn
}


// Read reads data from the websocket connection.
func (w *WebSocketReadWriteCloser) Read(p []byte) (n int, err error) {
    messageType, message, err := w.conn.ReadMessage()
    if err != nil {
        return 0, err
    }
    if messageType != websocket.TextMessage {
        return 0, io.ErrUnexpectedEOF
    }
    copy(p, message)
    return len(message), nil
}


// Write writes data to the websocket connection.
func (w *WebSocketReadWriteCloser) Write(p []byte) (n int, err error) {
    err = w.conn.WriteMessage(websocket.TextMessage, p)
    if err != nil {
        return 0, err
    }
    return len(p), nil
}


// Close closes the websocket connection.
func (w *WebSocketReadWriteCloser) Close() error {
    return w.conn.Close()
}

