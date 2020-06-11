%{
package expr
import "expr/types"
import "strconv"
%}

%union {
  node *AstNode
}
%token NULL
%token INT
%token STR
%token RAWSTR
%token BOOL
%token FLOAT
%token CONST
%token OR
%token AND
%token NOT
%token LIKE
%token NEQ
%token GT
%token LT
%token GTE
%token LTE
%token EQ
%token ADD
%token MINUS
%token MUL
%token DIV
%token CONTAINS
%token ID
%token IND
%token COMMA
%token ANY
%token FUNC
%token LP
%token RP
%token DOLLAR
%token VAR
%left OR
%left AND
%left NOT
%left GT LT GTE LTE EQ NEQ
%left LIKE CONTAINS
%left ADD MINUS
%left MUL DIV
%right LP
%right DOLLAR
%%

input:    e  { yylex.(*Lexer).parseResult=$1.node;};

e:    INT              { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Int); }
    | STR              { str,_:=strconv.Unquote(yylex.(*Lexer).Text());$$.node =newAst(CONST,str,types.Text); }
    | RAWSTR           { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Text); }
    | FLOAT            { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Float); }
    | BOOL             { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Bool); }
    | e AND e          { $$.node =newAst(FUNC,"and",types.Any,$1.node,$3.node); }
    | e OR e           { $$.node =newAst(FUNC,"or",types.Any,$1.node,$3.node); }
    | e ADD e          { $$.node =newAst(FUNC,"add",types.Any,$1.node,$3.node); }
    | e MINUS e        { $$.node =newAst(FUNC,"minus",types.Any,$1.node,$3.node); }
    | e DIV e          { $$.node =newAst(FUNC,"div",types.Any,$1.node,$3.node); }
    | e MUL e          { $$.node =newAst(FUNC,"mul",types.Any,$1.node,$3.node); }
    | NOT e            { $$.node =newAst(FUNC,"not",types.Any,$1.node); }
    | e GT e           { $$.node =newAst(FUNC,"gt",types.Any,$1.node,$3.node); }
    | e GTE e          { $$.node =newAst(FUNC,"gte",types.Any,$1.node,$3.node); }
    | e LT e           { $$.node =newAst(FUNC,"lt",types.Any,$1.node,$3.node); }
    | e LTE e          { $$.node =newAst(FUNC,"lte",types.Any,$1.node,$3.node); }
    | e EQ e           { $$.node =newAst(FUNC,"eq",types.Any,$1.node,$3.node); }
    | e NEQ e          { $$.node =newAst(FUNC,"neq",types.Any,$1.node,$3.node); }
    | DOLLAR INT       { $$.node =newAst(VAR,yylex.(*Lexer).Text(),types.Any);}
    | LP e RP          { $$.node =$2.node ;}
    | IDD LP e_list RP { $$.node =newAst(FUNC,$1.node.Value,types.Any,$3.node.Children...);}
    | IDD LP RP        { $$.node =newAst(FUNC,$1.node.Value,types.Any);}
;

IDD: ID {$$.node=newAst(NULL,yylex.(*Lexer).Text(),types.Any);};

e_list:   e  {$$.node =newAst(NULL,"",types.Any,$1.node);}
        | e_list COMMA e  {$$.node=newAst(NULL,"",types.Any,append($1.node.Children,$3.node)...);}


