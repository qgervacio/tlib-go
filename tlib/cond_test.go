// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cond_Tiff(t *testing.T) {
	assert.Equal(t, Tiff(true, "a", "b"), "a")
	assert.Equal(t, Tiff(false, "a", "b"), "b")
	assert.Equal(t, Tiff(true, 1, 2), 1)
}
