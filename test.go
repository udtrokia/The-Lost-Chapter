package main

import (
	"strconv"
	"fmt"
	"time"
	"bytes"
	"crypto/sha256"
)

//strconv.formaInt
func strconv_test() {
	// int64
	//raw := time.Now().Unix()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// converter between int64 and 
	fmt.Printf(timestamp)
	fmt.Println()

}

func bytes_sha256_test() {	
	headers := bytes.Join([][]byte{[]byte("eric"), []byte("udtrokia"), []byte("mercurio")}, []byte{})
	// [][]byte: try `[]btye` instead of `[][]byte` ERR occurs below, you know what I mean.
	// cannot use []byte literal (type []byte) as type [][]byte in argument to bytes.Join
	// cannot use ([]byte)("eric") (type []byte) as type byte in array or slice literal
	// cannot use ([]byte)("udtrokia") (type []byte) as type byte in array or slice literal
	// cannot use ([]byte)("mercurio") (type []byte) as type byte in array or slice literal
	
	hash := sha256.Sum256(headers)
	fmt.Printf("%x\n",hash)
	// pay attention to this. if you fmt.Printf(hash), err will come below
	// ERR: cannot use hash (type [32]byte) as type string in argument to fmt.Printf

	fmt.Printf("%x\n",hash[:])
	// hash[:]	
	// convert [32]byte to []byte: in the block.go file, leave out [:], ERR occurs below
	// cannot use hash (type [32]byte) as type []byte in assignment
}

