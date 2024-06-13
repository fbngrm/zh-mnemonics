package mnemonic

import (
	"fmt"

	"github.com/fbngrm/zh-mnemonics/index"
	"github.com/fbngrm/zh-mnemonics/pinyin"
	"github.com/fbngrm/zh-mnemonics/tone"
)

type Builder struct {
	table *pinyin.Table
	index map[string]string
}

func NewBuilder(path string) (*Builder, error) {
	i, err := index.New(path)
	if err != nil {
		return nil, err
	}
	return &Builder{
		table: pinyin.NewTable(),
		index: i,
	}, nil
}

func (b *Builder) Lookup(s string) string {
	m, _ := b.index[s]
	return m
}

type Mnemonic struct {
	Mnemonic       string
	Pronounciation string
}

func (b *Builder) Get(s string) (Mnemonic, error) {
	loc := tone.GetLocation(s)
	pinyinAlpha := tone.ReplaceToneCharacters(s)
	w, err := b.table.Find(pinyinAlpha)
	if err != nil {
		return Mnemonic{}, err
	}
	return Mnemonic{
		Mnemonic:       fmt.Sprintf("%s %s %s\n", w.Initial, loc, w.Final),
		Pronounciation: w.Pronounciation,
	}, nil
}
