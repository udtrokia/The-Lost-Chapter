
package main

import (
	"crypto/sha256"
	"strconv"
	"bytes"
	"time"
)

/*
  区块内容定义
*/
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

/*
  SetHash 通过区块链存储的数据为区块设置哈希值。
*/
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// formaInt converter between int and string
	/* formaInt 将 Int 类型与 string 相互转换 */
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// [][]byte - think through []byte arr
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
	// convert [32]byte to []byte
	/* 将[32]byte 转换为 []byte 的格式*/
}

/* 通过传入数据创建一个新的区块 */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

/* 创建创世区块 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

