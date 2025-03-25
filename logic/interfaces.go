package logic

// WriterPort defines the contract for output adaptors.
type WriterPort interface {
	// Write takes a string, appends a newline, and writes to some io.Writer.
	//
	// Parameters:
	//   - value is the result string. Remember that this value will be appended
	//     with a newline.
	//
	// Returns:
	//   - err matches the errors possible for whatever io.Writer is being used.
	//     for more details, see the specific adaptor.
	Write(value string) (err error)
}
