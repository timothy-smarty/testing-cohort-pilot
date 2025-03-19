package logic

import "fmt"

var (
	ErrorNumberTooBig   = fmt.Errorf("the input value is too big, should be no more than [%d]", maxInput)
	ErrorNumberTooSmall = fmt.Errorf("the input value is too small, should be at least [%d]", minInput)
	ErrorNotANumber     = fmt.Errorf("Input is not an integer value")
)
