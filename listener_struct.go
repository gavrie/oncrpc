package main

import (
	"fmt"
	"strings"

	"github.com/gavrie/oncrpc/parser"
)

type struct_ struct {
	OriginalName string
	Identifier   identifier
	Fields       []declaration
}

// 'struct' IDENTIFIER structBody ';' # typeDefStruct
func (l *myListener) EnterTypeDefStruct(ctx *parser.TypeDefStructContext) {
	l.values[ctx] = struct_{
		OriginalName: ctx.IDENTIFIER().GetText(),
		Identifier:   exportedIdentifier(ctx.IDENTIFIER().GetText()),
	}
}

func (l *myListener) ExitTypeDefStruct(ctx *parser.TypeDefStructContext) {
	st := l.values[ctx].(struct_)
	fmt.Printf(`
		// struct %v
		type %v struct %v
		`,
		st.OriginalName,
		st.Identifier,
		l.values[ctx.StructBody()],
	)
}

// structBody: '{' (declaration ';') (declaration ';')* '}';
func (l *myListener) ExitStructBody(ctx *parser.StructBodyContext) {
	st := l.values[ctx.GetParent()].(struct_)

	for _, decl := range ctx.AllDeclaration() {
		fieldDecl := l.values[decl].(declaration)
		fieldDecl.structName = "s"
		st.Fields = append(st.Fields, fieldDecl)
	}
	l.values[ctx] = strings.Join([]string{
		st.fieldDeclarations(),
		st.makeEncode(),
		// st.makeDecode(),
	}, "\n")

}

func (s struct_) fieldDeclarations() string {
	return fillTemplate(`
		{
			{{range .Fields -}}
			{{.}}
			{{end}}
		}
		`, s)
}

func (s struct_) makeEncode() string {
	return fillTemplate(`
		func (s *{{.Identifier}}) Encode(e *xdr.Encoder) error {
			{{range .Fields -}}
			{{.Encode}}
			{{end}}
			return nil
		}
		`, s)
}

func (s struct_) makeDecode() string {
	return fillTemplate(`
		func (s *{{.Identifier}}) Decode(d *xdr.Decoder) error {
			{{range .Fields -}}
			{{.Decode}}
			{{end}}
			return nil
		}
		`, s)
}
