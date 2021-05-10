package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorSix(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cur := FindFactorSixCount(12)
		assert.Equal(t, 1, cur)
	})
}
