package models

import block "github.com/dianabejan/blockchain/block"

type Blockchain interface {
	CalculateHash() string
	NextBlock(BPM int) (block.Block, error)
	IsBlockValid(newBlock, oldBlock block.Block) bool
}
