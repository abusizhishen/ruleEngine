grammar Rule;
NUM:[0-9]+ ;
NOW:'now()';

AND:'and';
OR:'or';
NOT:'not';

IF:'if';
THEN:'then';
ELIF:'elif';
ELSE:'else';
END:'end';

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

num:NUM#NUM;
identify:IDENTIFY#IDENTIFY;

calculateValue
    :identify
    |num
    ;

calculateStatement
    :calculateStatement (op=(MUL|DIV) calculateStatement)+ #MULDIV
    |calculateStatement (op=(ADD|SUB) calculateStatement)+ #ADDSUB
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
    :identify
    |num
    |calculateStatement
    ;
setValueStatement
    : IDENTIFY '=' valueType ';'?
    ;
ifStatement:
    IF boolStatement ':'
        setValueStatement*
    (ELIF boolStatement ':'
        setValueStatement*) ?
    (ELSE ':'
        setValueStatement*) ?
    END
    ;

statement
    :boolStatement
    |compareStatement
    |calculateStatement
    |ifStatement
    |setValueStatement
    ;
init:statement* EOF?;