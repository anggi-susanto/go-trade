package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueString(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cekArr := "bbssdd"
		cur := FirstOccurence(cekArr)
		assert.Equal(t, "bsd", cur)
	})
	t.Run("#2.", func(t *testing.T) {
		cekArr := "bbssddaaa"
		cur := LexicoGraphically(cekArr)
		assert.Equal(t, "abds", cur)
	})
}
