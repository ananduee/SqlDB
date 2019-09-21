package compiler

import (
	"fmt"
	"strings"
)

// StatementType - Type of parsed statement
type StatementType int

const (
	INSERT StatementType = iota
	SELECT
	UNRECOGNIZED
)

// Row to insert to DB.
type Row struct {
	ID       uint32
	Username [32]byte
	Email    [256]byte
}

// Statement to contain parsed information
type Statement struct {
	Type    StatementType
	DataRow Row
}

// Parse raw statement and extract grammer info.
func Parse(rawStatement string) Statement {
	if strings.HasPrefix(rawStatement, "insert") {
		row := Row{}
		var tokenName, userName, email string
		tokenCount, err := fmt.Fscanln(strings.NewReader(rawStatement), &tokenName, &row.ID, &userName, &email)
		if err != nil || tokenCount != 4 {
			return Statement{Type: UNRECOGNIZED}
		}
		// Todo - add validations for input data.
		copy(row.Username[:], []byte(userName))
		copy(row.Email[:], []byte(email))
		return Statement{Type: INSERT, DataRow: row}
	}
	if strings.HasPrefix(rawStatement, "select") {
		return Statement{Type: SELECT}
	}
	return Statement{Type: UNRECOGNIZED}
}
