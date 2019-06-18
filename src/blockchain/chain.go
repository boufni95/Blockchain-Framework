package blockchain

import (
	"errors"
)

//NewChain : create a new chain
func NewChain(blocks []Block, dir string) Chain {
	return &chain{
		blocks,
		dir,
	}

}

//Chain type for chain and methods
type Chain interface {
	Get() []Block
	GetBlock(int) (Block, error)
	Replace(Chain) bool
	Append(Block) bool
}
type chain struct {
	chain []Block
	dir   string
}

func (c *chain) Get() []Block {
	return c.chain
}
func (c *chain) GetBlock(i int) (Block, error) {
	var b Block
	if i >= len(c.chain) {
		return b, errors.New("cant find block")
	}
	return c.chain[i], nil
}
func (c *chain) Replace(newC Chain) bool {
	if len(newC.Get()) > len(c.chain) {
		c.chain = newC.Get()
		return true
	}
	return false
}
func (c *chain) Append(b Block) bool {
	c.chain = append(c.chain, b)
	return true
}
