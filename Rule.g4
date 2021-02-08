grammar Rule;
IDENTIFY:[a-zA-Z_]+;
NUM:[0-9]+;
NOW:'now()';

AND:'and';
OR:'or';
NOT:'not';

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
    :compareStatement op=compare compareStatement
    |IDENTIFY
    |TRUE
    |FALSE
    |'(' compareStatement')'
    ;
calculateStatement
    :calculateStatement op = calculate calculateStatement
    |IDENTIFY
    |NUM
    |'(' calculateStatement')'
    ;

statement
    :calculateStatement
    |compareStatement
    ;

init:statement* EOF;