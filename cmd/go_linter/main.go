package main

import (
	"github.com/MaximKovyrzin/go-linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(go_linter.Analyzer)
}
