package main

import (
	blkchain "blockchainGo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	serverPort := flag.String("port", "8000", "http port number where server will run")
	flag.Parse()

	blockchain := blkchain.NewBlockchain()
	nodeID := strings.Replace(blkchain.PseudoUUID(), "-", "", -1)

	log.Printf("Starting blkchain HTTP Server. Listening at port %q", *serverPort)

	http.Handle("/", blkchain.NewHandler(blockchain, nodeID))
	http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), nil)
}
