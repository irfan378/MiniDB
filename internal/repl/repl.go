package repl

import (
"fmt"
"os"
"github.com/irfan378/MiniDB/internal/input"
)

func Start(){
	inputBuffer:=input.NewInputBuffer()

	for{
		fmt.Print("db> ")

		err:=input.ReadInput(inputBuffer)
		if(err!=nil){
			fmt.Println("Error:",err)
			os.Exit(1)
		}

		if inputBuffer.Buffer==".exit"{
			os.Exit(0)
		}

		fmt.Printf("Unrecognized command '%s'.\n",inputBuffer.Buffer)
		
	}
}
