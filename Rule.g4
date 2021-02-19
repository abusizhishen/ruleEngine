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
    :compareStatement op=compare compareStatement #COMPARE
    |calculateValue #ITEMCOMP
    |calculateStatement #Calcu
    |'(' compareStatement')' #COMPAREX
    ;

calculateValue
    :IDENTIFY #IDEN
    |NUM #NUM
    ;

calculateStatement
    :calculateStatement (op=(MUL|DIV) calculateStatement)+ #MULDIV
    |calculateStatement (op=(ADD|SUB) calculateStatement)+ #ADDSUB
    |calculateValue #ITEMCALCU
    //|calculateValue (op=calculate calculateValue)+ #d
    |'(' calculateStatement')' #CALCULATEX
    ;

boolStatement
    : boolStatement op=boolOperate boolStatement #BOOLOP
    |('true'|'false') #BOOL
    |IDENTIFY #IDENBOOL
    |compareStatement #COMPAREBOOL
    |'(' boolStatement')' #BOOLOPX
    ;

statement
    :boolStatement
    |compareStatement
    |calculateStatement
    ;
init:statement* EOF?;