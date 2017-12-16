grammar cell;

file            : forms*
                | EOF
                ;
forms           : literal
                | symbolicForm
                ;

literal         : string
                | number
                | character
                | bool
                | symbol
                | list
                ;

symbolicForm    : quote
                | unquote
                ;

quote           : '\'' forms 
                ;

unquote         : '~' forms
                ;

string          : STRING
                ;


number          : FLOAT
                | INT
                ;

character       :
                ;

bool            : TRUE
                | FALSE
                ;

symbol          : SYMBOL
                ;

list            : '(' forms ')'
                ;

STRING          : '"' ( ~'"' | '\\' '"')* '"'
                ;

TRUE            : 'true'
                ;

FALSE           : 'false'
                ;

SYMBOL          : [a-zA-Z]+
                ;

FLOAT           : '-'? DIGIT TAIL
                | '-'? 'inf'
                | EXP
                ;

INT             : DIGIT
                ;

fragment TAIL   : '.' DIGIT EXP?
                ;

fragment DIGIT  : [0-9]+
                ;

fragment EXP    : 'e^' '-'? FLOAT
                ;