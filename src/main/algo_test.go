package main

import (
	"fmt"
	"testing"
)

func TestGraphAlgo_Match(t *testing.T) {
	g := MakeAlgo()
	players := []Player{
		{"权", 110},
		{"礼", 100},
		{"一雷", 100},
		{"梓坤", 150},
		{"毅", 100},
		{"明", 120},
		{"健宁", 50},
		{"命文", 50},
	}
	res := g.Match(players)
	for _, r := range res {
		fmt.Printf("%s\t%s\tVS\t%s\t%s\n", r[0], r[1], r[2], r[3])
	}
}
