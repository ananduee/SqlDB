package compiler

import (
	"strings"
)

// StatementType - Type of parsed statement
type StatementType int

const (
	INSERT StatementType = iota
	SELECT
	UNRECOGNIZED
)

// Statement to contain parsed information
type Statement struct {
	Type StatementType
}

// Parse raw statement and extract grammer info.
func Parse(rawStatement string) Statement {
	if strings.HasPrefix(rawStatement, "insert") {
		return Statement{Type: INSERT}
	}
	if strings.HasPrefix(rawStatement, "select") {
		return Statement{Type: SELECT}
	}
	return Statement{Type: UNRECOGNIZED}
}
