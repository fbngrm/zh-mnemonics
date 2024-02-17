package main

import (
	"log"

	"github.com/fbngrm/zh-mnemonics/mnemonic"
	"github.com/fbngrm/zh-mnemonics/pinyin"
)

func main() {
	p := "nan1"
	b := mnemonic.NewBuilder(pinyin.NewTable())
	s, err := b.GetBase(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

}
