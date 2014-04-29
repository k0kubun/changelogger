package main

import (
	"fmt"
	"github.com/agtorre/gocolorize"
	"github.com/sergi/go-diff/diffmatchpatch"
	"strings"
)

func showDiff(oldText, newText string) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldText, newText, true)

	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			color := gocolorize.NewColor("green")
			printLinesWithPrefix(color.Paint("+"), diff.Text)
		case diffmatchpatch.DiffDelete:
			color := gocolorize.NewColor("red")
			printLinesWithPrefix(color.Paint("-"), diff.Text)
		}
	}
}

func printLinesWithPrefix(prefix string, text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			fmt.Printf("%s ", prefix)
			fmt.Println(line)
		}
	}
}
