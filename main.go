package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gavrie/oncrpc/parser"
	"github.com/pboyer/antlr4/runtime/Go/antlr"
)

func main() {
	filename := os.Args[1]
	packageName := strings.Split(filepath.Base(filename), ".")[0]

	input := antlr.NewFileStream(filename)
	lexer := parser.NewONCRPCv2Lexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewONCRPCv2Parser(tokens)
	tree := p.Oncrpcv2Specification()

	l := newMyListener(packageName)
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
}
