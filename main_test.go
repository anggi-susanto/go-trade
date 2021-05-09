package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_processBody(t *testing.T) {
	t.Parallel()

	t.Run("#1.", func(t *testing.T) {
		cekStr := "5 4 3 2 1"
		diff, _ := processBody(cekStr)
		exp := 0
		assert.Equal(t, exp, diff)
	})
	t.Run("#2.", func(t *testing.T) {
		cekStr := "3 2 1 5 6 2"
		diff, _ := processBody(cekStr)
		exp := 5
		assert.Equal(t, exp, diff)
	})
}
