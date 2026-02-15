package go_linter

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "check log messages for stylistic errors",
	Run:  run,
}

func isLogFunc(pass *analysis.Pass, call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	obj := pass.TypesInfo.Uses[sel.Sel]
	if obj == nil {
		return false
	}
	if obj.Pkg() == nil {
		return false
	}

	return obj.Pkg().Name() == "log"
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if isLogFunc(pass, call) {
				checkLogRules(pass, call)
			}

			return true
		})
	}
	return nil, nil
}
