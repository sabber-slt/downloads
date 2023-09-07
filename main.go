package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"github.com/refraction-networking/utls"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	uTLSConf := &utls.Config{
		InsecureSkipVerify: true,
	}
	uTLSClientHelloID := utls.HelloChrome_Auto
	dialConn, err := net.Dial("tcp", "your_websocket_server_address:port")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	uconn := utls.UClient(dialConn, uTLSConf, uTLSClientHelloID)
	err = uconn.Handshake()
	if err != nil {
		log.Fatalf("uTLS handshake failed: %v", err)
	}

	// ایجاد یک اتصال WebSocket با استفاده از تنظیمات TLS
	wsConn, _, err := websocket.NewClient(uconn, "your_websocket_url", http.Header{}, 1024, 1024)
	if err != nil {
		log.Fatalf("WebSocket connection failed: %v", err)
	}
	defer wsConn.Close()

	// ارسال و دریافت پیام‌ها از سرور
	err = wsConn.WriteMessage(websocket.TextMessage, []byte("Hello, server!"))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	_, message, err := wsConn.ReadMessage()
	if err != nil {
		log.Fatalf("Failed to read message: %v", err)
	}
	fmt.Printf("Received: %s\n", message)
}
