package blockchainGo

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Block struct
type Block struct {
	Index        int64         `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int64         `json:"proof"`
	PreviousHash string        `json:"previous_hash"`
}

// Transaction struct
type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int64  `json:"amount"`
}

// Blockchain struct
type Blockchain struct {
	chain        []Block
	transactions []Transaction
	nodes        StringSet
}

type blockchainInfo struct {
	Length int     `json:"length"`
	Chain  []Block `json:"chain"`
}

//BlockchainService Interfaces
type BlockchainService interface {
	// Add a new node to the list of nodes
	RegisterNode(address string) bool

	// Determine if a given blockchain is valid
	ValidChain(chain Blockchain) bool

	// This is our Consensus Algorithm, it resolves conflicts by replacing
	// our chain with the longest one in the network.
	ResolveConflicts() bool

	// Create a new Block in the Blockchain
	NewBlock(proof int64, previousHash string) Block

	// Creates a new transaction to go into the next mined Block
	NewTransaction(tx Transaction) int64

	// Returns the last block on the chain
	LastBlock() Block

	// Simple Proof of Work Algorithm:
	// - Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
	// - p is the previous proof, and p' is the new proof
	ProofOfWork(lastProof int64)

	// Validates the Proof: Does hash(lastProof, proof) contain 4 leading zeroes?
	VerifyProof(lastProof, proof int64) bool
}

// NewBlock start a new block
func (bc *Blockchain) NewBlock(proof int64, previousHash string) Block {
	prevHash := previousHash
	if previousHash == "" {
		prevBlock := bc.chain[len(bc.chain)-1]
		prevHash = computeHashForBlock(prevBlock)
	}

	newBlock := Block{
		Index:        int64(len(bc.chain) + 1),
		Timestamp:    time.Now().UnixNano(),
		Transactions: bc.transactions,
		Proof:        proof,
		PreviousHash: prevHash,
	}

	bc.transactions = nil
	bc.chain = append(bc.chain, newBlock)
	return newBlock
}

// NewTransaction of blockcain
func (bc *Blockchain) NewTransaction(tx Transaction) int64 {
	bc.transactions = append(bc.transactions, tx)
	return bc.LastBlock().Index + 1
}

// LastBlock of blockchain
func (bc *Blockchain) LastBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

// ProofOfWork of blockchain
func (bc *Blockchain) ProofOfWork(lastProof int64) int64 {
	var proof int64
	for !bc.ValidProof(lastProof, proof) {
		proof++
	}
	return proof
}

// ValidProof of blockchain
func (bc *Blockchain) ValidProof(lastProof, proof int64) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := ComputeHashSha256([]byte(guess))
	return guessHash[:4] == "0000"
}

// ValidChain check block is valid
func (bc *Blockchain) ValidChain(chain *[]Block) bool {
	lastBlock := (*chain)[0]
	currentIndex := 1
	for currentIndex < len(*chain) {
		block := (*chain)[currentIndex]
		// Check that the hash of the block is correct
		if block.PreviousHash != computeHashForBlock(lastBlock) {
			return false
		}
		// Check that the Proof of Work is correct
		if !bc.ValidProof(lastBlock.Proof, block.Proof) {
			return false
		}
		lastBlock = block
		currentIndex++
	}
	return true
}

// RegisterNode of blockchain
func (bc *Blockchain) RegisterNode(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	return bc.nodes.Add(u.Host)
}

// ResolveConflicts of blockchain
func (bc *Blockchain) ResolveConflicts() bool {
	neighbours := bc.nodes
	newChain := make([]Block, 0)

	// We're only looking for chains longer than ours
	maxLength := len(bc.chain)

	// Grab and verify the chains from all the nodes in our network
	for _, node := range neighbours.Keys() {
		otherBlockchain, err := findExternalChain(node)
		if err != nil {
			continue
		}

		// Check if the length is longer and the chain is valid
		if otherBlockchain.Length > maxLength && bc.ValidChain(&otherBlockchain.Chain) {
			maxLength = otherBlockchain.Length
			newChain = otherBlockchain.Chain
		}
	}
	// Replace our chain if we discovered a new, valid chain longer than ours
	if len(newChain) > 0 {
		bc.chain = newChain
		return true
	}

	return false
}

// NewBlockchain new block
func NewBlockchain() *Blockchain {
	newBlockchain := &Blockchain{
		chain:        make([]Block, 0),
		transactions: make([]Transaction, 0),
		nodes:        NewStringSet(),
	}
	// Initial, sentinel block
	newBlockchain.NewBlock(100, "1")
	return newBlockchain
}

func computeHashForBlock(block Block) string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, block)
	return ComputeHashSha256(buf.Bytes())
}

func findExternalChain(address string) (blockchainInfo, error) {
	response, err := http.Get(fmt.Sprintf("http://%s/chain", address))
	if err == nil && response.StatusCode == http.StatusOK {
		var bi blockchainInfo
		if err := json.NewDecoder(response.Body).Decode(&bi); err != nil {
			return blockchainInfo{}, err
		}
		return bi, nil
	}
	return blockchainInfo{}, err
}
