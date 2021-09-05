// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// PGConn represents the PG connection cred type
type PGConn struct {
	Host     string `validate:"required" yaml:"host" json:"host"`
	Port     int    `validate:"required,numeric" yaml:"port"  json:"port"`
	User     string `validate:"required" yaml:"user" json:"user"`
	Password string `validate:"required" yaml:"password" json:"password"`
	Dbname   string `validate:"required" yaml:"name" json:"name"`
	Extras   string `yaml:"extras" json:"extras"`
}

// String returns the string representation of PGConn
func (p *PGConn) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		p.User, p.Password, p.Host, p.Port, p.Dbname, p.Extras)
}

// Peek returns the string representation of PGConn without password
func (p *PGConn) Peek() string {
	return fmt.Sprintf(
		"postgres://%s:***@%s:%d/%s?%s",
		p.User, p.Host, p.Port, p.Dbname, p.Extras)
}

// Validate the values
func (p *PGConn) Validate() error {
	v := validator.New()
	if err := v.Struct(p); err != nil {
		return err
	}
	return nil
}
