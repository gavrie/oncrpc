package main

import (
	"fmt"

	"github.com/gavrie/oncrpc/parser"
)

// declaration

type declaration struct {
	Identifier identifier
	TypeSpec   typeSpec
	InnerDecl  *declaration
	Comment    string
	structName string
	customName string
}

func (d *declaration) String() string {
	return fmt.Sprintf("%v %v%v", d.Identifier, d.TypeSpec.Name, d.Comment)
}

func (d *declaration) FieldName() string {
	if d.structName != "" {
		return fmt.Sprintf("%v.%v", d.structName, d.Identifier)
	}
	if d.customName != "" {
		return d.customName
	}
	return string(d.Identifier)
}

func (d *declaration) Encode() string {
	var enc string
	switch d.TypeSpec.Kind {
	case kindInt32:
		enc = `if _, err := e.EncodeInt({{.FieldName}}); err != nil { return err }`
	case kindUInt32:
		enc = `if _, err := e.EncodeUint({{.FieldName}}); err != nil { return err }`
	case kindInt64:
		enc = `if _, err := e.EncodeHyper({{.FieldName}}); err != nil { return err }`
	case kindUInt64:
		enc = `if _, err := e.EncodeUhyper({{.FieldName}}); err != nil { return err }`
	case kindBool:
		enc = `if _, err := e.EncodeBool({{.FieldName}}); err != nil { return err }`
	case kindString:
		enc = `
			// {{.FieldName}}: {{.TypeSpec.Name}} <{{.TypeSpec.Size}}>
			if len({{.FieldName}}) > {{.TypeSpec.Size}} {
				return types.ErrArrayTooLarge
			}
			if _, err := e.EncodeString({{.FieldName}}); err != nil {
				return err
			}
		`
	case kindByteSlice:
		enc = `
			// {{.FieldName}}: {{.TypeSpec.Name}} <{{.TypeSpec.Size}}>
			if len({{.FieldName}}) > {{.TypeSpec.Size}} {
				return types.ErrArrayTooLarge
			}
			if _, err := e.EncodeOpaque({{.FieldName}}); err != nil {
				return err
			}
		`
	case kindByteArray:
		enc = `
			// {{.FieldName}}: {{.TypeSpec.Name}} <{{.TypeSpec.Size}}>
			if _, err := e.EncodeFixedOpaque({{.FieldName}}[:]); err != nil {
				return err
			}
		`
	case kindSlice:
		enc = `
			// {{.FieldName}}: {{.TypeSpec.Name}} <{{.TypeSpec.Size}}>
			{
				dataLength := uint32(len({{.FieldName}}))

				if dataLength > {{.TypeSpec.Size}} {
					return types.ErrArrayTooLarge
				}

				if _, err := e.EncodeUint(dataLength); err != nil {
					return err
				}

				for _, value := range {{.FieldName}} {
					{{.InnerDecl.Encode}}
				}
			}
		`
	case kindIdentifier:
		enc = `
			if err := {{.FieldName}}.Encode(e); err != nil {
				return err
			}
		`
	case kindOptional:
		enc = `
			// {{.FieldName}}: {{.TypeSpec.Name}}
			if {{.FieldName}} != nil {
				if _, err := e.EncodeBool(true); err != nil {
					return err
				}
				if err := {{.FieldName}}.Encode(e); err != nil {
					return err
				}
			} else {
				if _, err := e.EncodeBool(false); err != nil {
					return err
				}
			}
		`
	default:
		panic(fmt.Errorf("unsupported type: %v", d.TypeSpec.Kind))
	}
	return fillTemplate(enc, d)
}

/*
func (d *declaration) Decode() string {
	if d.TypeSpec.Kind == kindIdentifier {
		return fillTemplate(`
		if err := s.{{.Identifier}}.Decode(d); err != nil {
			return err
		}
	`, d)
	}
	return d.TypeSpec.Decode(d.Identifier)
}
*/

// typeSpecifier IDENTIFIER                # declSimple
func (l *myListener) ExitDeclSimple(ctx *parser.DeclSimpleContext) {
	typeSpec := l.values[ctx.TypeSpecifier()].(typeSpec)
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())
	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec:   typeSpec,
	}
}

// typeSpecifier IDENTIFIER '[' value ']'  # declFixed
func (l *myListener) ExitDeclFixed(ctx *parser.DeclFixedContext) {
	var size string
	if ctx.Value() != nil {
		size = string(exportedIdentifier(ctx.Value().GetText()))
	}
	ts := l.values[ctx.TypeSpecifier()].(typeSpec)
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())

	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec: typeSpec{
			Name: fmt.Sprintf("[%v]%v", size, ts.Name),
			Kind: kindArray,
			Size: size,
		},
		InnerDecl: &declaration{
			Identifier: "value",
			TypeSpec:   ts,
		},
	}
}

// 'opaque' IDENTIFIER '[' value ']'       # declOpaqueFixed
func (l *myListener) ExitDeclOpaqueFixed(ctx *parser.DeclOpaqueFixedContext) {
	var size string
	if ctx.Value() != nil {
		size = string(exportedIdentifier(ctx.Value().GetText()))
	}
	ts := typeSpec{Name: "byte"}
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())

	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec: typeSpec{
			Name: fmt.Sprintf("[%v]%v", size, ts.Name),
			Kind: kindByteArray,
			Size: size,
		},
	}
}

// typeSpecifier IDENTIFIER '<' value? '>' # declVar
func (l *myListener) ExitDeclVar(ctx *parser.DeclVarContext) {
	var size string
	if ctx.Value() != nil {
		size = string(exportedIdentifier(ctx.Value().GetText()))
	} else {
		size = "MaxArrayLength"
	}
	ts := l.values[ctx.TypeSpecifier()].(typeSpec)
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())

	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec: typeSpec{
			Name: fmt.Sprintf("[]%v", ts.Name),
			Kind: kindSlice,
			Size: size,
		},
		InnerDecl: &declaration{
			Identifier: "value",
			TypeSpec:   ts,
		},
		Comment: fmt.Sprintf(" // Max length: %v", size),
	}
}

// 'opaque' IDENTIFIER '<' value? '>'      # declOpaqueVar
func (l *myListener) ExitDeclOpaqueVar(ctx *parser.DeclOpaqueVarContext) {
	var size string
	if ctx.Value() != nil {
		size = string(exportedIdentifier(ctx.Value().GetText()))
	} else {
		size = "MaxArrayLength"
	}
	ts := typeSpec{Name: "byte"}
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())

	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec: typeSpec{
			Name: fmt.Sprintf("[]%v", ts.Name),
			Kind: kindByteSlice,
			Size: size,
		},
		Comment: fmt.Sprintf(" // Max length: %v", size),
	}
}

// 'string' IDENTIFIER '<' value? '>'      # declString
func (l *myListener) ExitDeclString(ctx *parser.DeclStringContext) {
	var size string
	if ctx.Value() != nil {
		size = string(exportedIdentifier(ctx.Value().GetText()))
	} else {
		size = "MaxArrayLength"
	}
	ts := typeSpec{Name: "string"}
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())

	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec: typeSpec{
			Name: ts.Name,
			Kind: kindString,
			Size: size,
		},
		Comment: fmt.Sprintf(" // Max length: %v", size),
	}
}

// typeSpecifier '*' IDENTIFIER            # declOptional
func (l *myListener) ExitDeclOptional(ctx *parser.DeclOptionalContext) {
	ts := l.values[ctx.TypeSpecifier()].(typeSpec)
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())

	l.values[ctx] = declaration{
		Identifier: identifier,
		TypeSpec: typeSpec{
			Name: fmt.Sprintf("*%v", ts.Name),
			Kind: kindOptional,
		},
		InnerDecl: &declaration{
			TypeSpec: ts,
		},
	}
}

// 'void'                                  # declVoid
func (l *myListener) ExitDeclVoid(ctx *parser.DeclVoidContext) {
	l.values[ctx] = declaration{}
}
