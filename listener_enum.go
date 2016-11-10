package main

import (
	"fmt"
	"strings"

	"github.com/gavrie/oncrpc/parser"
)

type enum struct {
	OriginalName string
	Identifier   identifier
	Items        []enumItem
}

type enumItem struct {
	Key   identifier
	Value string
}

// 'enum' IDENTIFIER enumBody ';'     # typeDefEnum
func (l *myListener) EnterTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	l.values[ctx] = enum{
		OriginalName: ctx.IDENTIFIER().GetText(),
		Identifier:   exportedIdentifier(ctx.IDENTIFIER().GetText()),
	}
}

func (l *myListener) ExitTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	enum := l.values[ctx].(enum)
	fmt.Printf(`
		// enum %v
		type %v enum
		%v
		`,
		enum.OriginalName,
		enum.Identifier,
		l.values[ctx.EnumBody()],
	)
}

// enumBody: '{' (IDENTIFIER '=' value) (',' IDENTIFIER '=' value)* '}';
func (l *myListener) ExitEnumBody(ctx *parser.EnumBodyContext) {
	enum := l.values[ctx.GetParent()].(enum)

	vals := ctx.AllValue()
	for i, key := range ctx.AllIDENTIFIER() {
		enum.Items = append(enum.Items, enumItem{
			Key:   exportedIdentifier(key.GetText()),
			Value: vals[i].GetText(),
		})
	}

	l.values[ctx] = strings.Join([]string{
		enum.constDeclarations(),
		enum.makeString(),
		enum.makeFieldCheck(),
		enum.makeEncode(),
		enum.makeDecode(),
	}, "\n")
}

func (e enum) constDeclarations() string {
	return fillTemplate(`
		const (
			{{range .Items -}}
			{{.Key}} {{$.Identifier}} = {{.Value}}
			{{end}}
		)
		`, e)
}

func (e enum) makeString() string {
	return fillTemplate(`
		func (v {{.Identifier}}) String() string {
			switch v {
			{{range .Items -}}
			case {{.Key}}:
				return "{{.Key}}"
			{{end -}}
			}
			return fmt.Sprintf("{{.Identifier}}(%d)", v)
		}
		`, e)
}

func (e enum) makeFieldCheck() string {
	return fillTemplate(`
		func (v {{.Identifier}}) fieldCheck() error {
			switch v {
			{{range .Items -}}
			case {{.Key}}:
			{{end -}}
			default:
				return types.ErrInvalidEnumOpt
			}
			return nil
		}
		`, e)
}

func (e enum) makeEncode() string {
	return fillTemplate(`
		func (v *{{.Identifier}}) Encode(e *xdr.Encoder) error {
			if err := v.fieldCheck(); err != nil {
				return err
			}
			if _, err := e.EncodeInt(int32(*v)); err != nil {
				return err
			}
			return nil
		}
		`, e)
}

func (e enum) makeDecode() string {
	return fillTemplate(`
		func (v *{{.Identifier}}) Decode(d *xdr.Decoder) error {
			intValue, _, err := d.DecodeInt()
			if err != nil {
				return err
			}
			*v = {{.Identifier}}(intValue)
			return v.fieldCheck()
		}
		`, e)
}
