package functions

import (
	"github.com/yjhatfdu/expr/types"
	"regexp"
	"strings"
	"unicode/utf8"
)

type similarToFunc struct {
	regexp *regexp.Regexp
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

type replaceAllFunc struct {
	regexp *regexp.Regexp
}

func (s *replaceAllFunc) Handle(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
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
		out.Set(i, string(s.regexp.ReplaceAllString(input.Values[i], replace)), false)
		return nil
	})
}

type regexpMatchFunc struct {
	regexp *regexp.Regexp
	group  int
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
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.TrimSpace(input.Values[i]), false)
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

	replaceAll, _ := NewFunction("replaceAll")
	replaceAll.Comment("support Re2 regexp")
	replaceAll.OverloadHandler(
		[]types.BaseType{types.Text, types.TextS, types.TextS},
		types.Text,
		func() IHandler { return &replaceAllFunc{} },
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
