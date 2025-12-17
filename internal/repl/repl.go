package repl

import (
	"fmt"
	"os"

	"github.com/irfan378/MiniDB/internal/db"
	"github.com/irfan378/MiniDB/internal/input"
	"github.com/irfan378/MiniDB/internal/statement"
)

func Start() {
	database, err := db.Open("minidb.db")
	if err != nil {
		fmt.Println("Failed to open database:", err)
		os.Exit(1)
	}
	inputBuffer := input.NewInputBuffer()

	for {
		fmt.Print("db> ")

		err := input.ReadInput(inputBuffer)
		if err != nil {
			fmt.Println("Error reading Input:", err)
			continue
		}

		if len(inputBuffer.Buffer) > 0 && inputBuffer.Buffer[0] == '.' {
			switch statement.DoMetaCommand(inputBuffer.Buffer) {
			case statement.META_COMMAND_SUCCESS:
				if inputBuffer.Buffer == ".exit" {
					os.Exit(0)
				}
			case statement.META_COMMAND_UNRECOGNIZED:
				fmt.Printf("Unrecognized command '%s'\n", inputBuffer.Buffer)
			}
			continue

		}

		var stmt statement.Statement

		switch statement.PrepareStatement(inputBuffer.Buffer, &stmt) {
		case statement.PREPARE_SUCCESS:

		case statement.PREPARE_UNRECOGNIZED:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", inputBuffer.Buffer)
			continue
		}

		if err := statement.ExecuteStatement(&stmt, database); err != nil {
			fmt.Println("Execution error:", err)
			continue
		}

		fmt.Println("Executed.")
	}
}
