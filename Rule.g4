grammar Rule;
NUM:[0-9]+;
NOW:'now()';

AND:'and';
OR:'or';
NOT:'not';

boolOperate
    :AND
    |OR
    |NOT
    ;
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
TRUE:'true';
FALSE:'false';

logical
    :AND
    |OR
    |NOT
    ;

compareStatement
    :calculateValue op=compare calculateValue #COMPARE
    |calculateStatement #COMPAREX
    |'(' compareStatement')' #COMPAREX
    ;

calculateValue
    :IDENTIFY #IDEN
    |NUM #NUM
    ;

calculateStatement
    :calculateValue op=calculate calculateValue
    |'(' calculateStatement')'
    ;

boolValue
    :TRUE #BOOL
    |FALSE #BOOL
    |IDENTIFY # IDENBOOL
    |compareStatement #COMPAREVALUE
    ;

boolStatement
    : boolValue op=boolOperate boolValue
    |'(' boolStatement')'
    ;

statement
    :boolStatement
    |compareStatement
    |calculateStatement
    ;
init:statement* EOF?;