package main

import (
	"fmt"

	"github.com/gavrie/oncrpc/parser"
)

type myListener struct {
	*parser.BaseONCRPCv2Listener

	values      map[interface{}]interface{}
	packageName string
}

func newMyListener(packageName string) *myListener {
	return &myListener{
		values:      make(map[interface{}]interface{}),
		packageName: packageName,
	}
}

func (l *myListener) EnterOncrpcv2Specification(ctx *parser.Oncrpcv2SpecificationContext) {
	fmt.Printf(`// package %[1]s is automatically generated from %[1]s.x file by ...

package %[1]s

import (
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/gavrie/oncrpc/types"
)

// Prevent unused imports
var _ = types.ErrInvalidEnumOpt
var _ = xdr.ErrBadArguments

type enum int32 // RFC 1014

const MaxArrayLength = 0x7fffffff

`, l.packageName)
}

// constantDef: 'const' IDENTIFIER '=' constant ';';
func (l *myListener) ExitConstantDef(ctx *parser.ConstantDefContext) {
	identifier := exportedIdentifier(ctx.IDENTIFIER().GetText())
	value := ctx.Constant().GetText()
	switch value {
	case "TRUE":
		value = "true"
	case "FALSE":
		value = "false"
	}
	fmt.Printf("const %v = %s\n", identifier, value)
}
