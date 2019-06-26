package blockchain

import (
	"errors"
	"sync"
)

//NewChain : create a new chain
func NewChain(blocks []Block, dir string) Chain {
	var c chain
	c.chain = blocks
	c.dir = dir
	return &c

}

//Chain type for chain and methods
type Chain interface {
	Get() []Block
	GetBlock(int) (Block, error)
	Replace(Chain) bool
	Append(Block) bool
	Lock()
	RLock()
	Unlock()
	RUnlock()
}
type chain struct {
	sync.RWMutex
	chain []Block
	dir   string
}

func (c *chain) Get() []Block {
	return c.chain
}
func (c *chain) GetBlock(i int) (Block, error) {
	var b Block
	c.RLock()
	if i >= len(c.chain) {
		c.RUnlock()
		return b, errors.New("cant find block")
	}
	c.RUnlock()
	return c.chain[i], nil
}
func (c *chain) Replace(newC Chain) bool {
	c.Lock()
	if len(newC.Get()) > len(c.chain) {
		c.chain = newC.Get()
		c.Unlock()
		return true
	}
	c.Unlock()
	return false
}
func (c *chain) Append(b Block) bool {
	c.Lock()
	c.chain = append(c.chain, b)
	c.Unlock()
	return true
}
