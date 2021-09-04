// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// PGConn represents the PG connection cred type
type PGConn struct {
	Host     string `validate:"required" yaml:"host" json:"host"`
	Port     string `validate:"required,numeric" yaml:"port"  json:"port"`
	User     string `validate:"required" yaml:"user" json:"user"`
	Password string `validate:"required" yaml:"password" json:"password"`
	Dbname   string `validate:"required" yaml:"name" json:"name"`
	Sslmode  string `validate:"required" yaml:"sslmode" json:"sslmode"`
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
