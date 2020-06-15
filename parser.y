%{
package expr
import "expr/types"
import "strconv"
import "fmt"
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
%token CAST
%left OR
%left AND
%left NOT
%left GT LT GTE LTE EQ NEQ
%left LIKE CONTAINS
%left ADD MINUS
%left MUL DIV
%right LP
%right CAST
%right DOLLAR
%%

input:    e  { yylex.(*Lexer).parseResult=$1.node;};

e:    INT              { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Int); }
    | STR              { $$.node =newAst(CONST,str,types.Text); }
    | RAWSTR           { $$.node =newAst(CONST,unquoteRawString(yylex.(*Lexer).Text()),types.Text); }
    | FLOAT            { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Float); }
    | BOOL             { $$.node =newAst(CONST,yylex.(*Lexer).Text(),types.Bool); }
    | e binary_op e    { $$.node =newAst(FUNC,$2.node.Value,types.Any,$1.node,$3.node).SetOffset($2.node.Offset,$2.node.Length);}
    | NOT e            { $$.node =newAst(FUNC,"not",types.Any,$1.node); }
    | DOLLAR INT       { $$.node =newAst(VAR,yylex.(*Lexer).Text(),types.Any).setOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
    | LP e RP          { $$.node =$2.node ;}
    | IDD LP e_list RP { $$.node =newAst(FUNC,$1.node.Value,types.Any,$3.node.Children...);}
    | IDD LP RP        { $$.node =newAst(FUNC,$1.node.Value,types.Any);}
    | e cast_func      { $$.node =newAst(FUNC,$2.node.Value,types.Any,$1.node);}
;

binary_op: AND         {$$.node=newAst(FUNC,"and",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | OR          {$$.node=newAst(FUNC,"or",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | ADD         {$$.node=newAst(FUNC,"add",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | MINUS       {$$.node=newAst(FUNC,"MINUS",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | DIV         {$$.node=newAst(FUNC,"DIV",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | MUL         {$$.node=newAst(FUNC,"MUL",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | GT          {$$.node=newAst(FUNC,"GT",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | GTE         {$$.node=newAst(FUNC,"GTE",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | LT          {$$.node=newAst(FUNC,"LT",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | LTE         {$$.node=newAst(FUNC,"LTE",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | EQ          {$$.node=newAst(FUNC,"EQ",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
         | NEQ         {$$.node=newAst(FUNC,"NEQ",types.Any).SetOffset(yylex.(*Lexer).Offset,len(yylex.(*Lexer).Text()));}
;
IDD: ID {$$.node=newAst(NULL,yylex.(*Lexer).Text(),types.Any);};

cast_func: CAST IDD    { $$.node =newAst(FUNC,"to"+$2.node.Value,types.Any);}

e_list:   e  {$$.node =newAst(NULL,"",types.Any,$1.node);}
        | e_list COMMA e  {$$.node=newAst(NULL,"",types.Any,append($1.node.Children,$3.node)...);}


