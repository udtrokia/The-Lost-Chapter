package main

/* 定义 Blockchain 的类型 */
type Blockchain struct {
	blocks []*Block
}

// addblocks
/* 添加区块函数 */
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

/* 创建新的区块链 */
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
