package main

import (
	"fmt"
	"strings"

	"github.com/gavrie/oncrpc/parser"
)

type typedef struct {
	Declaration *declaration
}

// 'typedef' declaration ';'          # typeDefTypedef
func (l *myListener) ExitTypeDefTypedef(ctx *parser.TypeDefTypedefContext) {
	decl := l.values[ctx.Declaration()].(declaration)
	decl.customName = "value"
	td := typedef{
		Declaration: &decl,
	}

	fmt.Printf("type %v\n", strings.Join([]string{
		td.Declaration.String(),
		td.makeEncode(),
		// td.makeDecode(),
	}, "\n"))
}

func (td *typedef) makeEncode() string {
	return fillTemplate(`
		func (t *{{.Identifier}}) Encode(e *xdr.Encoder) error {
			value := {{.TypeSpec.Name}}(*t)
			{{.Encode}}
			return nil
		}
		`, td.Declaration)
}

func (td *typedef) makeDecode() string {
	return fillTemplate(`
		func (t *{{.Identifier}}) Decode(d *xdr.Decoder) error {
			{{.Decode}}
			return nil
		}
		`, td.Declaration)
}
