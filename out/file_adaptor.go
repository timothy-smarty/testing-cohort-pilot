package out

import (
	"fmt"
	"io"
)

// FileAdaptor fulfills the [logic.WriterPort] interface by writing to a passed
// io.Writer.
type FileAdaptor struct {
	writer io.Writer
}

// NewFileAdaptor creates and returns a pointer to a new file adaptor.
func NewFileAdaptor(writer io.Writer) *FileAdaptor {
	return &FileAdaptor{
		writer: writer,
	}
}

// Write writes to the io.Writer that was originally passed in.
//
// Parameters:
//   - value is the output string to write to the io.Writer. A newline will be
//     added to the end of the string.
//
// Returns:
//   - err is the same error expected for io.Writer.
func (this *FileAdaptor) Write(value string) error {
	_, err := this.writer.Write([]byte(fmt.Sprintln(value)))
	return err
}
