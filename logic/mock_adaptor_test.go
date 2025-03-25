package logic

import (
	"bytes"
	"fmt"
)

type MockAdaptor struct {
	writer *bytes.Buffer
}

func NewMockAdaptor(buffer *bytes.Buffer) *MockAdaptor {
	return &MockAdaptor{
		writer: buffer,
	}
}

func (this *MockAdaptor) Write(value string) error {
	_, err := this.writer.Write([]byte(fmt.Sprintln(value)))
	return err
}
