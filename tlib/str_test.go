// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Env_EnvEval(t *testing.T) {
	assert.Equal(t, EnvEval("$a $b ${} {{}} ${city}, ${state} ${zip} ${city}"),
		"$a $b ${} {{}} ${city}, ${state} ${zip} ${city}")

	_ = os.Setenv("SOME_TEST0", "abc")
	_ = os.Setenv("SOME_TEST1", "123")
	assert.Equal(t, EnvEval("hello ${SOME_TEST0} and ${SOME_TEST1}"),
		"hello abc and 123")
}
