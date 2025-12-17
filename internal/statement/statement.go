package statement

import (
	"errors"
	"fmt"
	"strings"

	"github.com/irfan378/MiniDB/internal/db"
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
	STATEMENT_CREATE_TABLE
)

type Statement struct {
	Type      StatementType
	TableName string
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

	if strings.HasPrefix(input, "create table") {
		tokens := strings.Fields(input)
		if len(tokens) >= 3 {
			stmt.Type = STATEMENT_CREATE_TABLE
			stmt.TableName = tokens[2]
			return PREPARE_SUCCESS
		}
	}

	return PREPARE_UNRECOGNIZED
}

func ExecuteStatement(stmt *Statement, database *db.Database) error {

	switch stmt.Type {
	case STATEMENT_CREATE_TABLE:
		if err := database.CreateTable(stmt.TableName); err != nil {
			fmt.Println("Error creating table:", err)
		} else {
			fmt.Println("Table created.")
		}

	case STATEMENT_INSERT:
		fmt.Println("Insert statement executed")

	case STATEMENT_SELECT:
		fmt.Println("Select statement executed")

	default:
		return errors.New("Unknown statement type")
	}

	return nil
}
