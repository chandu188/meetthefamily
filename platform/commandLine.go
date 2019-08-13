package platform

import (
	"bufio"
	"fmt"
	"io"
)

// IOReadWriter is an interface for retrieveing a line from the input
//and writing a line to output
type IOReadWriter interface {
	RetrieveInput() (string, error)
	WriteOutput(s string) error
}

type commandLine struct {
	input  *bufio.Scanner
	output *bufio.Writer
}

// NewPlatform returns an implementaion of IoReadWriter, which reads a line from r, and
// writes the given string to w
func NewPlatform(r io.Reader, w io.Writer) IOReadWriter {
	return &commandLine{
		input:  bufio.NewScanner(r),
		output: bufio.NewWriter(w),
	}
}

func (c *commandLine) RetrieveInput() (string, error) {
	var str string
	var err error
	scanner := c.input
	err = scanner.Err()
	if scanner.Scan() {
		str = scanner.Text()
	} else {
		err = io.EOF
	}
	return str, err
}

func (c *commandLine) WriteOutput(s string) error {
	_, err := fmt.Fprintln(c.output, s)
	if err != nil {
		return err
	}
	return c.output.Flush()
}
