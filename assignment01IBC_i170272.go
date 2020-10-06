package assignment01IBC


import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func VerifyChain(chainHead *Block) {

	var temp *Block
	//	var temp string
	for temp = chainHead; temp != nil; {
		temp2 := temp.prevPointer
		fmt.Println("transaction=", temp.transactions)
		if temp2 != nil {
			pHash := CalculateHash(temp.prevPointer)
			if pHash != temp.prevHash {
				fmt.Println("change detected")
				return
			}
		}
		temp = temp.prevPointer
	}

}

func CalculateHash(inputBlock *Block) string {
	a := len(inputBlock.transactions)
	var temp string
	for b := 0; b < a; b++ {
		temp = temp + inputBlock.transactions[b]
	}
	fmt.Println(temp)
	//obj := sha256.New()
	obj := sha256.New()
	obj.Write([]byte(fmt.Sprintf("%x", temp)))

	return fmt.Sprintf("%x", obj.Sum(nil))
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
	//	var NewBlock *Person = new(Person)

	//NewBlock.Name =

	if chainHead == nil {

		var temp *Block = new(Block)

		temp.transactions = transactionsToInsert
		temp.prevPointer = nil
		hash := CalculateHash(temp)
		temp.currentHash = hash
		return temp

	}

	if chainHead != nil {
		//fmt.Println(" not genesis block ")
		var temp *Block = new(Block)

		temp.transactions = transactionsToInsert

		temp.prevPointer = chainHead

		temp.prevHash = chainHead.currentHash
		temp.currentHash = CalculateHash(temp)
		fmt.Println("previous block content :-", temp.prevPointer.transactions)
		return temp

	}
	return nil
	//fmt.Println("____________________")
}

func ListBlocks(chainHead *Block) {
	//a := len(chainHead.transactions)
	var temp *Block
	//	var temp string
	for temp = chainHead; temp != nil; {
		fmt.Println("transaction=", temp.transactions)
		fmt.Println("--------")
		temp = temp.prevPointer
	}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	//tempPerson := chainHead
	var temp *Block

	for temp = chainHead; temp != nil; {

		a := len(temp.transactions)

		for b := 0; b < a; b++ {
			if temp.transactions[b] == oldTrans {
				fmt.Println("found and updated")
				temp.transactions[b] = newTrans
				//tempPerson.CurrentHash = CalculateHash(chainHead)
				//fmt.Println("hash updated=", tempPerson.CurrentHash, "         ", tempPerson.Transactions)
				return
				// Found!
			}

		}
		temp = temp.prevPointer

	}
}