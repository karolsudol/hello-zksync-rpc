package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/zksync-sdk/zksync2-go"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("load env key err: %v", err)
	}

	key := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("private key ECDSA err: %v", err)
	}

	pkBytes := crypto.FromECDSA(privateKey)

	// or from raw PrivateKey bytes
	ethereumSigner, err := zksync2.NewEthSignerFromRawPrivateKey(pkBytes, 280)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// also, init ZkSync Provider, specify ZkSync2 RPC URL (e.g. testnet)
	zkSyncProvider, err := zksync2.NewDefaultProvider("https://zksync2-testnet.zksync.dev")

	// then init Wallet, passing just created Ethereum Signer and ZkSync Provider
	w, err := zksync2.NewWallet(ethereumSigner, zkSyncProvider)

	// init default RPC client to Ethereum node (Goerli network in case of ZkSync2 testnet)
	// ethRpc, err := rpc.Dial("https://goerli.infura.io/v3/<your_infura_node_id>")

	// and use it to create Ethereum Provider by Wallet
	// ethereumProvider, err := w.CreateEthereumProvider(ethRpc)

	calldata := crypto.Keccak256([]byte("greet()"))[:4]
	hash, err := w.Execute(
		common.HexToAddress("0x61893345eE37292bAb3c9e7078010E750dED2F0E"),
		calldata,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx hash", hash)

	// var object Greeter

}
