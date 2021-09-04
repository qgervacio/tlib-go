// Copyright (c) 2021. Quirino Gervacio

package tlib

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	block = regexp.MustCompile(`(\${[^}]*})`)
	token = regexp.MustCompile(`\${(.*)}`)
)

// EnvEval evaluates all matching ${...}
// pattern by replacing it with the actual
// OS env value.
func EnvEval(s string) string {
	evaluated := s
	matches := block.FindAllString(s, -1)
	for _, m := range matches {
		e := token.FindStringSubmatch(m)
		ev := os.Getenv(e[1])
		evaluated = strings.ReplaceAll(evaluated, e[0],
			fmt.Sprintf("%v", Tiff(ev == "", e[0], ev)))
	}
	return evaluated
}
