// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// PGConn represents the PG connection cred type
type PGConn struct {
	Host     string `validate:"required"`
	Port     string `validate:"required,numeric"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Dbname   string `validate:"required"`
	Sslmode  string `validate:"required"`
}

// String returns the string representation of PGConn
func (p *PGConn) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.Dbname, p.Sslmode)
}

// Peek returns the string representation of PGConn without password
func (p *PGConn) Peek() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Dbname, p.Sslmode)
}

// Validate the values
func (p *PGConn) Validate() error {
	v := validator.New()
	if err := v.Struct(p); err != nil {
		return err
	}
	return nil
}
