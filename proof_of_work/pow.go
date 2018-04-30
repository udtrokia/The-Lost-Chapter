package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 24

// ProofOfWork represents a proof-of-work
/* 工作证明资产 */
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
/* ProofOfWork 生成函数 */
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// NewInt(x) returns an *Int set to the value of the int64 argument x
	/* NewInt 用来设置 int64 的值 */
	
	target.Lsh(target, uint(256-targetBits))
	// target means total workload
	// targetBits means workload 
	// Lsh sets z: z = x << n and returnd z
	////// is targetBits bigger, targets smaller? //////
	// 2 << 3 == 16
	/* target 代表当前 POW 难度*/
	
	pow := &ProofOfWork{b, target}
	// target is the size of a block
	
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		// this is a block
		[]byte{},
	)
	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
		// Cmp compares x and y and returns
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
