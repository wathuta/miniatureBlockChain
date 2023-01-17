package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

const targetBits = 24

type BlockChain struct {
	blocks []*Block
}

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

type Block struct {
	Data []byte
	BlockHeader
}
type BlockHeader struct {
	Hash          []byte
	Timestamp     int64
	PrevBlockHash []byte
}

const targetBits = 24

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.BlockHeader.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	Block := &Block{
		Data: []byte(data),
		BlockHeader: BlockHeader{
			Timestamp:     time.Now().Unix(),
			PrevBlockHash: prevBlockHash,
		},
	}
	return Block
}
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newblock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newblock)
}
func NewGenesisBlock() *Block {
	//to do what does the genesis block contain
	return NewBlock("Genesis Block", []byte{})
}
func NewBlockChain() *BlockChain {
	return &BlockChain{blocks: []*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to ivan")
	bc.AddBlock("send 2 more BTC to Ivan")
	for _, block := range bc.blocks {
		fmt.Println(string(block.Data))
	}
}
