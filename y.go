// Code generated by goyacc parser.y. DO NOT EDIT.

//line parser.y:2
package expr

import __yyfmt__ "fmt"

//line parser.y:2
import "github.com/yjhatfdu/expr/types"

//line parser.y:6
type yySymType struct {
	yys    int
	offset int
	node   *AstNode
}

const NULL = 57346
const INT = 57347
const STR = 57348
const RAWSTR = 57349
const BOOL = 57350
const FLOAT = 57351
const CONST = 57352
const OR = 57353
const AND = 57354
const NOT = 57355
const LIKE = 57356
const NEQ = 57357
const GT = 57358
const LT = 57359
const GTE = 57360
const LTE = 57361
const EQ = 57362
const ADD = 57363
const MINUS = 57364
const MUL = 57365
const DIV = 57366
const CONTAINS = 57367
const ID = 57368
const IND = 57369
const COMMA = 57370
const ANY = 57371
const FUNC = 57372
const LP = 57373
const RP = 57374
const DOLLAR = 57375
const VAR = 57376
const CAST = 57377
const PIPE = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NULL",
	"INT",
	"STR",
	"RAWSTR",
	"BOOL",
	"FLOAT",
	"CONST",
	"OR",
	"AND",
	"NOT",
	"LIKE",
	"NEQ",
	"GT",
	"LT",
	"GTE",
	"LTE",
	"EQ",
	"ADD",
	"MINUS",
	"MUL",
	"DIV",
	"CONTAINS",
	"ID",
	"IND",
	"COMMA",
	"ANY",
	"FUNC",
	"LP",
	"RP",
	"DOLLAR",
	"VAR",
	"CAST",
	"PIPE",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 203

var yyAct = [...]int{
	59, 2, 62, 57, 32, 31, 66, 60, 62, 33,
	37, 36, 61, 15, 30, 34, 12, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	53, 19, 18, 35, 16, 29, 24, 26, 25, 27,
	28, 20, 21, 23, 22, 17, 20, 21, 23, 22,
	13, 11, 56, 38, 1, 32, 31, 39, 0, 0,
	32, 31, 0, 65, 64, 19, 18, 0, 16, 29,
	24, 26, 25, 27, 28, 20, 21, 23, 22, 17,
	0, 0, 54, 55, 0, 0, 0, 0, 0, 32,
	31, 18, 0, 16, 29, 24, 26, 25, 27, 28,
	20, 21, 23, 22, 17, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 32, 31, 16, 29, 24, 26,
	25, 27, 28, 20, 21, 23, 22, 17, 3, 4,
	5, 7, 6, 0, 23, 22, 8, 32, 31, 0,
	0, 0, 0, 0, 0, 14, 32, 31, 0, 15,
	16, 0, 0, 0, 10, 63, 9, 20, 21, 23,
	22, 17, 3, 4, 5, 7, 6, 0, 0, 0,
	8, 32, 31, 0, 3, 4, 5, 7, 6, 14,
	0, 0, 8, 15, 0, 0, 0, 0, 10, 58,
	9, 14, 0, 0, 0, 15, 0, 0, 0, 0,
	10, 0, 9,
}

var yyPact = [...]int{
	169, -1000, 54, -1000, -1000, -1000, -1000, -1000, 169, 10,
	169, -1000, -1000, -21, 48, -1000, 169, 169, 169, 169,
	169, 169, 169, 169, 169, 169, 169, 169, 169, 169,
	-1000, -13, -13, 102, -1000, -1000, 20, 157, -1000, -1000,
	25, 25, 102, 79, 111, 111, -31, -31, 136, 136,
	136, 136, 136, 136, -24, -1000, -1000, -20, -1000, 54,
	123, -1000, 169, -1000, -26, 54, -1000,
}

var yyPgo = [...]int{
	0, 54, 0, 51, 16, 50, 3, 14,
}

var yyR1 = [...]int{
	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 4, 4, 3,
	3, 3, 3, 3, 3, 3, 5, 7, 6, 6,
}

var yyR2 = [...]int{
	0, 1, 1, 1, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 3, 3, 3, 3,
	3, 3, 2, 2, 3, 1, 1, 2, 2, 4,
	3, 1, 2, 3, 5, 6, 1, 2, 1, 3,
}

var yyChk = [...]int{
	-1000, -1, -2, 5, 6, 7, 9, 8, 13, 33,
	31, -3, -4, -5, 22, 26, 14, 25, 12, 11,
	21, 22, 24, 23, 16, 18, 17, 19, 20, 15,
	-7, 36, 35, -2, 5, 23, -2, 31, 5, 9,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, -5, -5, 32, -6, 32, -2,
	31, 32, 28, 32, -6, -2, 32,
}

var yyDef = [...]int{
	0, -2, 1, 2, 3, 4, 5, 6, 0, 0,
	0, 25, 26, 31, 0, 36, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	32, 0, 0, 15, 22, 23, 0, 0, 27, 28,
	7, 8, 9, 10, 11, 12, 13, 14, 16, 17,
	18, 19, 20, 21, 33, 37, 24, 0, 30, 38,
	0, 29, 0, 34, 0, 39, 35,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:57
		{
			yylex.(*Lexer).parseResult = yyDollar[1].node
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:59
		{
			yyVAL.node = newAst(CONST, yylex.(*Lexer).Text(), types.Int, yyDollar[1].offset)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:60
		{
			yyVAL.node = newAst(CONST, yylex.(*Lexer).Text(), types.Text, yyDollar[1].offset)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:61
		{
			yyVAL.node = newAst(CONST, unquoteRawString(yylex.(*Lexer).Text()), RAWSTR, yyDollar[1].offset)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:62
		{
			yyVAL.node = newAst(CONST, yylex.(*Lexer).Text(), types.Float, yyDollar[1].offset)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:63
		{
			yyVAL.node = newAst(CONST, yylex.(*Lexer).Text(), types.Bool, yyDollar[1].offset)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:64
		{
			yyVAL.node = newAst(FUNC, "like", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:65
		{
			yyVAL.node = newAst(FUNC, "contains", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:66
		{
			yyVAL.node = newAst(FUNC, "and", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:67
		{
			yyVAL.node = newAst(FUNC, "or", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:68
		{
			yyVAL.node = newAst(FUNC, "add", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:69
		{
			yyVAL.node = newAst(FUNC, "minus", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:70
		{
			yyVAL.node = newAst(FUNC, "div", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:71
		{
			yyVAL.node = newAst(FUNC, "mul", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:72
		{
			yyVAL.node = newAst(FUNC, "not", types.Any, yyDollar[1].offset, yyDollar[2].node)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:73
		{
			yyVAL.node = newAst(FUNC, "gt", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:74
		{
			yyVAL.node = newAst(FUNC, "gte", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:75
		{
			yyVAL.node = newAst(FUNC, "lt", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:76
		{
			yyVAL.node = newAst(FUNC, "lte", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:77
		{
			yyVAL.node = newAst(FUNC, "eq", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:78
		{
			yyVAL.node = newAst(FUNC, "neq", types.Any, yyDollar[2].offset, yyDollar[1].node, yyDollar[3].node)
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:79
		{
			yyVAL.node = newAst(VAR, yylex.(*Lexer).Text(), types.Any, yyDollar[1].offset)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:80
		{
			yyVAL.node = newAst(VAR, "ALL", types.Any, yyDollar[1].offset)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:81
		{
			yyVAL.node = yyDollar[2].node
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:82
		{
			yyVAL.node = yyDollar[1].node
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:83
		{
			yyVAL.node = yyDollar[1].node
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:85
		{
			yyVAL.node = newAst(CONST, "-"+yylex.(*Lexer).Text(), types.Int, yyDollar[2].offset)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:86
		{
			yyVAL.node = newAst(CONST, "-"+yylex.(*Lexer).Text(), types.Float, yyDollar[2].offset)
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:88
		{
			yyVAL.node = newAst(FUNC, yyDollar[1].node.Value, types.Any, yyDollar[1].offset, yyDollar[3].node.Children...)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:89
		{
			yyVAL.node = newAst(FUNC, yyDollar[1].node.Value, types.Any, yyDollar[1].offset)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:90
		{
			yyVAL.node = newAst(FUNC, yyDollar[1].node.Value, types.Any, yyDollar[1].offset)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:91
		{
			yyVAL.node = newAst(FUNC, yyDollar[2].node.Value, types.Any, yyDollar[2].offset, yyDollar[1].node)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:92
		{
			yyDollar[3].node.Children = append([]*AstNode{yyDollar[1].node}, yyDollar[3].node.Children...)
			yyVAL.node = yyDollar[3].node
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:93
		{
			yyDollar[3].node.Children = append([]*AstNode{yyDollar[1].node}, yyDollar[3].node.Children...)
			yyVAL.node = yyDollar[3].node
		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:94
		{
			yyVAL.node = newAst(FUNC, yyDollar[3].node.Value, types.Any, yyDollar[3].offset, append([]*AstNode{yyDollar[1].node}, yyDollar[5].node.Children...)...)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:96
		{
			yyVAL.node = newAst(FUNC, yylex.(*Lexer).Text(), types.Any, yyDollar[1].offset)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:98
		{
			yyVAL.node = newAst(FUNC, "to"+yyDollar[2].node.Value, types.Any, yyDollar[2].offset)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:100
		{
			yyVAL.node = newAst(NULL, "", types.Any, yyDollar[1].offset, yyDollar[1].node)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:101
		{
			yyVAL.node = newAst(NULL, "", types.Any, yyDollar[3].offset, append(yyDollar[1].node.Children, yyDollar[3].node)...)
		}
	}
	goto yystack /* stack new state and value */
}
