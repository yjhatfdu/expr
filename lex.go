package expr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

//go:generate go get golang.org/x/tools/cmd/goyacc
//go:generate goyacc parser.y

var keywords = map[string]int{
	"+":       ADD,
	"-":       MINUS,
	"*":       MUL,
	"/":       DIV,
	">":       GT,
	"gt":      GT,
	">=":      GTE,
	"gte":     GTE,
	"=":       EQ,
	"==":      EQ,
	"aeq":     EQ,
	"<":       LT,
	"<=":      LTE,
	"lte":     LTE,
	"!=":      NEQ,
	"!==":     NEQ,
	"ne":      NEQ,
	"and":     AND,
	"AND":     AND,
	"or":      OR,
	"OR":      OR,
	"!":       NOT,
	"not":     NOT,
	"NOT":     NOT,
	"like":    LIKE,
	"LIKE":    LIKE,
	"true":    BOOL,
	"TRUE":    BOOL,
	"FALSE":   BOOL,
	"false":   BOOL,
	",":       COMMA,
	"(":       LP,
	")":       RP,
	"null":    NULL,
	"$":       DOLLAR,
	":":       CAST,
	"|":       PIPE,
	"case":    CASE,
	"when":    WHEN,
	"then":    THEN,
	"end":     END,
	"in":      IN,
	"else":    ELSE,
	"is":      IS,
	"similar": SIMILAR,
	"to":      TO,
}

type Lexer struct {
	s           string
	state       int
	pos         int
	bufStartPos int
	parseResult *AstNode
}

type token struct {
	t   int
	s   string
	pos int
}

const EOF = 0

const (
	StateNone = iota
	StateToken
	StateQuote
	StateDoubleQuote
	StateRawQuote
	StateNumber
)

var terminators = map[string]int{
	"+":   ADD,
	"-":   MINUS,
	"*":   MUL,
	"/":   DIV,
	">":   GT,
	">=":  GTE,
	"=":   EQ,
	"==":  EQ,
	"<":   LT,
	"<=":  LTE,
	"!=":  NEQ,
	"!==": NEQ,
	"<>":  NEQ,
	"!":   NOT,
	",":   COMMA,
	"(":   LP,
	")":   RP,
	"$":   DOLLAR,
	":":   CAST,
	"|":   PIPE,
}
var maxKeywordsLength = 3

func init() {
	yyErrorVerbose = true
	for k, _ := range keywords {
		if len(k) > maxKeywordsLength {
			maxKeywordsLength = len(k)
		}
	}
}

func NewLexer(s string) *Lexer {
	return &Lexer{
		s: s,
	}
}

func (l *Lexer) lookAheadKeyword() string {
	for i := l.pos + 3; i > l.pos; i-- {
		if i > len(l.s) {
			continue
		}
		k := strings.ToLower(l.s[l.pos:i])
		if terminators[k] > 0 {
			return k
		}
	}
	return ""
}

func (l *Lexer) lookAhead(n int) string {
	if l.pos+n > len(l.s) {
		return ""
	} else {
		return l.s[l.pos : l.pos+n]
	}
}

func (l *Lexer) Lex(lval *yySymType) (out int) {
	t, err := l.lex()
	if err != nil {
		l.Error(err.Error())
		return 0
	}
	//fmt.Println(t)
	lval.offset = t.pos
	lval.text = t.s
	if t.t == 0 {
		return 0
	}
	if t.t != STR && t.t != INT && t.t != FLOAT {
		if s, ok := keywords[strings.ToLower(t.s)]; ok {
			return s
		} else {
			return ID
		}
	} else {
		return t.t
	}
}

func (l *Lexer) Error(s string) {
	errInfo := fmt.Sprintf("\n%s\n%s\n", l.s, strings.Repeat(" ", l.pos)+"^")
	panic(errInfo + s)
}

func (l *Lexer) lex() (token, error) {
	for {
		switch l.state {
		case StateNone:
			if s := l.lookAheadKeyword(); s != "" {
				l.pos += len(s)
				l.bufStartPos = l.pos
				return token{
					t:   keywords[s],
					s:   s,
					pos: l.pos - len(s),
				}, nil
			}
			switch s := l.lookAhead(1); {
			case s == "'":
				l.state = StateQuote
				l.bufStartPos = l.pos
			case s == "\"":
				l.state = StateDoubleQuote
				l.bufStartPos = l.pos
			case s == "`":
				l.state = StateRawQuote
				l.bufStartPos = l.pos
			case s == "":
				return token{
					t:   EOF,
					s:   "",
					pos: l.pos + 1,
				}, nil
			case s[0] <= '9' && s[0] >= '0':
				l.state = StateNumber
				l.bufStartPos = l.pos
			case s == "\n" || s == " " || s == "\r":
				l.pos++
				continue
			default:
				s0 := s[0]
				if s0 >= 'a' && s0 <= 'z' || s0 >= 'A' && s0 <= 'Z' || s0 == '_' {
					l.state = StateToken
					l.bufStartPos = l.pos
				} else {
					return token{}, fmt.Errorf("\n"+l.s+"\n"+strings.Repeat(" ", l.pos)+"^\n"+
						strings.Repeat(" ", l.pos)+"unexpected token %s", s)
				}
			}
			l.pos++
		case StateQuote:
			switch s := l.lookAhead(1); {
			case s == "'":
				switch s2 := l.lookAhead(2); {
				case s2 == "" || s2[1] != '\'':
					l.pos++
					l.state = StateNone
					str, err := unquote(l.s[l.bufStartPos:l.pos])
					if err != nil {
						err = fmt.Errorf("%v: %s", err, l.s[l.bufStartPos:l.pos])
					}
					return token{
						t:   STR,
						s:   str,
						pos: l.bufStartPos,
					}, err
				case s2 == "''":
					l.pos += 2
				}
			case s == "":
				return token{}, errors.New("unexpected EOF, incomplete string")
			case s == "\\":
				l.pos += 2
			default:
				l.pos++
			}
		case StateDoubleQuote:
			s := l.lookAhead(1)
			switch {
			case s == "\"":
				switch s2 := l.lookAhead(2); {
				case s2 == "" || s2[1] != '"':
					l.pos++
					l.state = StateNone
					str, err := unquoteDouble(l.s[l.bufStartPos:l.pos])
					if err != nil {
						err = fmt.Errorf("%v: %s", err, l.s[l.bufStartPos:l.pos])
					}
					return token{
						t:   STR,
						s:   str,
						pos: l.bufStartPos,
					}, err
				case s2 == "\"\"":
					l.pos += 2
				}
			case s == "":
				return token{}, errors.New("unexpected EOF, incomplete string")
			case s == "\\":
				l.pos += 2
			default:
				l.pos++
			}
		case StateRawQuote:
			s := l.lookAhead(1)
			switch {
			case s == "`":
				switch s2 := l.lookAhead(2); {
				case s2 == "" || s2[1] != '`':
					l.pos++
					l.state = StateNone

					return token{
						t:   STR,
						s:   strings.ReplaceAll(l.s[l.bufStartPos+1:l.pos-1], "``", "`"),
						pos: l.bufStartPos,
					}, nil
				case s2 == "``":
					l.pos += 2
				}
			case s == "":
				return token{}, errors.New("unexpected EOF, incomplete string")
			default:
				l.pos++
			}
		case StateToken:
			if s := l.lookAheadKeyword(); s != "" {
				l.state = StateNone
				return token{
					t:   ID,
					s:   l.s[l.bufStartPos:l.pos],
					pos: l.bufStartPos,
				}, nil
			}
			s := l.lookAhead(1)
			if s == "" || s == "'" || s == "\"" || s == "`" || s == " " || s == "\r" || s == "\n" {
				l.state = StateNone
				return token{
					t:   ID,
					s:   l.s[l.bufStartPos:l.pos],
					pos: l.bufStartPos,
				}, nil
			}
			l.pos++
		case StateNumber:
			if s := l.lookAheadKeyword(); s != "" {
				l.state = StateNone
				numberText := l.s[l.bufStartPos:l.pos]
				if strings.Contains(numberText, ".") {
					return token{
						t:   FLOAT,
						s:   numberText,
						pos: l.bufStartPos,
					}, nil
				} else {
					return token{
						t:   INT,
						s:   numberText,
						pos: l.bufStartPos,
					}, nil
				}
			}
			s := l.lookAhead(1)
			if s == "" || !(s[0] <= '9' && s[0] >= '0' || s[0] == '.') {
				l.state = StateNone
				numberText := l.s[l.bufStartPos:l.pos]
				if strings.Contains(numberText, ".") {
					return token{
						t:   FLOAT,
						s:   numberText,
						pos: l.bufStartPos,
					}, nil
				} else {
					return token{
						t:   INT,
						s:   numberText,
						pos: l.bufStartPos,
					}, nil
				}
			}
			l.pos++
		}
	}
}

func unquote(s string) (string, error) {
	if s == `''` {
		return "", nil
	}
	return _unquote(strings.ReplaceAll(s, "''", "\\'"))
}

func unquoteDouble(s string) (string, error) {
	if s == `""` {
		return "", nil
	}
	return _unquote(strings.ReplaceAll(s, `""`, `"`))

}

var ErrSyntax = errors.New("invalid quote syntax")

func contains(s string, c byte) bool {
	return strings.IndexByte(s, c) != -1
}

func _unquote(s string) (string, error) {
	n := len(s)
	if n < 2 {
		return "", ErrSyntax
	}
	quote := s[0]
	if quote != s[n-1] {
		return "", ErrSyntax
	}
	s = s[1 : n-1]

	if quote == '`' {
		if contains(s, '`') {
			return "", ErrSyntax
		}
		if contains(s, '\r') {
			// -1 because we know there is at least one \r to remove.
			buf := make([]byte, 0, len(s)-1)
			for i := 0; i < len(s); i++ {
				if s[i] != '\r' {
					buf = append(buf, s[i])
				}
			}
			return string(buf), nil
		}
		return s, nil
	}
	if quote != '"' && quote != '\'' {
		return "", ErrSyntax
	}
	if contains(s, '\n') {
		return "", ErrSyntax
	}

	// Is it trivial? Avoid allocation.
	if !contains(s, '\\') && !contains(s, quote) {
		switch quote {
		case '"':
			if utf8.ValidString(s) {
				return s, nil
			}
		case '\'':
			r, size := utf8.DecodeRuneInString(s)
			if size == len(s) && (r != utf8.RuneError || size != 1) {
				return s, nil
			}
		}
	}

	var runeTmp [utf8.UTFMax]byte
	buf := make([]byte, 0, 3*len(s)/2) // Try to avoid more allocations.
	for len(s) > 0 {
		c, multibyte, ss, err := strconv.UnquoteChar(s, quote)
		if err != nil {
			return "", err
		}
		s = ss
		if c < utf8.RuneSelf || !multibyte {
			buf = append(buf, byte(c))
		} else {
			n := utf8.EncodeRune(runeTmp[:], c)
			buf = append(buf, runeTmp[:n]...)
		}
	}
	return string(buf), nil
}
