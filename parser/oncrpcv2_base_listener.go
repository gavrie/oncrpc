// Generated from ONCRPCv2.g4 by ANTLR 4.5.3.

package parser // ONCRPCv2

import "github.com/pboyer/antlr4/runtime/Go/antlr"

// BaseONCRPCv2Listener is a complete listener for a parse tree produced by ONCRPCv2Parser.
type BaseONCRPCv2Listener struct{}

var _ ONCRPCv2Listener = &BaseONCRPCv2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseONCRPCv2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseONCRPCv2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseONCRPCv2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseONCRPCv2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgramDef is called when production programDef is entered.
func (s *BaseONCRPCv2Listener) EnterProgramDef(ctx *ProgramDefContext) {}

// ExitProgramDef is called when production programDef is exited.
func (s *BaseONCRPCv2Listener) ExitProgramDef(ctx *ProgramDefContext) {}

// EnterVersionDef is called when production versionDef is entered.
func (s *BaseONCRPCv2Listener) EnterVersionDef(ctx *VersionDefContext) {}

// ExitVersionDef is called when production versionDef is exited.
func (s *BaseONCRPCv2Listener) ExitVersionDef(ctx *VersionDefContext) {}

// EnterProcedureDef is called when production procedureDef is entered.
func (s *BaseONCRPCv2Listener) EnterProcedureDef(ctx *ProcedureDefContext) {}

// ExitProcedureDef is called when production procedureDef is exited.
func (s *BaseONCRPCv2Listener) ExitProcedureDef(ctx *ProcedureDefContext) {}

// EnterProcReturn is called when production procReturn is entered.
func (s *BaseONCRPCv2Listener) EnterProcReturn(ctx *ProcReturnContext) {}

// ExitProcReturn is called when production procReturn is exited.
func (s *BaseONCRPCv2Listener) ExitProcReturn(ctx *ProcReturnContext) {}

// EnterProcFirstArg is called when production procFirstArg is entered.
func (s *BaseONCRPCv2Listener) EnterProcFirstArg(ctx *ProcFirstArgContext) {}

// ExitProcFirstArg is called when production procFirstArg is exited.
func (s *BaseONCRPCv2Listener) ExitProcFirstArg(ctx *ProcFirstArgContext) {}

// EnterOncrpcv2Specification is called when production oncrpcv2Specification is entered.
func (s *BaseONCRPCv2Listener) EnterOncrpcv2Specification(ctx *Oncrpcv2SpecificationContext) {}

// ExitOncrpcv2Specification is called when production oncrpcv2Specification is exited.
func (s *BaseONCRPCv2Listener) ExitOncrpcv2Specification(ctx *Oncrpcv2SpecificationContext) {}

// EnterDeclSimple is called when production declSimple is entered.
func (s *BaseONCRPCv2Listener) EnterDeclSimple(ctx *DeclSimpleContext) {}

// ExitDeclSimple is called when production declSimple is exited.
func (s *BaseONCRPCv2Listener) ExitDeclSimple(ctx *DeclSimpleContext) {}

// EnterDeclFixed is called when production declFixed is entered.
func (s *BaseONCRPCv2Listener) EnterDeclFixed(ctx *DeclFixedContext) {}

// ExitDeclFixed is called when production declFixed is exited.
func (s *BaseONCRPCv2Listener) ExitDeclFixed(ctx *DeclFixedContext) {}

// EnterDeclVar is called when production declVar is entered.
func (s *BaseONCRPCv2Listener) EnterDeclVar(ctx *DeclVarContext) {}

// ExitDeclVar is called when production declVar is exited.
func (s *BaseONCRPCv2Listener) ExitDeclVar(ctx *DeclVarContext) {}

// EnterDeclOpaqueFixed is called when production declOpaqueFixed is entered.
func (s *BaseONCRPCv2Listener) EnterDeclOpaqueFixed(ctx *DeclOpaqueFixedContext) {}

// ExitDeclOpaqueFixed is called when production declOpaqueFixed is exited.
func (s *BaseONCRPCv2Listener) ExitDeclOpaqueFixed(ctx *DeclOpaqueFixedContext) {}

// EnterDeclOpaqueVar is called when production declOpaqueVar is entered.
func (s *BaseONCRPCv2Listener) EnterDeclOpaqueVar(ctx *DeclOpaqueVarContext) {}

// ExitDeclOpaqueVar is called when production declOpaqueVar is exited.
func (s *BaseONCRPCv2Listener) ExitDeclOpaqueVar(ctx *DeclOpaqueVarContext) {}

// EnterDeclString is called when production declString is entered.
func (s *BaseONCRPCv2Listener) EnterDeclString(ctx *DeclStringContext) {}

// ExitDeclString is called when production declString is exited.
func (s *BaseONCRPCv2Listener) ExitDeclString(ctx *DeclStringContext) {}

// EnterDeclOptional is called when production declOptional is entered.
func (s *BaseONCRPCv2Listener) EnterDeclOptional(ctx *DeclOptionalContext) {}

// ExitDeclOptional is called when production declOptional is exited.
func (s *BaseONCRPCv2Listener) ExitDeclOptional(ctx *DeclOptionalContext) {}

// EnterDeclVoid is called when production declVoid is entered.
func (s *BaseONCRPCv2Listener) EnterDeclVoid(ctx *DeclVoidContext) {}

// ExitDeclVoid is called when production declVoid is exited.
func (s *BaseONCRPCv2Listener) ExitDeclVoid(ctx *DeclVoidContext) {}

// EnterValue is called when production value is entered.
func (s *BaseONCRPCv2Listener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseONCRPCv2Listener) ExitValue(ctx *ValueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseONCRPCv2Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseONCRPCv2Listener) ExitConstant(ctx *ConstantContext) {}

// EnterTypeSpecInt is called when production typeSpecInt is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecInt(ctx *TypeSpecIntContext) {}

// ExitTypeSpecInt is called when production typeSpecInt is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecInt(ctx *TypeSpecIntContext) {}

// EnterTypeSpecHyper is called when production typeSpecHyper is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecHyper(ctx *TypeSpecHyperContext) {}

// ExitTypeSpecHyper is called when production typeSpecHyper is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecHyper(ctx *TypeSpecHyperContext) {}

// EnterTypeSpecFloat is called when production typeSpecFloat is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecFloat(ctx *TypeSpecFloatContext) {}

// ExitTypeSpecFloat is called when production typeSpecFloat is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecFloat(ctx *TypeSpecFloatContext) {}

// EnterTypeSpecDouble is called when production typeSpecDouble is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecDouble(ctx *TypeSpecDoubleContext) {}

// ExitTypeSpecDouble is called when production typeSpecDouble is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecDouble(ctx *TypeSpecDoubleContext) {}

// EnterTypeSpecQuad is called when production typeSpecQuad is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecQuad(ctx *TypeSpecQuadContext) {}

// ExitTypeSpecQuad is called when production typeSpecQuad is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecQuad(ctx *TypeSpecQuadContext) {}

// EnterTypeSpecBool is called when production typeSpecBool is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecBool(ctx *TypeSpecBoolContext) {}

// ExitTypeSpecBool is called when production typeSpecBool is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecBool(ctx *TypeSpecBoolContext) {}

// EnterTypeSpecStructIdentifier is called when production typeSpecStructIdentifier is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecStructIdentifier(ctx *TypeSpecStructIdentifierContext) {}

// ExitTypeSpecStructIdentifier is called when production typeSpecStructIdentifier is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecStructIdentifier(ctx *TypeSpecStructIdentifierContext) {}

// EnterTypeSpecEnum is called when production typeSpecEnum is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecEnum(ctx *TypeSpecEnumContext) {}

// ExitTypeSpecEnum is called when production typeSpecEnum is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecEnum(ctx *TypeSpecEnumContext) {}

// EnterTypeSpecStruct is called when production typeSpecStruct is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecStruct(ctx *TypeSpecStructContext) {}

// ExitTypeSpecStruct is called when production typeSpecStruct is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecStruct(ctx *TypeSpecStructContext) {}

// EnterTypeSpecUnion is called when production typeSpecUnion is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecUnion(ctx *TypeSpecUnionContext) {}

// ExitTypeSpecUnion is called when production typeSpecUnion is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecUnion(ctx *TypeSpecUnionContext) {}

// EnterTypeSpecIdentifier is called when production typeSpecIdentifier is entered.
func (s *BaseONCRPCv2Listener) EnterTypeSpecIdentifier(ctx *TypeSpecIdentifierContext) {}

// ExitTypeSpecIdentifier is called when production typeSpecIdentifier is exited.
func (s *BaseONCRPCv2Listener) ExitTypeSpecIdentifier(ctx *TypeSpecIdentifierContext) {}

// EnterEnumTypeSpec is called when production enumTypeSpec is entered.
func (s *BaseONCRPCv2Listener) EnterEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// ExitEnumTypeSpec is called when production enumTypeSpec is exited.
func (s *BaseONCRPCv2Listener) ExitEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseONCRPCv2Listener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseONCRPCv2Listener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterStructTypeSpec is called when production structTypeSpec is entered.
func (s *BaseONCRPCv2Listener) EnterStructTypeSpec(ctx *StructTypeSpecContext) {}

// ExitStructTypeSpec is called when production structTypeSpec is exited.
func (s *BaseONCRPCv2Listener) ExitStructTypeSpec(ctx *StructTypeSpecContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BaseONCRPCv2Listener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BaseONCRPCv2Listener) ExitStructBody(ctx *StructBodyContext) {}

// EnterUnionTypeSpec is called when production unionTypeSpec is entered.
func (s *BaseONCRPCv2Listener) EnterUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// ExitUnionTypeSpec is called when production unionTypeSpec is exited.
func (s *BaseONCRPCv2Listener) ExitUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// EnterUnionBody is called when production unionBody is entered.
func (s *BaseONCRPCv2Listener) EnterUnionBody(ctx *UnionBodyContext) {}

// ExitUnionBody is called when production unionBody is exited.
func (s *BaseONCRPCv2Listener) ExitUnionBody(ctx *UnionBodyContext) {}

// EnterCaseSpec is called when production caseSpec is entered.
func (s *BaseONCRPCv2Listener) EnterCaseSpec(ctx *CaseSpecContext) {}

// ExitCaseSpec is called when production caseSpec is exited.
func (s *BaseONCRPCv2Listener) ExitCaseSpec(ctx *CaseSpecContext) {}

// EnterDefaultCaseSpec is called when production defaultCaseSpec is entered.
func (s *BaseONCRPCv2Listener) EnterDefaultCaseSpec(ctx *DefaultCaseSpecContext) {}

// ExitDefaultCaseSpec is called when production defaultCaseSpec is exited.
func (s *BaseONCRPCv2Listener) ExitDefaultCaseSpec(ctx *DefaultCaseSpecContext) {}

// EnterConstantDef is called when production constantDef is entered.
func (s *BaseONCRPCv2Listener) EnterConstantDef(ctx *ConstantDefContext) {}

// ExitConstantDef is called when production constantDef is exited.
func (s *BaseONCRPCv2Listener) ExitConstantDef(ctx *ConstantDefContext) {}

// EnterTypeDefTypedef is called when production typeDefTypedef is entered.
func (s *BaseONCRPCv2Listener) EnterTypeDefTypedef(ctx *TypeDefTypedefContext) {}

// ExitTypeDefTypedef is called when production typeDefTypedef is exited.
func (s *BaseONCRPCv2Listener) ExitTypeDefTypedef(ctx *TypeDefTypedefContext) {}

// EnterTypeDefEnum is called when production typeDefEnum is entered.
func (s *BaseONCRPCv2Listener) EnterTypeDefEnum(ctx *TypeDefEnumContext) {}

// ExitTypeDefEnum is called when production typeDefEnum is exited.
func (s *BaseONCRPCv2Listener) ExitTypeDefEnum(ctx *TypeDefEnumContext) {}

// EnterTypeDefStruct is called when production typeDefStruct is entered.
func (s *BaseONCRPCv2Listener) EnterTypeDefStruct(ctx *TypeDefStructContext) {}

// ExitTypeDefStruct is called when production typeDefStruct is exited.
func (s *BaseONCRPCv2Listener) ExitTypeDefStruct(ctx *TypeDefStructContext) {}

// EnterTypeDefUnion is called when production typeDefUnion is entered.
func (s *BaseONCRPCv2Listener) EnterTypeDefUnion(ctx *TypeDefUnionContext) {}

// ExitTypeDefUnion is called when production typeDefUnion is exited.
func (s *BaseONCRPCv2Listener) ExitTypeDefUnion(ctx *TypeDefUnionContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseONCRPCv2Listener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseONCRPCv2Listener) ExitDefinition(ctx *DefinitionContext) {}

// EnterXdrSpecification is called when production xdrSpecification is entered.
func (s *BaseONCRPCv2Listener) EnterXdrSpecification(ctx *XdrSpecificationContext) {}

// ExitXdrSpecification is called when production xdrSpecification is exited.
func (s *BaseONCRPCv2Listener) ExitXdrSpecification(ctx *XdrSpecificationContext) {}
