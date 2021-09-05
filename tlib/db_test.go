// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
)

func Test_Db_String(t *testing.T) {
	pgc := &PGConn{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		Dbname:   "dbname",
		Extras:   "sslMode=disable",
	}
	assert.Equal(t,
		"postgres://user:password@localhost:5432/dbname?sslMode=disable",
		pgc.String())
}

func Test_Db_Peek(t *testing.T) {
	pgc := &PGConn{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		Dbname:   "dbname",
		Extras:   "sslMode=disable",
	}
	assert.Equal(t,
		"postgres://user:***@localhost:5432/dbname?sslMode=disable",
		pgc.Peek())
}

func Test_Db_Validate_Pass(t *testing.T) {
	pgc := &PGConn{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		Dbname:   "dbname",
		Extras:   "sslMode=disable",
	}
	assert.Nil(t, pgc.Validate())
}

func Test_Db_Validate_Fail(t *testing.T) {
	pgc := &PGConn{}
	err := pgc.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, 5, len(err.(validator.ValidationErrors)))

	// missing port
	pgc = &PGConn{
		Host:     "localhost",
		User:     "user",
		Password: "password",
		Dbname:   "dbname",
		Extras:   "sslMode=disable",
	}
	err = pgc.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, 1, len(err.(validator.ValidationErrors)))
}
