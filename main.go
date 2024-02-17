package main

import (
	"fbngrm/zh-mnemo/mnemonic"
	"fbngrm/zh-mnemo/pinyin"
	"log"
)

func main() {
	p := "nàn"
	b := mnemonic.NewBuilder(pinyin.NewTable())
	s, err := b.GetBase(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

}
