grammar DRConfig;

classDef : STATIC?(classIdentifier) (EXTENDS parentClass)?
           classEnter
           (classDef|property)*
           classLeave
;

classEnter : PARENL ;
classLeave : PARENR ;

classIdentifier : IDENTIFIER | ASTERISK ;

parentClass : IDENTIFIER ;

property : propertyName ASSIGN propertyValue SEMI;

propertyValue : (
                EXCL?IDENTIFIER|
                (IDENTIFIER COLON IDENTIFIER)|
                NUMBER|SINGLESTR|DOUBLESTR|VECTOR3
            ) ;

propertyName : IDENTIFIER ;


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

