package models

import block "github.com/dianabejan/blockchain/bpm"

type Blockchain interface {
	CalculateHash() string
	NextBlock(BPM int) (block.Block, error)
	IsBlockValid(newBlock, oldBlock block.Block) bool
	// ReplaceChain(newBlocks, Blockchain []block.Block) bool
}
