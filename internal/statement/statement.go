package statement

import (
	"errors"
	"fmt"
	"strings"
)

type MetaCommandResult int
type PrepareResult int
type StatementType int

const (
	META_COMMAND_SUCCESS MetaCommandResult = iota
	META_COMMAND_UNRECOGNIZED
)

const (
	PREPARE_SUCCESS PrepareResult = iota
	PREPARE_UNRECOGNIZED
)

const (
	STATEMENT_INSERT StatementType = iota
	STATEMENT_SELECT
)

type Statement struct {
	Type StatementType
}

func DoMetaCommand(input string) MetaCommandResult {
	if input == ".exit" {
		return META_COMMAND_SUCCESS
	}
	return META_COMMAND_UNRECOGNIZED
}

func PrepareStatement(input string, stmt *Statement) PrepareResult {
	if strings.HasPrefix(input, "insert") {
		stmt.Type = STATEMENT_INSERT
		return PREPARE_SUCCESS
	}

	if strings.HasPrefix(input, "select") {
		stmt.Type = STATEMENT_SELECT
		return PREPARE_SUCCESS
	}

	return PREPARE_UNRECOGNIZED

}

func ExecuteStatement(stmt *Statement) error {
	switch stmt.Type {
	case STATEMENT_INSERT:
		fmt.Println("Insert statement executed")
	case STATEMENT_SELECT:
		fmt.Println("Select statement executed")
	default:
		return errors.New("Unknown statement type")
	}
	return nil
}
