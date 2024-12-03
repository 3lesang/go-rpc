package main

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RPCResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}
