package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringUnpacking(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		var data = []string{"a4bc2d5e", "abcd", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
		var correct = []string{"aaaabccddddde", "abcd", "", "qwe45", "qwe44444", "qwe\\\\\\\\\\"}
		for i, v := range data {
			str, err := StringUnpacking(v)
			assert.Equal(t, str, correct[i])
			assert.NoError(t, err)
		}
	})

	t.Run("Incorrect", func(t *testing.T) {
		data := "45"
		_, err := StringUnpacking(data)
		assert.Error(t, err)
	})

}
