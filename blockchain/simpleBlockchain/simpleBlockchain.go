package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

var Blockchain []Block

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

func main() {
	//Создаем первый блок
	t := time.Now()
	genesisBlock := Block{0, t.String(), "init", "", ""}
	Blockchain = append(Blockchain, genesisBlock)
	time.Sleep(1 * time.Second)
	//Записываем новый блок
	addBlock(Blockchain[len(Blockchain)-1], "two block")
	time.Sleep(1 * time.Second)
	//Записываем новый блок
	addBlock(Blockchain[len(Blockchain)-1], "three block")
	time.Sleep(1 * time.Second)
	//Записываем невалидный блок
	addBlock(Blockchain[len(Blockchain)-2], "false block")

	for _, block := range Blockchain {
		fmt.Printf("Index: %x\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Data string) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func addBlock(oldBlock Block, data string) {
	newBlock := generateBlock(oldBlock, data)
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
	}

}

func isBlockValid(newBlock, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		fmt.Println("error: ", "oldBlock.Index+1 != newBlock.Index")
		return false
	} else if oldBlock.Hash != newBlock.PrevHash {
		fmt.Println("error: ", "oldBlock.Hash != newBlock.PrevHash")
		return false
	} else if calculateHash(newBlock) != newBlock.Hash {
		fmt.Println("error: ", "calculateHash(newBlock) != newBlock.Hash")
		return false
	}
	return true
}
