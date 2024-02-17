package mnemonic

import (
	"fmt"

	"github.com/fbngrm/zh-mnemonics/pinyin"
	"github.com/fbngrm/zh-mnemonics/tone"
)

type Builder struct {
	table *pinyin.Table
}

func NewBuilder(t *pinyin.Table) *Builder {
	return &Builder{
		table: t,
	}
}

func (b *Builder) GetBase(s string) (string, error) {
	loc := tone.GetLocation(s)
	pinyinAlpha := tone.ReplaceToneCharacters(s)
	w, err := b.table.Find(pinyinAlpha)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s %s\n", w.Initial, loc, w.Final), nil
}
