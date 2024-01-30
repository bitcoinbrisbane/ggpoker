package main

import (
	// "crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anthdm/ggpoker/p2p"
	"github.com/miguelmota/go-ethereum-hdwallet"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto"
)

func makeServerAndStart(addr, apiAddr string) *p2p.Server {
	cfg := p2p.ServerConfig{
		Version:       "GGPOKER V0.3-alpha",
		ListenAddr:    addr,
		APIListenAddr: apiAddr,
		GameVariant:   p2p.TexasHoldem,
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	time.Sleep(time.Millisecond * 200)

	return server
}

// // Each player has a private key and a public key
// func makeKeyPair() (string) {
// 	privateKey, err := crypto.GenerateKey()
//     if err != nil {
//         fmt.Println(err)
//     }

//     privateKeyBytes := crypto.FromECDSA(privateKey)
//     fmt.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))

//     publicKey := privateKey.Public()
//     publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//     if !ok {
//         fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//     }

//     publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
//     fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

//     address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
//     fmt.Println("Address:", address)

// 	// return hexutil.Encode(privateKeyBytes), address
// 	return address
// }

func makeKeyPair() (string) {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559

	return account.Address.Hex()
}

func main() {
	addr := ":3000"
	apiAddr := ":3001"

	if len(os.Args) > 1 {
		addr = os.Args[0]
		apiAddr = os.Args[1]
	}

	player := makeServerAndStart(addr, apiAddr) // dealer
	fmt.Println("Player :", player.ListenAddr)

	// playerB := makeServerAndStart(":4000", ":4001") // sb
	// playerC := makeServerAndStart(":5000", ":5001") // bb
	// playerD := makeServerAndStart(":7000", ":7001") // bb + 2

	go func() {
		time.Sleep(time.Second * 2)
		http.Get(fmt.Sprintf("http://localhost%a/ready", apiAddr))

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:4001/ready")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:5001/ready")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:7001/ready")

		// [3000:D, 4000:sb, 5000:bb, 7000]
		// PREFLOP
		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:4001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:5001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:7001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:3001/fold")

		// // FLOP
		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:4001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:5001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:7001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:3001/fold")

		// // TURN
		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:4001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:5001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:7001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:3001/fold")

		// // RIVER
		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:4001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:5001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:7001/fold")

		// time.Sleep(time.Second * 2)
		// http.Get("http://localhost:3001/fold")

	}()

	// time.Sleep(time.Millisecond * 200)
	// playerB.Connect(playerA.ListenAddr)

	// time.Sleep(time.Millisecond)
	// playerC.Connect(playerB.ListenAddr)

	// time.Sleep(time.Millisecond * 200)
	// playerD.Connect(playerC.ListenAddr)

	select {}
}

func init() {
	// Generate a key pair for the dealer
}
