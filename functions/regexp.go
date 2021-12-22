package functions

import (
	"errors"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/yjhatfdu/expr/types"
	"strings"
	"unicode/utf8"
)

type regexpReplaceAllIndex struct {
	regexp *regexp2.Regexp
	start  int
	end    int
}

func (s *regexpReplaceAllIndex) Init(consts []string, env map[string]string) error {
	return nil
}
func (s *regexpReplaceAllIndex) Handle(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
	if s.regexp == nil {
		r := vectors[1].Index(0).(string)
		var err error
		s.regexp, err = regexp2.Compile(r, regexp2.None)
		if err != nil {
			return nil, err
		}
	}

	var needSub bool
	var unlimitedEnd bool
	var startIndex int
	var endIndex int

	if len(vectors) == 4 {
		startIndex = int(vectors[3].Index(0).(int64))
		needSub = true
		unlimitedEnd = true
	} else if len(vectors) == 5 {
		startIndex = int(vectors[3].Index(0).(int64))
		endIndex = int(vectors[4].Index(0).(int64))
		needSub = true

		if endIndex < 0 {
			return nil, errors.New(fmt.Sprintf("regexpReplace 函数的 endIndex 必须大于等于 0, 实际值 startIndex: %d, endIndex: %d", startIndex, endIndex))
		}

		if startIndex == endIndex {
			endIndex += 1
		}

		if startIndex > endIndex {
			return nil, errors.New(fmt.Sprintf("regexpReplace 函数的 startIndex 必须小于等于 endIndex, 实际值 startIndex: %d, endIndex: %d", startIndex, endIndex))
		}
	} else {
		needSub = false
	}

	if needSub {
		if startIndex < 0 {
			return nil, errors.New(fmt.Sprintf("regexpReplace 函数的 startIndex 必须大于等于 0, 实际值 startIndex: %d", startIndex))
		}
	}

	replace := vectors[2].Index(0).(string)
	input := vectors[0].(*types.NullableText)
	out := &types.NullableText{}
	return BroadCast1(vectors[0], out, func(i int) error {
		var realEnd int
		v := input.Index(i).(string)
		if needSub {
			if unlimitedEnd {
				realEnd = len(v)
			} else {
				if len(v) <= endIndex {
					realEnd = len(v)
				} else {
					realEnd = endIndex
				}
			}
		}

		if needSub {
			sub := string([]rune(v)[startIndex:realEnd])
			str, err := s.regexp.Replace(sub, replace, 0, -1)
			if err != nil {
				return err
			}
			out.Seti(i, string([]rune(v)[0:startIndex])+str+string([]rune(v)[realEnd:]))
		} else {
			str, err := s.regexp.Replace(v, replace, 0, -1)
			if err != nil {
				return err
			}
			out.Seti(i, str)
		}

		return nil
	})
}

type likeFunc struct {
	regexp *regexp2.Regexp
	not    bool
}

func (s *likeFunc) Init(cons []string, env map[string]string) error {
	if len(cons) != 2 {
		return errors.New(fmt.Sprintf("like 函数仅接受两个参数，实际参数个数 %d", len(cons)))
	}

	pattern := cons[1]
	re, err := regexp2.Compile(strings.ReplaceAll(strings.ReplaceAll(pattern, "%", ".*?"), "_", "."), regexp2.None)
	if err != nil {
		return errors.New(fmt.Sprintf("未能成功编译 like 函数表达式 %s，异常信息 %s", cons[1], err))
	}
	s.regexp = re
	return nil
}
func (s *likeFunc) Handle(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
	input := vectors[0].(*types.NullableText)
	output := &types.NullableBool{}
	return BroadCast1(input, output, func(i int) error {
		if s.not {
			matched, err := s.regexp.MatchString(input.Index(i).(string))
			if err != nil {
				return err
			}
			output.Seti(i, !matched)
		} else {
			matched, err := s.regexp.MatchString(input.Index(i).(string))
			if err != nil {
				return err
			}
			output.Seti(i, matched)
		}
		return nil
	})
}

func init() {
	trim, _ := NewFunction("trim")
	trim.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.TrimSpace(input.Index(i).(string)), false)
			return nil
		})
	})
	trim.Overload([]types.BaseType{types.Text, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableText)
		t := vectors[1].(*types.NullableText).Index(0).(string)
		return BroadCast1(input, output, func(i int) error {
			var v = input.Index(i).(string)
			for strings.HasPrefix(v, t) {
				v = strings.TrimLeft(v, t)
			}

			for strings.HasSuffix(v, t) {
				v = strings.TrimSuffix(v, t)
			}

			output.Seti(i, v)
			return nil
		})
	})

	length, _ := NewFunction("length")
	length.Overload([]types.BaseType{types.Text}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, int64(utf8.RuneCountInString(input.Index(i).(string))), false)
			return nil
		})
	})

	toLower, _ := NewFunction("toLower")
	toLower.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.ToLower(input.Index(i).(string)), false)
			return nil
		})
	})

	toUpper, _ := NewFunction("toUpper")
	toUpper.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.ToUpper(input.Index(i).(string)), false)
			return nil
		})
	})

	like, _ := NewFunction("like")
	like.OverloadHandler(
		[]types.BaseType{types.Text, types.TextS},
		types.Bool,
		func() IHandler { return &likeFunc{} },
	)
	notLike, _ := NewFunction("notLike")
	notLike.OverloadHandler(
		[]types.BaseType{types.Text, types.TextS},
		types.Bool,
		func() IHandler { return &likeFunc{not: true} },
	)

	contains, _ := NewFunction("contains")
	contains.Overload([]types.BaseType{types.Text, types.Text}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		containsStr := vectors[1].(*types.NullableText).Index(0).(string)
		output := &types.NullableBool{}
		return BroadCast1(input, output, func(i int) error {
			output.Seti(i, strings.Contains(input.Index(i).(string), containsStr))
			return nil
		})
	})

	// similarTo
	// this function is for SQL Style `SIMILAR TO` operator
	// similar to LIKE, except interprets the pattern using SQL standard's
	// _ and % also can be used as wildcard characters.
	// [(comparable to . and .*) so we can just replace _ and % with . and .*]
	// according to https://postgresql.org/docs/13/functions-matching.html
	_similarTo, _ := NewFunction("similarTo")
	_similarTo.Overload(
		[]types.BaseType{types.Text, types.TextS},
		types.Bool,
		func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
			text := vectors[0].(*types.NullableText)
			pattern := vectors[1].(*types.NullableText).Index(0).(string)
			output := &types.NullableBool{}

			// copy from https://github.com/postgres/postgres/blob/master/src/backend/utils/adt/regexp.c#L833
			// take a look.
			pattern = `^(?:` + pattern + `)$`
			var r, err = regexp2.Compile(strings.ReplaceAll(strings.ReplaceAll(pattern, "%", ".*"), "_", "."), regexp2.None)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("SimilarTo函数匹配语法未能编译成正确的正则表达式, %v %v", pattern, err.Error()))
			}

			return BroadCast1(text, output, func(i int) error {
				matched, err := r.MatchString(text.Index(i).(string))
				if err != nil {
					return err
				}
				output.Seti(i, matched)
				return nil
			})
		})

	_notSimilarTo, _ := NewFunction("notSimilarTo")
	_notSimilarTo.Overload(
		[]types.BaseType{types.Text, types.TextS},
		types.Bool,
		func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
			text := vectors[0].(*types.NullableText)
			pattern := vectors[1].(*types.NullableText).Index(0).(string)
			output := &types.NullableBool{}

			pattern = `^(?:` + pattern + `)$`
			var r, err = regexp2.Compile(strings.ReplaceAll(strings.ReplaceAll(pattern, "%", ".*"), "_", "."), regexp2.None)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("SimilarTo函数匹配语法未能编译成正确的正则表达式, %v %v", pattern, err.Error()))
			}

			return BroadCast1(text, output, func(i int) error {
				matched, err := r.MatchString(text.Index(i).(string))
				if err != nil {
					return err
				}
				output.Seti(i, !matched)
				return nil
			})
		})

	replaceAll, _ := NewFunction("regexpReplace")
	replaceAll.OverloadHandler(
		[]types.BaseType{types.Text, types.TextS, types.TextS},
		types.Text,
		func() IHandler { return &regexpReplaceAllIndex{} },
	)
	replaceAll.OverloadHandler(
		[]types.BaseType{types.Text, types.TextS, types.TextS, types.IntS},
		types.Text,
		func() IHandler { return &regexpReplaceAllIndex{} },
	)
	replaceAll.OverloadHandler(
		[]types.BaseType{types.Text, types.TextS, types.TextS, types.IntS, types.IntS},
		types.Text,
		func() IHandler { return &regexpReplaceAllIndex{} },
	)

	regexpMatch, _ := NewFunction("regexpMatch")
	regexpMatch.Overload([]types.BaseType{types.Text, types.TextS}, types.TextA, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTextArray{}
		input := vectors[0].(*types.NullableText)
		reStr := vectors[1].Index(0).(string)
		re, err := regexp2.Compile(reStr, regexp2.None)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("regexpMatch函数未能成功编译正则表达式 %s, %s", reStr, err.Error()))
		}
		return BroadCast1(vectors[0], output, func(i int) error {
			var ret = make([]string, 0)
			match, err := re.FindStringMatch(input.Index(i).(string))
			for err == nil {
				if match != nil {
					ret = append(ret, match.String())
					match, err = re.FindNextMatch(match)
				} else {
					break
				}
			}
			if err != nil {
				return err
			}

			output.Seti(i, ret)
			return nil
		})
	})
}
