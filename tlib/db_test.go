// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
)

func Test_Db_String(t *testing.T) {
	pgc := &PGConn{
		Host:     "a",
		Port:     "2",
		User:     "c",
		Password: "d",
		Dbname:   "e",
		Sslmode:  "f",
	}
	assert.Equal(t,
		"host=a port=2 user=c password=d dbname=e sslmode=f",
		pgc.String())
}

func Test_Db_Peek(t *testing.T) {
	pgc := &PGConn{
		Host:     "a",
		Port:     "2",
		User:     "c",
		Password: "d",
		Dbname:   "e",
		Sslmode:  "f",
	}
	assert.Equal(t,
		"host=a port=2 user=c dbname=e sslmode=f",
		pgc.Peek())
}

func Test_Db_Validate_Pass(t *testing.T) {
	pgc := &PGConn{
		Host:     "a",
		Port:     "2",
		User:     "c",
		Password: "d",
		Dbname:   "e",
		Sslmode:  "f",
	}
	assert.Nil(t, pgc.Validate())
}

func Test_Db_Validate_Fail(t *testing.T) {
	pgc := &PGConn{}
	err := pgc.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, 6, len(err.(validator.ValidationErrors)))

	pgc = &PGConn{
		Host:     "a",
		Port:     "b", // must be numeric
		User:     "c",
		Password: "d",
		Dbname:   "e",
		Sslmode:  "f",
	}
	err = pgc.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err.(validator.ValidationErrors)))
}
