
package input

import (
	"bufio"
	"os"
	"strings"
)

type InputBuffer struct {
	Buffer string
}

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{}
}

func ReadInput(input *InputBuffer) error {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	input.Buffer = strings.TrimRight(line, "\n")
	return nil
}
