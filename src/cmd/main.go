package main

import (
	blk "blockchaingo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	serverPort := flag.String("port", "8000", "http port number where server will run")
	flag.Parse()

	blockchain := blk.NewBlockchain()
	nodeID := strings.Replace(blk.PseudoUUID(), "-", "", -1)

	log.Printf("Starting blk HTTP Server. Listening at port %q", *serverPort)

	http.Handle("/", blk.NewHandler(blockchain, nodeID))
	http.ListenAndServe(fmt.Sprintf(":%s", *serverPort), nil)
}
