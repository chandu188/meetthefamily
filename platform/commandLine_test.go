package platform

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandLineIO(t *testing.T) {
	str := "Hello\nworld!!"
	var buf bytes.Buffer
	cmdLine := NewPlatform(strings.NewReader(str), &buf)
	in, err := cmdLine.RetrieveInput()
	assert.Nil(t, err)
	assert.Equal(t, "Hello", in)
	cmdLine.WriteOutput(in)
	in, err = cmdLine.RetrieveInput()
	assert.Nil(t, err)
	assert.Equal(t, "world!!", in)
	cmdLine.WriteOutput(in)
	_, err = cmdLine.RetrieveInput()
	assert.NotNil(t, err)
	assert.Equal(t, "Hello\nworld!!\n", buf.String())
}
