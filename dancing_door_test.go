package dancingdoor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {
	smallCorpus := []string{"a"}
	smallSuffixes := []string{"b"}

	t.Run("Default options work", func(t *testing.T) {
		require.Equal(t, " ", MakeOptions().Separator)
		require.Equal(t, "-", MakeOptions().WithSeparator("-").Separator)
	})

	t.Run("Generation works", func(t *testing.T) {
		opts := MakeOptions().WithCorpus(smallCorpus).WithSuffixes(smallSuffixes)
		require.Contains(t, []string{"a a b", "b a", "a b"}, Codename(opts))
	})
}
