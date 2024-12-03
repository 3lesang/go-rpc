package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)

var rpcMethods = map[string]func(params interface{}) (interface{}, error){
	"sayHello": func(params interface{}) (interface{}, error) {
		return "Hello, " + params.(string), nil
	},
}

func handleRPC(request RPCRequest) RPCResponse {
	handler, exists := rpcMethods[request.Method]
	if !exists {
		return RPCResponse{Error: "method not found"}
	}

	result, err := handler(request.Params)
	if err != nil {
		return RPCResponse{Error: err.Error()}
	}

	return RPCResponse{Result: result}
}

func websocketHandler(c *websocket.Conn) {
	defer c.Close()

	for {
		// Read message
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// Decode RPC request
		var request RPCRequest
		if err := json.Unmarshal(msg, &request); err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(`{"error":"invalid request"}`))
			continue
		}

		// Process request and encode response
		response := handleRPC(request)
		responseBytes, _ := json.Marshal(response)

		// Send response
		if err = c.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
