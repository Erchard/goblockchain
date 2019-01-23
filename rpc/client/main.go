package RpcClient

import (
	"../server"
	"bytes"
	"github.com/gorilla/rpc/v2/json"
	"log"
	"net/http"
)

func Request() {
	url := "http://localhost:1234/rpc"
	args := RpcServer.Args{
		A: 2,
		B: 3,
	}
	message, err := json.EncodeClientRequest("Arith.Multiply", args)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	log.Printf("%s\n", message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error is sending request to %s. %s\n", url, err)
	}
	defer resp.Body.Close()

	var result RpcServer.Result
	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode responce %s.\n", err)
	}
	log.Printf("%d * %d = %d\n", args.A, args.B, result)

}
