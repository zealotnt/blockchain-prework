package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is the incomplete declaration we need.
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	fmt.Println("******TODO: IMPLEMENT InitBlockchain!******")
	spew.Dump(Blockchain)
	// Fill me in, noble warrior.
}

func NewBlock(oldBlock Block, data string) Block {
	fmt.Println("******TODO: IMPLEMENT NewBlock!******")
	return Block{}
}

func AddBlock(b Block) error {
	fmt.Println("******TODO: IMPLEMENT AddBlock!******")
	spew.Dump(Blockchain)
	// Fill me in, brave wizard.
	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
