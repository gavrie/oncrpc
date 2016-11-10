package main

import (
	"strings"

	"github.com/gavrie/oncrpc/parser"
)

// typeSpecifier

//go:generate stringer -type=typeKind

type typeKind int

const (
	kindInvalid typeKind = iota
	kindIdentifier
	kindInt32
	kindUInt32
	kindInt64
	kindUInt64
	kindFloat32
	kindFloat64
	kindBool
	kindArray
	kindSlice
	kindByteArray
	kindByteSlice
	kindString
	kindOptional
)

type typeSpec struct {
	Name string
	Kind typeKind
	Size string
}

// ('unsigned' | ('unsigned'? 'int'))   # typeSpecInt
func (l *myListener) ExitTypeSpecInt(ctx *parser.TypeSpecIntContext) {
	var val typeSpec
	if strings.HasPrefix(ctx.GetText(), "unsigned") {
		val.Kind = kindUInt32
		val.Name = "uint32"
	} else {
		val.Kind = kindInt32
		val.Name = "int32"
	}
	l.values[ctx] = val
}

// 'unsigned'? 'hyper' # typeSpecHyper
func (l *myListener) EnterTypeSpecHyper(ctx *parser.TypeSpecHyperContext) {
	var val typeSpec
	if strings.HasPrefix(ctx.GetText(), "unsigned") {
		val.Kind = kindUInt64
		val.Name = "uint64"
	} else {
		val.Kind = kindInt64
		val.Name = "int64"
	}
	l.values[ctx] = val
}

// 'float'             # typeSpecFloat
func (l *myListener) EnterTypeSpecFloat(ctx *parser.TypeSpecFloatContext) {
	l.values[ctx] = typeSpec{
		Kind: kindFloat32,
		Name: "float32",
	}
}

// 'double'            # typeSpecDouble
func (l *myListener) EnterTypeSpecDouble(ctx *parser.TypeSpecDoubleContext) {
	l.values[ctx] = typeSpec{
		Kind: kindFloat64,
		Name: "float64",
	}
}

// 'quadruple'         # typeSpecQuad
func (l *myListener) EnterTypeSpecQuad(ctx *parser.TypeSpecQuadContext) {
	panic("unsupported type 'quadruple'")
}

// 'bool'              # typeSpecBool
func (l *myListener) EnterTypeSpecBool(ctx *parser.TypeSpecBoolContext) {
	l.values[ctx] = typeSpec{
		Kind: kindBool,
		Name: "bool",
	}
}

// 'struct' IDENTIFIER # typeSpecStructIdentifier
func (l *myListener) EnterTypeSpecStructIdentifier(ctx *parser.TypeSpecStructIdentifierContext) {
	panic("not implemented yet")
	// l.values[ctx] = typeSpec(exportedIdentifier(ctx.IDENTIFIER().GetText()))
}

// enumTypeSpec        # typeSpecEnum
func (l *myListener) EnterTypeSpecEnum(ctx *parser.TypeSpecEnumContext) {
	panic("not implemented yet")
}

// structTypeSpec      # typeSpecStruct
func (l *myListener) EnterTypeSpecStruct(ctx *parser.TypeSpecStructContext) {
	panic("not implemented yet")
}

// unionTypeSpec       # typeSpecUnion
func (l *myListener) EnterTypeSpecUnion(ctx *parser.TypeSpecUnionContext) {
	panic("not implemented yet")
}

// IDENTIFIER          # typeSpecIdentifier
func (l *myListener) EnterTypeSpecIdentifier(ctx *parser.TypeSpecIdentifierContext) {
	l.values[ctx] = typeSpec{
		Name: string(exportedIdentifier(ctx.GetText())),
		Kind: kindIdentifier,
	}
}
