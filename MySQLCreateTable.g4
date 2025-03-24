grammar MySQLCreateTable;

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];

// 关键字定义
CREATE: C R E A T E;
TEMPORARY: T E M P O R A R Y;
TABLE: T A B L E;
IF: I F;
NOT: N O T;
EXISTS: E X I S T S;
PRIMARY: P R I M A R Y;
KEY: K E Y;
FOREIGN: F O R E I G N;
REFERENCES: R E F E R E N C E S;
INDEX: I N D E X;
UNIQUE: U N I Q U E;
CONSTRAINT: C O N S T R A I N T;
CHECK: C H E C K;
DEFAULT: D E F A U L T;
AUTO_INCREMENT: A U T O '_' I N C R E M E N T;
ENGINE: E N G I N E;
CHARSET: C H A R S E T;
COLLATE: C O L L A T E;
COMMENT: C O M M E N T;
NULL: N U L L;
CURRENT_TIMESTAMP: C U R R E N T '_' T I M E S T A M P;
ON: O N;
UPDATE: U P D A T E;
CHARACTER: C H A R A C T E R;
SET: S E T;
USING: U S I N G;
JSON: J S O N;
AND: A N D;
OR: O R;

// 数据类型
INT: I N T;
BIGINT: B I G I N T;
TINYINT: T I N Y I N T;
VARCHAR: V A R C H A R;
CHAR: C H A R;
DECIMAL: D E C I M A L;
DATETIME: D A T E T I M E;
DATE: D A T E;
TIMESTAMP: T I M E S T A M P;
TEXT: T E X T;
UNSIGNED: U N S I G N E D;

// 操作符和符号
OPERATOR: '=' | '>' | '<' | '<>' | '!=' | '<=>' | '>=' | '<=' ;
COMMA: ',';
LPAREN: '(';
RPAREN: ')';
SEMI: ';';
DOT: '.';
DELETE: D E L E T E;
RESTRICT: R E S T R I C T;
CASCADE: C A S C A D E;
NO: N O;
ACTION: A C T I O N;

// 标识符
BACKTICK_QUOTED_ID: '`' (~'`' | '``')* '`' -> type(IDENTIFIER);
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

// 字面量
STRING_LITERAL: '\'' (~'\'' | '\'\'')* '\'';
NUMERIC_LITERAL: [0-9]+ ('.' [0-9]*)? ([eE][+-]?[0-9]+)?;

// 空白和注释
WS: [ \t\r\n]+ -> skip;
LINE_COMMENT: '--' ~[\r\n]* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;

create_table_stmt:
    CREATE TEMPORARY? TABLE (IF NOT EXISTS)? table_name
    LPAREN table_element (COMMA table_element)* RPAREN
    table_option* SEMI?
    ;

table_element:
    column_def
    | table_constraint
    ;

column_def:
    column_name data_type 
    (NOT NULL | NULL)?
    (DEFAULT ( literal 
             | CURRENT_TIMESTAMP (ON UPDATE CURRENT_TIMESTAMP)? 
    ))?  // 关键修正：闭合DEFAULT子句的括号
    (AUTO_INCREMENT)?
    (PRIMARY KEY | UNIQUE KEY?)?
    (COMMENT STRING_LITERAL)?
    (CHARACTER SET charset_name)? 
    (COLLATE collation_name)? 
    (REFERENCES table_name LPAREN column_name RPAREN 
      (ON (DELETE | UPDATE) action)* 
    )? 
    (CHECK LPAREN expr RPAREN)?
    ;

data_type:
    INT (LPAREN NUMERIC_LITERAL RPAREN)? (UNSIGNED)?          #IntegerType
    | BIGINT (LPAREN NUMERIC_LITERAL RPAREN)? (UNSIGNED)?     #BigintType
    | TINYINT (LPAREN NUMERIC_LITERAL RPAREN)? (UNSIGNED)?    #TinyintType
    | VARCHAR LPAREN NUMERIC_LITERAL RPAREN                   #VarcharType
    | CHAR LPAREN NUMERIC_LITERAL RPAREN                      #CharType
    | DECIMAL (LPAREN NUMERIC_LITERAL (COMMA NUMERIC_LITERAL)? RPAREN )?  #DecimalType
    | DATETIME                                                #DatetimeType
    | DATE                                                    #DateType
    | TIMESTAMP                                               #TimestampType
    | TEXT                                                    #TextType
    | JSON                                                    #JsonType
    ;

table_constraint:
    (CONSTRAINT identifier)? 
    (
        PRIMARY KEY LPAREN column_list RPAREN 
        | FOREIGN KEY LPAREN column_list RPAREN 
            REFERENCES table_name LPAREN column_list RPAREN 
            (ON (DELETE | UPDATE) action)* 
        | CHECK LPAREN expr RPAREN 
        | (INDEX | KEY) identifier? LPAREN column_list RPAREN (USING index_type)? 
        | UNIQUE (INDEX | KEY)? identifier? LPAREN column_list RPAREN (USING index_type)?
    )
    ;

action: RESTRICT | CASCADE | SET NULL | NO ACTION;

table_option:
    ENGINE OPERATOR engine_name 
    | DEFAULT? (CHARSET | CHARACTER SET) OPERATOR charset_name 
    | DEFAULT? COLLATE OPERATOR collation_name 
    | AUTO_INCREMENT OPERATOR NUMERIC_LITERAL 
    | COMMENT OPERATOR STRING_LITERAL
    ;

// 辅助规则
table_name: IDENTIFIER (DOT IDENTIFIER)?;
column_name: IDENTIFIER;
column_list: column_name (COMMA column_name)*;
engine_name: IDENTIFIER;
charset_name: IDENTIFIER;
collation_name: IDENTIFIER;
literal: STRING_LITERAL | NUMERIC_LITERAL | NULL | CURRENT_TIMESTAMP;
identifier: IDENTIFIER;
index_type: IDENTIFIER;

// 表达式规则
expr: logical_expr;

logical_expr:
    logical_expr (AND | OR) comparison_expr  #LogicalOp
    | comparison_expr                        #SingleComparison
    ;

comparison_expr: 
    primary_expr (OPERATOR primary_expr)?    #ComparisonOp
    | LPAREN expr RPAREN                     #ParenExpr
    ;

primary_expr:
    column_name                              #ColumnRef
    | literal                                #LiteralValue
    | function_call                          #FunctionExpr
    ;

function_call: IDENTIFIER LPAREN (expr (COMMA expr)* )? RPAREN;