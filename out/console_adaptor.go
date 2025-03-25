package out

import (
	"os"
)

// ConsoleAdaptor fulfills the [logic.WriterPort] interface by writing to the
// standard-out.
type ConsoleAdaptor struct {
	FileAdaptor
}

// NewConsoleAdaptor creates and returns a pointer to a new console adaptor.
func NewConsoleAdaptor() *ConsoleAdaptor {
	return &ConsoleAdaptor{
		*NewFileAdaptor(os.Stdout),
	}
}

// Write writes to the standard-out file.
//
// Parameters:
//   - value is the output string to write to the console. A newline will be
//     added to the end of the string.
//
// Returns:
//   - err will be nil.
func (this *ConsoleAdaptor) Write(value string) (err error) {
	return this.FileAdaptor.Write(value)
}
