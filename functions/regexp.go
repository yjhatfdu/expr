package functions

import (
	"errors"
	"fmt"
	"github.com/yjhatfdu/expr/types"
	"regexp"
	"strings"
	"unicode/utf8"
)

type similarToFunc struct {
	regexp *regexp.Regexp
}

func (s *similarToFunc) Init([]string, map[string]string) error {
	return nil
}
func (s *similarToFunc) Handle(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
	if s.regexp == nil {
		r := vectors[1].Index(0).(string)
		var err error
		s.regexp, err = regexp.Compile(r)
		if err != nil {
			return nil, err
		}
	}
	input := vectors[0].(*types.NullableText)
	out := &types.NullableBool{}
	return BroadCast1(vectors[0], out, func(i int) error {
		out.Set(i, s.regexp.MatchString(input.Values[i]), false)
		return nil
	})
}

type regexpReplaceAll struct {
	regexp *regexp.Regexp
}

func (s *regexpReplaceAll) Init(consts []string, env map[string]string) error {
	return nil
}
func (s *regexpReplaceAll) Handle(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
	if s.regexp == nil {
		r := vectors[1].Index(0).(string)
		var err error
		s.regexp, err = regexp.Compile(r)
		if err != nil {
			return nil, err
		}
	}
	replace := vectors[2].Index(0).(string)
	input := vectors[0].(*types.NullableText)
	out := &types.NullableText{}
	return BroadCast1(vectors[0], out, func(i int) error {
		out.Set(i, s.regexp.ReplaceAllString(input.Values[i], replace), false)
		return nil
	})
}

type regexpReplaceAllIndex struct {
	regexp *regexp.Regexp
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
		s.regexp, err = regexp.Compile(r)
		if err != nil {
			return nil, err
		}
	}

	var needSub bool
	var startIndex int
	var endIndex int

	if len(vectors) == 4 {
		startIndex = int(vectors[3].Index(0).(int64))
		needSub = true
	} else if len(vectors) == 5 {
		startIndex = int(vectors[3].Index(0).(int64))
		endIndex = int(vectors[4].Index(0).(int64))
		needSub = true
	} else {
		needSub = false
	}

	if needSub {
		if startIndex < 0 || endIndex < 0 {
			return nil, errors.New(fmt.Sprintf("regexpReplace 函数的 startIndex 与 endIndex 必须大于等于 0, 实际值 startIndex: %d, endIndex: %d", startIndex, endIndex))
		}

		if startIndex == endIndex {
			endIndex += 1
		}

		if startIndex > endIndex {
			return nil, errors.New(fmt.Sprintf("regexpReplace 函数的 startIndex 必须小于等于 endIndex, 实际值 startIndex: %d, endIndex: %d", startIndex, endIndex))
		}
	}

	replace := vectors[2].Index(0).(string)
	input := vectors[0].(*types.NullableText)
	out := &types.NullableText{}
	return BroadCast1(vectors[0], out, func(i int) error {
		var realEnd int
		v := input.Index(i).(string)
		if needSub {
			if len(v) <= endIndex {
				realEnd = len(v)
			} else {
				realEnd = endIndex
			}
		}

		if needSub {
			sub := string([]rune(v)[startIndex:realEnd])
			subReplaceStr := s.regexp.ReplaceAllString(sub, replace)
			out.Set(i, string([]rune(v)[0:startIndex])+subReplaceStr+string([]rune(v)[realEnd:]), false)
		} else {
			out.Set(i, s.regexp.ReplaceAllString(v, replace), false)
		}

		return nil
	})
}

type regexpMatchFunc struct {
	regexp *regexp.Regexp
	group  int
}

func (s *regexpMatchFunc) Init([]string, map[string]string) error {
	return nil
}
func (s *regexpMatchFunc) Handle(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
	if s.regexp == nil {
		r := vectors[1].Index(0).(string)
		var err error
		s.regexp, err = regexp.Compile(r)
		if err != nil {
			return nil, err
		}
		if len(vectors) == 3 {
			s.group = int(vectors[2].Index(0).(int64))
		}
	}
	input := vectors[0].(*types.NullableText)
	out := &types.NullableText{}
	return BroadCast1(vectors[0], out, func(i int) error {
		group := s.regexp.FindStringSubmatch(input.Values[i])
		if len(group) > s.group {
			out.Set(i, group[s.group], false)
		} else {
			out.Set(i, "", true)
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
			output.Set(i, strings.TrimSpace(input.Values[i]), false)
			return nil
		})
	})
	trim.Overload([]types.BaseType{types.Text, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableText)
		t := vectors[0].(*types.NullableText).Index(0).(string)
		return BroadCast1(input, output, func(i int) error {
			s := input.Index(i).(string)
			output.Seti(i, strings.TrimSuffix(strings.TrimPrefix(s, t), t))
			return nil
		})
	})
	length, _ := NewFunction("length")
	length.Overload([]types.BaseType{types.Text}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, int64(utf8.RuneCountInString(input.Values[i])), false)
			return nil
		})
	})
	toLower, _ := NewFunction("toLower")
	toLower.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.ToLower(input.Values[i]), false)
			return nil
		})
	})
	toUpper, _ := NewFunction("toUpper")
	toUpper.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.ToUpper(input.Values[i]), false)
			return nil
		})
	})
	contains, _ := NewFunction("contains")
	contains.Overload([]types.BaseType{types.Text, types.Text}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		input2 := vectors[1].(*types.NullableText)
		output := &types.NullableBool{}
		return BroadCast2(input, input2, output, func(index, i, j int) error {
			output.Set(index, strings.Contains(input.Values[i], input2.Values[j]), false)
			return nil
		})
	})
	similar, _ := NewFunction("similar")
	similar.Comment("support Re2 regexp")
	similar.OverloadHandler([]types.BaseType{types.Text, types.TextS}, types.Bool, func() IHandler { return &similarToFunc{} })

	replaceAll, _ := NewFunction("regexpReplace")
	replaceAll.Comment("support Re2 regexp")
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
	regexpMatch.OverloadHandler(
		[]types.BaseType{types.Text, types.Text},
		types.Text,
		func() IHandler {
			return &regexpMatchFunc{}
		},
	)
	regexpMatch.OverloadHandler(
		[]types.BaseType{types.Text, types.Text, types.Int},
		types.Text,
		func() IHandler {
			return &regexpMatchFunc{}
		},
	)
}
