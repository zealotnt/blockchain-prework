package main

import (
	"bytes"
	"testing"
)

// This test just verifies that life is ok.
func TestTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}

func TestBlock(t *testing.T) {
	InitBlockchain()
	length := len(Blockchain)
	if length != 1 {
		t.Error("New Blockchain should have length 1, but has length: ", length)
	}
}

func TestNewBlock(t *testing.T) {
	testOldBlock := Block{"test", 0, []byte{}, []byte{}}
	testStr := "testing new block!"
	b := NewBlock(testOldBlock, testStr)
	if b.Data != "testing new block!" {
		t.Error("New Block should have data: ", testStr, "but has data: ", b.Data)
	}
	if bytes.Compare(b.PrevHash, testOldBlock.calculateHash()) != 0 {
		t.Errorf("New Block should have PrevHash: %x, but has PrevHash: %x",
			testOldBlock.calculateHash(), b.PrevHash)
	}
	if bytes.Compare(b.Hash, b.calculateHash()) != 0 {
		t.Errorf("New Block should have Hash: %x, but has Hash: %x",
			b.calculateHash(), b.Hash)
	}
}

func TestAddBlock(t *testing.T) {
	InitBlockchain()
	err := AddBlock(Block{"asdf", 0, []byte{}, []byte{}})
	if err == nil {
		t.Error("Adding invalid block should have thrown an error.")
	}
}
