package main

import "bytes"

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{
		b.Data,
		b.PrevHash,
	}, []byte{})

}
