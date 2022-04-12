package main

import (
	"github.com/hashicorp/vault/api"
	"github.com/anonsachin/fabric-sdk-go/pkg/gateway"
)

func main(){
	// importing vault
	_ = api.DefaultConfig()

	// importing gateway
	_ = gateway.NewInMemoryWallet()
}