grammar XDR;

// parser rules

declaration:
    typeSpecifier IDENTIFIER                # declSimple
  | typeSpecifier IDENTIFIER '[' value ']'  # declFixed
  | typeSpecifier IDENTIFIER '<' value? '>' # declVar
  | 'opaque' IDENTIFIER '[' value ']'       # declOpaqueFixed
  | 'opaque' IDENTIFIER '<' value? '>'      # declOpaqueVar
  | 'string' IDENTIFIER '<' value? '>'      # declString
  | typeSpecifier '*' IDENTIFIER            # declOptional
  | 'void'                                  # declVoid
;
value: constant | IDENTIFIER;
constant: DECIMAL | HEXADECIMAL | OCTAL | BOOLEAN;
typeSpecifier:
    ('unsigned' | ('unsigned'? 'int'))   # typeSpecInt
  | 'unsigned'? 'hyper' # typeSpecHyper
  | 'float'             # typeSpecFloat
  | 'double'            # typeSpecDouble
  | 'quadruple'         # typeSpecQuad
  | 'bool'              # typeSpecBool
  | 'struct' IDENTIFIER # typeSpecStructIdentifier
  | enumTypeSpec        # typeSpecEnum
  | structTypeSpec      # typeSpecStruct
  | unionTypeSpec       # typeSpecUnion
  | IDENTIFIER          # typeSpecIdentifier
;
enumTypeSpec: 'enum' enumBody;
enumBody: '{' (IDENTIFIER '=' value) (',' IDENTIFIER '=' value)* '}';
structTypeSpec: 'struct' structBody;
structBody: '{' (declaration ';') (declaration ';')* '}';
unionTypeSpec: 'union' unionBody;
unionBody: 'switch' '(' declaration ')' '{'
        caseSpec
        caseSpec*
        defaultCaseSpec?
    '}';
caseSpec: ('case' value ':') ('case' value ':')* declaration ';';
defaultCaseSpec: ('default' ':' declaration ';');
constantDef: 'const' IDENTIFIER '=' constant ';';
typeDef:
    'typedef' declaration ';'          # typeDefTypedef
  | 'enum' IDENTIFIER enumBody ';'     # typeDefEnum
  | 'struct' IDENTIFIER structBody ';' # typeDefStruct
  | 'union' IDENTIFIER unionBody ';'   # typeDefUnion
;
definition: typeDef | constantDef;
xdrSpecification: definition+; //this is the top level rule for xdr (rfc 4506)

// lexer rules

COMMENT : '/*' .*? '*/' -> skip;
OCTAL : '0' [1-7] ([0-7])*;
DECIMAL : ('-')? ([0-9])+;
HEXADECIMAL : '0x' ([a-fA-F0-9])+;
BOOLEAN : 'TRUE' | 'FALSE';
IDENTIFIER : [a-zA-Z_] ([a-zA-Z0-9_])*;
WS : [ \t\r\n]+ -> skip;
