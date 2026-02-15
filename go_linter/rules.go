package go_linter

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"strings"
)

var sensitiveWords = [3]string{"password", "apikey", "token"}

func checkSensitiveWords(s string) bool {
	for _, word := range sensitiveWords {
		if strings.Contains(s, word) {
			return true
		}
	}
	return false
}

func checkLogRules(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	arg, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return
	}

	if arg.Kind != token.STRING {
		return
	}

	msg := strings.Trim(arg.Value, `""`)
	if len(msg) == 0 {
		return
	}

	if msg[0] >= 'A' && msg[0] <= 'Z' {
		pass.Reportf(arg.Pos(), "log messages shouldn't start with a capitalized letter.")
	}

	for _, l := range msg {
		if !(l >= 'a' && l <= 'z') {
			pass.Reportf(arg.Pos(), "log messages should be written in English.")
		}
	}

	if checkSensitiveWords(msg) {
		pass.Reportf(arg.Pos(), "log message potentionally contains sensitive information.")
	}
}
