package main

import (
	fmt
	strconv
)


type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

func (b *Block) Sethash() {
	fmt.Println("hello")
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.hash = hash[:]
}
