package main

import (
	blockchain "blockchaingo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	serverPort := flag.String("port", "8000", "http port number where server will run")
	flag.Parse()

	blkc := blockchain.NewBlockchain()
	nodeID := strings.Replace(blockchain.PseudoUUID(), "-", "", -1)

	fmt.Printf("%s\n\n", nodeID)

	log.Printf("Starting Blockchain HTTP Server. Listening at port %q", *serverPort)

	http.Handle("/", blockchain.NewHandler(blkc, nodeID))
	http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), nil)
}
