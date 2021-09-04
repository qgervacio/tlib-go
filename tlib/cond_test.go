// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cond_Tiff(t *testing.T) {
	assert.NotNil(t, Tiff(true, "a", "b"), "a")
	assert.NotNil(t, Tiff(false, "a", "b"), "b")
	assert.NotNil(t, Tiff(true, 1, 2), 1)
}
