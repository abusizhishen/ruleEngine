grammar Rule;
NUM:[0-9]+ ;
NOW:'now()';

AND:'&&';
OR:'||';
NOT:'!';

IF:'if';
ELSIF:'elsif';
ELSE:'else';

boolOperate
    :AND
    |OR
    |NOT
    ;

    //antlr

TRUE:'true';
FALSE:'false';

IDENTIFY:[a-zA-Z_]+;

GT:'>';
GTE:'>=';
EQ:'==';
LTE:'<=';
LT:'<';

ADD: '+';
SUB: '-';
MUL:'*';
DIV:'/';

WS:[\r\n\t ]+ -> skip;

calculate
    :ADD
    |SUB
    |MUL
    |DIV
    ;

compare
    :GT
    |GTE
    |EQ
    |LTE
    |LT
    ;

logical
    :AND
    |OR
    |NOT
    ;

compareStatement
    :compareStatement op=compare compareStatement #COMPARE
    |calculateValue #ITEMCOMP
    |calculateStatement #Calcu
    |'(' compareStatement')' #COMPAREX
    ;

key: stringValue|num ;
num:NUM#NUM;
boolValue:TRUE|FALSE;
identify:IDENTIFY#IDENTIFY;
stringValue:Str #STRING;

calculateValue
    :identify
    |num
    ;

calculateStatement
    :calculateStatement op=(MUL|DIV) calculateStatement #MULDIV
    |calculateStatement op=(ADD|SUB) calculateStatement #ADDSUB
    |calculateValue #ITEMCALCU
    |'(' calculateStatement')' #CALCULATEX
    ;

boolStatement
    : boolStatement op=boolOperate boolStatement #BOOLOP
    |('true'|'false') #BOOL
    |IDENTIFY #IDENBOOL
    |compareStatement #COMPAREBOOL
    |'(' boolStatement')' #BOOLOPX
    ;

valueType
    :boolValue
    |identify
    |num
    |calculateStatement
    |mapValue
    |getMapOrArrayValue
    |stringValue
    |array
    ;

setValueStatement
    : IDENTIFY '=' valueType ';'?
    ;

ifStatement:
    IF boolStatement '{'
        statement*
    '}'
    elseIfStatement*
    elseStatement?
    ;

elseIfStatement:
    ELSIF boolStatement '{'
        statement*
'}';

elseStatement:
    ELSE '{'
    statement*
'}';

returnStatement:
    'return'
       value= valueType|calculateStatement|boolStatement
    ;

Str: '"' [a-zA-Z_0-9 ']+ '"' ;
pair: mapKey=Str ':' valueType;
mapValue
    :'{' (pair (',' pair)*)? '}'
    ;
getMapOrArrayValue:identify ('[' IDENTIFY ']')+;

array:'[' (valueType (',' valueType)*)? ']';
statement
    :calculateStatement
    |boolStatement
    |compareStatement
    |ifStatement
    |setValueStatement
    |returnStatement
    ;

init:statement* EOF?;