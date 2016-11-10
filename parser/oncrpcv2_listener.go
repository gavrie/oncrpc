// Generated from ONCRPCv2.g4 by ANTLR 4.5.3.

package parser // ONCRPCv2

import "github.com/pboyer/antlr4/runtime/Go/antlr"

// ONCRPCv2Listener is a complete listener for a parse tree produced by ONCRPCv2Parser.
type ONCRPCv2Listener interface {
	antlr.ParseTreeListener

	// EnterProgramDef is called when entering the programDef production.
	EnterProgramDef(c *ProgramDefContext)

	// EnterVersionDef is called when entering the versionDef production.
	EnterVersionDef(c *VersionDefContext)

	// EnterProcedureDef is called when entering the procedureDef production.
	EnterProcedureDef(c *ProcedureDefContext)

	// EnterProcReturn is called when entering the procReturn production.
	EnterProcReturn(c *ProcReturnContext)

	// EnterProcFirstArg is called when entering the procFirstArg production.
	EnterProcFirstArg(c *ProcFirstArgContext)

	// EnterOncrpcv2Specification is called when entering the oncrpcv2Specification production.
	EnterOncrpcv2Specification(c *Oncrpcv2SpecificationContext)

	// EnterDeclSimple is called when entering the declSimple production.
	EnterDeclSimple(c *DeclSimpleContext)

	// EnterDeclFixed is called when entering the declFixed production.
	EnterDeclFixed(c *DeclFixedContext)

	// EnterDeclVar is called when entering the declVar production.
	EnterDeclVar(c *DeclVarContext)

	// EnterDeclOpaqueFixed is called when entering the declOpaqueFixed production.
	EnterDeclOpaqueFixed(c *DeclOpaqueFixedContext)

	// EnterDeclOpaqueVar is called when entering the declOpaqueVar production.
	EnterDeclOpaqueVar(c *DeclOpaqueVarContext)

	// EnterDeclString is called when entering the declString production.
	EnterDeclString(c *DeclStringContext)

	// EnterDeclOptional is called when entering the declOptional production.
	EnterDeclOptional(c *DeclOptionalContext)

	// EnterDeclVoid is called when entering the declVoid production.
	EnterDeclVoid(c *DeclVoidContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterTypeSpecInt is called when entering the typeSpecInt production.
	EnterTypeSpecInt(c *TypeSpecIntContext)

	// EnterTypeSpecHyper is called when entering the typeSpecHyper production.
	EnterTypeSpecHyper(c *TypeSpecHyperContext)

	// EnterTypeSpecFloat is called when entering the typeSpecFloat production.
	EnterTypeSpecFloat(c *TypeSpecFloatContext)

	// EnterTypeSpecDouble is called when entering the typeSpecDouble production.
	EnterTypeSpecDouble(c *TypeSpecDoubleContext)

	// EnterTypeSpecQuad is called when entering the typeSpecQuad production.
	EnterTypeSpecQuad(c *TypeSpecQuadContext)

	// EnterTypeSpecBool is called when entering the typeSpecBool production.
	EnterTypeSpecBool(c *TypeSpecBoolContext)

	// EnterTypeSpecStructIdentifier is called when entering the typeSpecStructIdentifier production.
	EnterTypeSpecStructIdentifier(c *TypeSpecStructIdentifierContext)

	// EnterTypeSpecEnum is called when entering the typeSpecEnum production.
	EnterTypeSpecEnum(c *TypeSpecEnumContext)

	// EnterTypeSpecStruct is called when entering the typeSpecStruct production.
	EnterTypeSpecStruct(c *TypeSpecStructContext)

	// EnterTypeSpecUnion is called when entering the typeSpecUnion production.
	EnterTypeSpecUnion(c *TypeSpecUnionContext)

	// EnterTypeSpecIdentifier is called when entering the typeSpecIdentifier production.
	EnterTypeSpecIdentifier(c *TypeSpecIdentifierContext)

	// EnterEnumTypeSpec is called when entering the enumTypeSpec production.
	EnterEnumTypeSpec(c *EnumTypeSpecContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterStructTypeSpec is called when entering the structTypeSpec production.
	EnterStructTypeSpec(c *StructTypeSpecContext)

	// EnterStructBody is called when entering the structBody production.
	EnterStructBody(c *StructBodyContext)

	// EnterUnionTypeSpec is called when entering the unionTypeSpec production.
	EnterUnionTypeSpec(c *UnionTypeSpecContext)

	// EnterUnionBody is called when entering the unionBody production.
	EnterUnionBody(c *UnionBodyContext)

	// EnterCaseSpec is called when entering the caseSpec production.
	EnterCaseSpec(c *CaseSpecContext)

	// EnterDefaultCaseSpec is called when entering the defaultCaseSpec production.
	EnterDefaultCaseSpec(c *DefaultCaseSpecContext)

	// EnterConstantDef is called when entering the constantDef production.
	EnterConstantDef(c *ConstantDefContext)

	// EnterTypeDefTypedef is called when entering the typeDefTypedef production.
	EnterTypeDefTypedef(c *TypeDefTypedefContext)

	// EnterTypeDefEnum is called when entering the typeDefEnum production.
	EnterTypeDefEnum(c *TypeDefEnumContext)

	// EnterTypeDefStruct is called when entering the typeDefStruct production.
	EnterTypeDefStruct(c *TypeDefStructContext)

	// EnterTypeDefUnion is called when entering the typeDefUnion production.
	EnterTypeDefUnion(c *TypeDefUnionContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterXdrSpecification is called when entering the xdrSpecification production.
	EnterXdrSpecification(c *XdrSpecificationContext)

	// ExitProgramDef is called when exiting the programDef production.
	ExitProgramDef(c *ProgramDefContext)

	// ExitVersionDef is called when exiting the versionDef production.
	ExitVersionDef(c *VersionDefContext)

	// ExitProcedureDef is called when exiting the procedureDef production.
	ExitProcedureDef(c *ProcedureDefContext)

	// ExitProcReturn is called when exiting the procReturn production.
	ExitProcReturn(c *ProcReturnContext)

	// ExitProcFirstArg is called when exiting the procFirstArg production.
	ExitProcFirstArg(c *ProcFirstArgContext)

	// ExitOncrpcv2Specification is called when exiting the oncrpcv2Specification production.
	ExitOncrpcv2Specification(c *Oncrpcv2SpecificationContext)

	// ExitDeclSimple is called when exiting the declSimple production.
	ExitDeclSimple(c *DeclSimpleContext)

	// ExitDeclFixed is called when exiting the declFixed production.
	ExitDeclFixed(c *DeclFixedContext)

	// ExitDeclVar is called when exiting the declVar production.
	ExitDeclVar(c *DeclVarContext)

	// ExitDeclOpaqueFixed is called when exiting the declOpaqueFixed production.
	ExitDeclOpaqueFixed(c *DeclOpaqueFixedContext)

	// ExitDeclOpaqueVar is called when exiting the declOpaqueVar production.
	ExitDeclOpaqueVar(c *DeclOpaqueVarContext)

	// ExitDeclString is called when exiting the declString production.
	ExitDeclString(c *DeclStringContext)

	// ExitDeclOptional is called when exiting the declOptional production.
	ExitDeclOptional(c *DeclOptionalContext)

	// ExitDeclVoid is called when exiting the declVoid production.
	ExitDeclVoid(c *DeclVoidContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitTypeSpecInt is called when exiting the typeSpecInt production.
	ExitTypeSpecInt(c *TypeSpecIntContext)

	// ExitTypeSpecHyper is called when exiting the typeSpecHyper production.
	ExitTypeSpecHyper(c *TypeSpecHyperContext)

	// ExitTypeSpecFloat is called when exiting the typeSpecFloat production.
	ExitTypeSpecFloat(c *TypeSpecFloatContext)

	// ExitTypeSpecDouble is called when exiting the typeSpecDouble production.
	ExitTypeSpecDouble(c *TypeSpecDoubleContext)

	// ExitTypeSpecQuad is called when exiting the typeSpecQuad production.
	ExitTypeSpecQuad(c *TypeSpecQuadContext)

	// ExitTypeSpecBool is called when exiting the typeSpecBool production.
	ExitTypeSpecBool(c *TypeSpecBoolContext)

	// ExitTypeSpecStructIdentifier is called when exiting the typeSpecStructIdentifier production.
	ExitTypeSpecStructIdentifier(c *TypeSpecStructIdentifierContext)

	// ExitTypeSpecEnum is called when exiting the typeSpecEnum production.
	ExitTypeSpecEnum(c *TypeSpecEnumContext)

	// ExitTypeSpecStruct is called when exiting the typeSpecStruct production.
	ExitTypeSpecStruct(c *TypeSpecStructContext)

	// ExitTypeSpecUnion is called when exiting the typeSpecUnion production.
	ExitTypeSpecUnion(c *TypeSpecUnionContext)

	// ExitTypeSpecIdentifier is called when exiting the typeSpecIdentifier production.
	ExitTypeSpecIdentifier(c *TypeSpecIdentifierContext)

	// ExitEnumTypeSpec is called when exiting the enumTypeSpec production.
	ExitEnumTypeSpec(c *EnumTypeSpecContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitStructTypeSpec is called when exiting the structTypeSpec production.
	ExitStructTypeSpec(c *StructTypeSpecContext)

	// ExitStructBody is called when exiting the structBody production.
	ExitStructBody(c *StructBodyContext)

	// ExitUnionTypeSpec is called when exiting the unionTypeSpec production.
	ExitUnionTypeSpec(c *UnionTypeSpecContext)

	// ExitUnionBody is called when exiting the unionBody production.
	ExitUnionBody(c *UnionBodyContext)

	// ExitCaseSpec is called when exiting the caseSpec production.
	ExitCaseSpec(c *CaseSpecContext)

	// ExitDefaultCaseSpec is called when exiting the defaultCaseSpec production.
	ExitDefaultCaseSpec(c *DefaultCaseSpecContext)

	// ExitConstantDef is called when exiting the constantDef production.
	ExitConstantDef(c *ConstantDefContext)

	// ExitTypeDefTypedef is called when exiting the typeDefTypedef production.
	ExitTypeDefTypedef(c *TypeDefTypedefContext)

	// ExitTypeDefEnum is called when exiting the typeDefEnum production.
	ExitTypeDefEnum(c *TypeDefEnumContext)

	// ExitTypeDefStruct is called when exiting the typeDefStruct production.
	ExitTypeDefStruct(c *TypeDefStructContext)

	// ExitTypeDefUnion is called when exiting the typeDefUnion production.
	ExitTypeDefUnion(c *TypeDefUnionContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitXdrSpecification is called when exiting the xdrSpecification production.
	ExitXdrSpecification(c *XdrSpecificationContext)
}
