package main

import (
	"fmt"
	"strings"

	"github.com/gavrie/oncrpc/parser"
)

type union struct {
	OriginalName   string
	ArbitratorDecl declaration
	TypeIdentifier identifier
	UnexportedName identifier
	Cases          []caseSpec
}

// 'union' IDENTIFIER unionBody ';'   # typeDefUnion
func (l *myListener) EnterTypeDefUnion(ctx *parser.TypeDefUnionContext) {
	originalName := ctx.IDENTIFIER().GetText()
	l.values[ctx] = union{
		OriginalName:   originalName,
		TypeIdentifier: exportedIdentifier(originalName),
		UnexportedName: unexportedIdentifier(originalName),
	}
}

func (l *myListener) ExitTypeDefUnion(ctx *parser.TypeDefUnionContext) {
	fmt.Printf("%v", l.values[ctx.UnionBody()])
}

// unionBody: 'switch' '(' declaration ')' '{'
//         caseSpec
//         caseSpec*
//         defaultCaseSpec?
//     '}';
func (l *myListener) ExitUnionBody(ctx *parser.UnionBodyContext) {
	u := l.values[ctx.GetParent()].(union)

	u.ArbitratorDecl = l.values[ctx.Declaration()].(declaration)
	u.ArbitratorDecl.structName = "s"

	for _, decl := range ctx.AllCaseSpec() {
		u.Cases = append(u.Cases, l.values[decl].(caseSpec))
	}
	if decl := ctx.DefaultCaseSpec(); decl != nil {
		u.Cases = append(u.Cases, l.values[decl].(caseSpec))
	}

	l.values[ctx] = strings.Join([]string{
		u.fieldDeclarations(),
		u.makeEncode(),
		u.makeDecode(),
	}, "\n")
}

func (u *union) fieldDeclarations() string {
	return fillTemplate(`
		// union {{.OriginalName}}

		type {{.TypeIdentifier}} struct {
			{{.ArbitratorDecl}} // Arbitrator
			Union {{.UnexportedName}}Union
		}

		type {{.UnexportedName}}Union interface {
			is{{.TypeIdentifier}}Union()
		}

		{{range .Cases}}
		{{if not .IsVoidOrZeroLength -}}
		func ({{.Declaration.TypeSpec.Name}}) is{{$.TypeIdentifier}}Union() {}
		{{- end}}
		{{end}}
	`, u)
}

func (u *union) makeEncode() string {
	return fillTemplate(`
		func (s *{{.TypeIdentifier}}) Encode(e *xdr.Encoder) error {
			// Arbitrator
			{{.ArbitratorDecl.Encode}}

			switch {{.ArbitratorDecl.FieldName}} {
			{{range .Cases -}}
				{{- if not .IsDefault -}}
					{{range .Values -}}
			case {{.}}:
					{{end}}
				{{- else -}}
			default:
				{{- end -}}
				{{- if not .IsVoidOrZeroLength -}}
				u, ok := s.Union.({{.Declaration.TypeSpec.Name}})
				if !ok {
					return types.ErrArbitratorValueMismatch
				}
				if err := u.Encode(e); err != nil {
					return err
				}
				{{else}}
				// Void
				{{- end -}}
			{{end}}
			}

			return nil
		}
	`, u)
}

func (u *union) makeDecode() string {
	return fillTemplate(`
		func (s *{{.TypeIdentifier}}) Decode(d *xdr.Decoder) error {
			return nil
		}
	`, u)
}

type caseSpec struct {
	Declaration declaration
	Values      []string
	IsDefault   bool
}

func (cs *caseSpec) IsVoidOrZeroLength() bool {
	return cs.IsVoid() || cs.IsZeroLength()
}

func (cs *caseSpec) IsVoid() bool {
	return cs.Declaration.TypeSpec.Name == ""
}

func (cs *caseSpec) IsZeroLength() bool {
	return cs.Declaration.TypeSpec.Size == "0"
}

// caseSpec: ('case' value ':') ('case' value ':')* declaration ';';
func (l *myListener) ExitCaseSpec(ctx *parser.CaseSpecContext) {
	var vals []string
	for _, v := range ctx.AllValue() {
		val := v.GetText()
		switch val {
		case "TRUE":
			val = "true"
		case "FALSE":
			val = "false"
		default:
			val = string(exportedIdentifier(val))
		}
		vals = append(vals, val)
	}

	l.values[ctx] = caseSpec{
		Declaration: l.values[ctx.Declaration()].(declaration),
		Values:      vals,
	}
}

// defaultCaseSpec: ('default' ':' declaration ';');
func (l *myListener) ExitDefaultCaseSpec(ctx *parser.DefaultCaseSpecContext) {
	l.values[ctx] = caseSpec{
		Declaration: l.values[ctx.Declaration()].(declaration),
		IsDefault:   true,
	}
}
