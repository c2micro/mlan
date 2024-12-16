grammar Mlan;

WS : [ \t\r\n] -> channel(HIDDEN) ;

program
    : block EOF
    ;

block
    : (statement | functionDefinition | includeSubmodule)*
    ;

statement
    : assignment ';'
    | closureInvoke ';'
    | functionInvoke ';'
    | whileStatement
    | forStatement
    | ifStatement
    | breakStatement ';'
    | continueStatement ';'
    | returnStatement ';'
    ;

whileStatement
    : 'while' expression '{' statement* '}'
    ;

forStatement
    : 'for' assignment ';' expression ';' assignment '{' statement* '}'
    ;

returnStatement
    : 'return' expression?
    ;

continueStatement
    : 'continue'
    ;

breakStatement
    : 'break'
    ;

assignment
    : varScalarName = Identifier '=' expression #assignmentRegular
    | varScalarName = Identifier '=' closureDefinition #assignmentClosure
    | varScalarName = Identifier AssignSum expression #assignmentSum
    | varScalarName = Identifier '-=' expression #assignmentSub
    | varScalarName = Identifier '*=' expression #assignmentMul
    | varScalarName = Identifier '/=' expression #assignmentDiv
    | varScalarName = Identifier '%=' expression #assignmentMod
    | varScalarName = Identifier '**=' expression #assignmentPow
    | varScalarName = Identifier index '=' expression #assignmentIndexRegular
    ;

list
    : '[' (expression (',' expression)*)? ']'
    ;

dictUnit
    : expression ':' expression
    ;

dict
    : '{' (dictUnit (',' dictUnit)*)? '}'
    ;

index
    : '[' expression ']'
    ;

functionInvoke
    : varFunctionName = Identifier '(' (expression (',' expression)*)? ')' #identifierFunctionInvoke
    ;

closureInvoke
    : Closure varClosureName = Identifier '(' (expression (',' expression)*)? ')' #identifierClosureInvoke
    ;

expression
    : Integer #expressionInteger
    | IntegerHex #expressionIntegerHex
    | Null #expressionNull
    | Bool #expressionBool
    | Identifier #expressionIdentifier
    | Float #expressionFloat
    | String #expressionString
    | closureDefinition #expressionClosure
    | functionInvoke #expressionFunctionInvoke
    | closureInvoke #expressionClosureInvoke
    | expression index #expressionIndex
    | Subtract expression #expressionUnaryNegation
    | Not expression #expressionLogicalNot
    | <assoc=right> expression Pow expression #expressionPow
    | expression op=(Multiply|Division|Modulus) expression #expressionMulDivMod
    | expression op=(Add|Subtract) expression #expressionSumSub
    | expression op=(GtEq|LtEq|Gt|Lt) expression #expressionComparison
    | expression op=(Eq|Neq) expression #expressionEqual
    | expression Xor expression #expressionXor
    | expression And expression #expressionLogicalAnd
    | expression Or expression #expressionLogicalOr
    | '(' expression ')' #expressionParentheses
    | list #expressionList
    | dict #expressionDict
    ;

ifBlock
    : 'if' expression '{' statement* '}' #ifBlockStatement
    ;

elifBlock
    : 'elif' expression '{' statement* '}' #elifBlockStatement
    ;

elseBlock
    : 'else' '{' statement* '}' #elseBlockStatement
    ;

ifStatement
    : ifBlock elifBlock* elseBlock?
    ;

functionParameters
    : Identifier (',' Identifier)*
    ;

functionDefinition
    : 'fn' varFunctionName = Identifier '(' functionParameters? ')' '{' statement* '}'
    ;

closureDefinition
    : 'fn' '(' functionParameters? ')' '{' statement* '}'
    ;

includeSubmodule
    : 'include' '(' expression ')' ';'
    ;

Eq : '==' ;
Neq : '!=' ;
Or : '||' ;
And : '&&' ;
Pow : '**' ;
GtEq : '>=' ;
LtEq : '<= ';
AssignSum : '+=' ;
Gt : '>' ;
Lt : '<' ;
Multiply : '*' ;
Division : '/' ;
Modulus : '%' ;
Add : '+' ;
Subtract : '-' ;
Xor : '^' ;
Not : '!' ;
Closure : '@' ;

Bool
    : 'true'
    | 'false'
    ;

Null
    : 'null'
    ;

Identifier
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

Integer
    : [0-9]+
    ;

IntegerHex
    : [0][xX][0-9a-fA-F]+
    ;

Float
    : [0-9]* '.' [0-9]+
    ;

String
    : ["] ( ~["\r\n\\] | '\\' ~[\r\n] )* ["]
    ;

Comment
    : ( '//' ~[\r\n]* | '/*' .*? '*/' ) -> channel(HIDDEN)
    ;
