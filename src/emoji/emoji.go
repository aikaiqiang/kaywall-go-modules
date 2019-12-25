package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/go-emoji/emoji"
)

func main() {
	emo := emoji.NewEmoji()
	fmt.Println(emo.Version)
	for _, i := range emo.ShortCodeList() {
		emo.Println(i)
	}
}
