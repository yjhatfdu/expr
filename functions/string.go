package functions

import (
	"github.com/yjhatfdu/expr/types"
	"regexp"
	"strings"
)

type similarToFunc struct {
	regexp *regexp.Regexp
}

func (s *similarToFunc) Handle(vectors []types.INullableVector) (types.INullableVector, error) {
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

func init() {
	trim, _ := NewFunction("trim")
	trim.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.TrimSpace(input.Values[i]), false)
			return nil
		})
	})
	length, _ := NewFunction("length")
	length.Overload([]types.BaseType{types.Text}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, int64(len(input.Values[i])), false)
			return nil
		})
	})
	toLower, _ := NewFunction("toLower")
	toLower.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.ToLower(input.Values[i]), false)
			return nil
		})
	})
	toUpper, _ := NewFunction("toUpper")
	toUpper.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableText)
		output := &types.NullableText{}
		return BroadCast1(input, output, func(i int) error {
			output.Set(i, strings.ToUpper(input.Values[i]), false)
			return nil
		})
	})
	contains, _ := NewFunction("contains")
	contains.Overload([]types.BaseType{types.Text, types.Text}, types.Bool, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
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
	similar.OverloadHandler([]types.BaseType{types.Text, types.TextS}, types.Bool, new(similarToFunc))
}
