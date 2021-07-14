// Define a grammar called Hello
grammar DRConfig;

classDef : STATIC?(classIdentifier) (EXTENDS parentClass)?
           PARENL
           (classDef|property)*
           PARENR
;

classIdentifier : IDENTIFIER | ASTERISK ;

parentClass : IDENTIFIER ;

property : IDENTIFIER ASSIGN (
                EXCL?IDENTIFIER|
                (IDENTIFIER COLON IDENTIFIER)|
                NUMBER|SINGLESTR|DOUBLESTR|VECTOR3
            )SEMI;

COMMENT : '//' ~[\r\n]* -> skip;
MLCOMMENT : '/*' .*? '*/' -> skip;
VECTOR3 : NUMBER COMMA NUMBER COMMA NUMBER;

EXTENDS : 'extends';
STATIC : 'static';
ASSIGN : '=';
PARENL : '{';
PARENR : '}';
EXCL : '!';
SEMI: ';';
DOT : '.';
COMMA : ',';
COLON : ':';
ASTERISK : '*';
SLASHSLASH : '//';

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
EOL: '\n';

SINGLESTR : ['].*?['];
DOUBLESTR : ["].*?["];
IDENTIFIER : ([a-zA-Z0-9_-]|DOT)+;
NUMBER : [-]?[0-9.]+;
ANY: .;

//r  : 'hello' ID ;         // match keyword hello followed by an identifier
//ID : [a-z]+ ;             // match lower-case identifiers

