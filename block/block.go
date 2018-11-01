package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents the structure of the blockchain composition unit
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// CalculateHash method uses SHA256 to compute the hash of the new added block
func (block Block) CalculateHash() string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	sha := sha256.New()
	sha.Write([]byte(record))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

// NextBlock method creates a new block based on the passed BPM argument and the current block
func (block Block) NextBlock(BPM int) (Block, error) {
	var nextBlock Block

	t := time.Now()
	nextBlock.Index = block.Index + 1
	nextBlock.Timestamp = t.String()
	nextBlock.BPM = BPM
	nextBlock.PrevHash = block.Hash
	nextBlock.Hash = nextBlock.CalculateHash()
	return nextBlock, nil
}

// IsBlockValid checks if the data is the block is valid
func IsBlockValid(newBlock, oldBlock Block) bool {
	if (oldBlock.Index + 1) != newBlock.Index {
		return false
	}

	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}

	if newBlock.CalculateHash() != newBlock.Hash {
		return false
	}
	return true
}
