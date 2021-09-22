%{
package expr
import "github.com/yjhatfdu/expr/types"
%}

%union {
  offset int
  node *AstNode
  text string
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
%token PIPE
%token WHEN
%token CASE
%token THEN
%token IN
%token END
%token ELSE
%token IS
%token SIMILAR
%token TO
%left OR
%left AND
%left NOT SIMILAR
%left GT LT GTE LTE EQ NEQ
%left LIKE CONTAINS TO
%left ADD MINUS
%left MUL DIV
%left PIPE
%right IS
%right IN
%right LP
%right DOLLAR
%right CAST
%%

input:    e      { yylex.(*Lexer).parseResult=$1.node;};

e:    INT              { $$.node =newAst(CONST,$1.text,types.Int,$1.offset); }
    | STR              { $$.node =newAst(CONST,$1.text,types.Text,$1.offset); }
    | RAWSTR           { $$.node =newAst(CONST,$1.text,RAWSTR,$1.offset); }
    | FLOAT            { $$.node =newAst(CONST,$1.text,types.Float,$1.offset); }
    | BOOL             { $$.node =newAst(CONST,$1.text,types.Bool,$1.offset); }
    | e LIKE e          { $$.node =newAst(FUNC,"like",types.Any,$2.offset,$1.node,$3.node); }
    | e SIMILAR TO e          { $$.node =newAst(FUNC,"similarTo",types.Any,$2.offset,$1.node,$4.node); }
    | e NOT SIMILAR TO e          { $$.node =newAst(FUNC,"notSimilarTo",types.Any,$2.offset,$1.node,$5.node); }
    | e NOT LIKE e          { $$.node =newAst(FUNC,"notLike",types.Any,$2.offset,$1.node,$4.node); }
    | e CONTAINS e          { $$.node =newAst(FUNC,"contains",types.Any,$2.offset,$1.node,$3.node); }
    | e AND e          { $$.node =newAst(FUNC,"and",types.Any,$2.offset,$1.node,$3.node); }
    | e OR e           { $$.node =newAst(FUNC,"or",types.Any,$2.offset,$1.node,$3.node); }
    | e ADD e          { $$.node =newAst(FUNC,"add",types.Any,$2.offset,$1.node,$3.node); }
    | e MINUS e        { $$.node =newAst(FUNC,"minus",types.Any,$2.offset,$1.node,$3.node); }
    | e DIV e          { $$.node =newAst(FUNC,"div",types.Any,$2.offset,$1.node,$3.node); }
    | e MUL e          { $$.node =newAst(FUNC,"mul",types.Any,$2.offset,$1.node,$3.node); }
    | NOT e            { $$.node =newAst(FUNC,"not",types.Any,$1.offset,$2.node); }
    | e GT e           { $$.node =newAst(FUNC,"gt",types.Any,$2.offset,$1.node,$3.node); }
    | e GTE e          { $$.node =newAst(FUNC,"gte",types.Any,$2.offset,$1.node,$3.node); }
    | e LT e           { $$.node =newAst(FUNC,"lt",types.Any,$2.offset,$1.node,$3.node); }
    | e LTE e          { $$.node =newAst(FUNC,"lte",types.Any,$2.offset,$1.node,$3.node); }
    | e EQ e           { $$.node =newAst(FUNC,"eq",types.Any,$2.offset,$1.node,$3.node); }
    | e NEQ e          { $$.node =newAst(FUNC,"neq",types.Any,$2.offset,$1.node,$3.node); }
    | e IS NULL        { $$.node =newAst(FUNC,"isNull",types.Any,$2.offset,$1.node); }
    | e IS NOT NULL    { $$.node =newAst(FUNC,"isNotNull",types.Any,$2.offset,$1.node); }
    | DOLLAR INT       { $$.node =newAst(VAR,$2.text,types.Any,$1.offset);}
    | DOLLAR MUL       { $$.node =newAst(VAR,"ALL",types.Any,$1.offset);}
    | LP e RP          { $$.node =$2.node;}
    | func_call        { $$.node =$1.node;}
    | negative 	       { $$.node =$1.node;}
    | NULL             { $$.node =newAst(CONST,"null",types.Null,$1.offset);}
    | CASE whenClause END { $$.node =$2.node;}
    | CASE whenClause ELSE e END {
        $$.node =$2.node
        $$.node.Children=append($$.node.Children,$4.node)
    }
    | e IN LP e_list RP  {
        $$.node=newAst(FUNC,"in",types.Any,$2.offset,append([]*AstNode{$1.node},$4.node.Children...)...)
    }
    | e NOT IN LP e_list RP  {
         $$.node=newAst(FUNC,"notIn",types.Any,$2.offset,append([]*AstNode{$1.node},$5.node.Children...)...)
     }


negative : MINUS INT { $$.node =newAst(CONST,"-" + $2.text,types.Int,$2.offset); }
	| MINUS FLOAT { $$.node =newAst(CONST,"-" + $2.text,types.Float,$2.offset); }

func_call :     IDD LP e_list RP { $$.node =newAst(FUNC,$1.node.Value,types.Any,$1.offset,$3.node.Children...);}
              | IDD LP RP        { $$.node =newAst(FUNC,$1.node.Value,types.Any,$1.offset);}
              | IDD              { $$.node =newAst(FUNC,$1.node.Value,types.Any,$1.offset);}
              | e cast_func      { $$.node =newAst(FUNC,$2.node.Value,types.Any,$2.offset,$1.node);}
              | e PIPE IDD       { $3.node.Children=append([]*AstNode{$1.node},$3.node.Children...);$$.node=$3.node}
              | e PIPE IDD LP RP { $3.node.Children=append([]*AstNode{$1.node},$3.node.Children...);$$.node=$3.node}
              | e PIPE IDD LP e_list RP { $$.node =newAst(FUNC,$3.node.Value,types.Any,$3.offset,append([]*AstNode{$1.node},$5.node.Children...)...);}

whenClause  : WHEN e  THEN e
            {$$.node = newAst(FUNC,"multiIf",types.Any,$1.offset,$2.node,$4.node)}
            | whenClause WHEN e  THEN e
            {
            $1.node.Children=append($1.node.Children,$3.node,$5.node)
            $$=$1
            }

IDD: ID {$$.node=newAst(FUNC,$1.text,types.Any,$1.offset);};

cast_func: CAST IDD    { $$.node =newAst(FUNC,"to"+$2.node.Value,types.Any,$2.offset);}

e_list:   e  {$$.node =newAst(NULL,"",types.Any,$1.offset,$1.node);}
        | e_list COMMA e  {$$.node=newAst(NULL,"",types.Any,$3.offset,append($1.node.Children,$3.node)...);}
