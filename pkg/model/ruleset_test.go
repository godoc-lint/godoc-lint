package model_test

import (
	"testing"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestRuleSetListReturnsSortedSlice(t *testing.T) {
	require := require.New(t)
	rs := model.RuleSet{}
	rs = rs.Add("z")
	rs = rs.Add("a")
	require.Equal([]string{"a", "z"}, rs.List())
}

func TestRuleSetListReturnsEmptySlice(t *testing.T) {
	require := require.New(t)
	rs := model.RuleSet{}
	require.Equal([]string{}, rs.List())
}

func TestRuleSetHas(t *testing.T) {
	require := require.New(t)

	rs := model.RuleSet{}
	require.False(rs.Has(""))
	require.False(rs.Has("unknown-rule"))

	rs = rs.Add("foo", "bar")
	require.True(rs.Has("foo"))
	require.True(rs.Has("bar"))
	require.False(rs.Has(""))
	require.False(rs.Has("unknown-rule"))
}

func TestRuleSetMerge(t *testing.T) {
	require := require.New(t)

	rs := model.RuleSet{}
	anotherEmpty := model.RuleSet{}
	anotherNonEmpty := anotherEmpty.Add("foo", "bar")

	rs = rs.Merge(anotherEmpty)
	require.Equal([]string{}, rs.List())

	rs = rs.Merge(anotherNonEmpty)
	require.Equal([]string{"bar", "foo"}, rs.List())

	rs = rs.Merge(rs)
	require.Equal([]string{"bar", "foo"}, rs.List())
}

func TestRuleSetAdd(t *testing.T) {
	require := require.New(t)

	rs := model.RuleSet{}
	rs = rs.Add("foo")
	require.Equal([]string{"foo"}, rs.List())

	rs = rs.Add("bar")
	require.Equal([]string{"bar", "foo"}, rs.List())

	rs = rs.Add("baz", "baz", "yolo")
	require.Equal([]string{"bar", "baz", "foo", "yolo"}, rs.List())
}

func TestRuleSetIsSupersetOf(t *testing.T) {
	require := require.New(t)

	rs := model.RuleSet{}.Add("foo", "bar")

	require.True(rs.IsSupersetOf(rs))
	require.True(rs.IsSupersetOf(model.RuleSet{}))
	require.True(rs.IsSupersetOf(model.RuleSet{}.Add("foo")))
	require.True(rs.IsSupersetOf(model.RuleSet{}.Add("foo", "bar")))
	require.False(rs.IsSupersetOf(model.RuleSet{}.Add("yolo")))

	require.True(model.RuleSet{}.IsSupersetOf(model.RuleSet{}))
	require.False(model.RuleSet{}.IsSupersetOf(model.RuleSet{}.Add("foo")))
}

func TestRuleSetHasCommonsWith(t *testing.T) {
	require := require.New(t)

	rs := model.RuleSet{}.Add("foo", "bar")

	require.True(rs.HasCommonsWith(rs))
	require.False(rs.HasCommonsWith(model.RuleSet{}))
	require.True(rs.HasCommonsWith(model.RuleSet{}.Add("foo")))
	require.True(rs.HasCommonsWith(model.RuleSet{}.Add("foo", "bar")))
	require.False(rs.HasCommonsWith(model.RuleSet{}.Add("yolo")))
	require.True(rs.HasCommonsWith(model.RuleSet{}.Add("foo", "yolo")))

	require.False(model.RuleSet{}.HasCommonsWith(model.RuleSet{}))
	require.False(model.RuleSet{}.HasCommonsWith(model.RuleSet{}.Add("foo")))
}
