package expr

import (
	"regexp"
	"strings"
	"text/scanner"
)

//go:generate go get golang.org/x/tools/cmd/goyacc
//go:generate goyacc parser.y

var keywords = map[string]int{
	"+":        ADD,
	"-":        MINUS,
	"*":        MUL,
	"/":        DIV,
	">":        GT,
	"=":        EQ,
	"==":       EQ,
	"eq":       EQ,
	"<":        LT,
	"!=":       NEQ,
	"not_eq":   NEQ,
	"!==":      NEQ,
	"and":      AND,
	"AND":      AND,
	"or":       OR,
	"OR":       OR,
	"not":      NOT,
	"NOT":      NOT,
	"like":     LIKE,
	"LIKE":     LIKE,
	"contains": CONTAINS,
	"CONTAINS": CONTAINS,
	"true":     BOOL,
	"TRUE":     BOOL,
	"FALSE":    BOOL,
	"false":    BOOL,
	",":        COMMA,
	"(":        LP,
	")":        RP,
	"null":     NULL,
	"$":        DOLLAR,
	"cast":     CAST,
	"::":       CAST,
}

var replaceMap = map[string]string{
	"==": " EQ ",
	"!=": " NOT_EQ ",
	"::": " CAST ",
}
var replacer = regexp.MustCompile("(==|!=|!==|::)")

type Lexer struct {
	s           scanner.Scanner
	parseResult *AstNode
	buffer      string
}

func NewLexer(str string) *Lexer {
	l := Lexer{}
	str = replacer.ReplaceAllStringFunc(str, func(s string) string {
		return replaceMap[strings.ToLower(s)]
	})
	l.s.Init(strings.NewReader(str))
	l.s.Mode = scanner.ScanStrings | scanner.ScanFloats | scanner.ScanInts | scanner.ScanIdents | scanner.SkipComments | scanner.ScanRawStrings
	return &l
}
func (l *Lexer) Text() string {
	r := l.s.TokenText()
	return r
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := l.s.Scan()
	switch token {
	case scanner.EOF:
		return 0
	case scanner.Int:
		return INT
	case scanner.Float:
		return FLOAT
	case scanner.String:
		return STR
	case scanner.RawString:
		return RAWSTR
	case scanner.Ident:
		ident := strings.ToLower(l.Text())
		keyword, ok := keywords[strings.ToLower(ident)]
		if ok {
			return keyword
		} else {
			return ID
		}
	default:
		ident := strings.ToLower(l.Text())
		keyword, ok := keywords[ident]
		if ok {
			return keyword
		} else {
			l.Error("lexer error: " + l.Text())
			return 0
		}
	}
}

func (l *Lexer) Error(s string) {
	panic(s)
}

func unquoteRawString(s string) string {
	return s[1 : len(s)-1]
}
