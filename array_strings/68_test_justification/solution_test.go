package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_test(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
		require.Equal(t,
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			}, res)
	})
	t.Run("2", func(t *testing.T) {
		res := fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
		require.Len(t, res, 3)
		require.Equal(t, []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}, res)
	})

	t.Run("3", func(t *testing.T) {
		res := fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
		require.Equal(t, []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to", // 6+1 +2+1 +7+1+2
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
			res,
		)

	})
}
