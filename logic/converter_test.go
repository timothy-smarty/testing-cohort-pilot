package logic

import (
	"testing"

	"github.com/smarty/assertions/should"
)

func TestConverter(t *testing.T) {
	testTable := map[string]struct {
		Input          string
		ExpectedOutput string
		ExpectedError  error
	}{
		"Convert perfect number":   {Input: "28", ExpectedOutput: "XXVIII", ExpectedError: nil},
		"Convert non-numeric":      {Input: "hello", ExpectedOutput: "", ExpectedError: ErrorNotANumber},
		"Convert max-value":        {Input: "3999", ExpectedOutput: "MMMCMXCIX", ExpectedError: nil},
		"Convert beyond-max-value": {Input: "4000", ExpectedOutput: "", ExpectedError: ErrorNumberTooBig},
		"Convert min-value":        {Input: "1", ExpectedOutput: "I", ExpectedError: nil},
		"Convert too-small":        {Input: "0", ExpectedOutput: "", ExpectedError: ErrorNumberTooSmall},
	}

	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("friggin panicked dude: %s", err.(error).Error())
				}
			}()

			actual, err := Convert(testCase.Input)
			should.So(t, actual, should.Equal, testCase.ExpectedOutput)
			if testCase.ExpectedError == nil {
				should.So(t, err, should.BeNil)
			} else {
				should.So(t, err, should.Wrap, testCase.ExpectedError)
			}
		})
	}
}

func TestSelectNumeral(t *testing.T) {
	testTable := map[string]struct {
		Index          int
		ExpectedOutput byte
	}{
		"selectNumeral -1": {Index: -1, ExpectedOutput: ' '},
		"selectNumeral 0":  {Index: 0, ExpectedOutput: 'I'},
		"selectNumeral 6":  {Index: 6, ExpectedOutput: 'M'},
		"selectNumeral 7":  {Index: 7, ExpectedOutput: ' '},
	}

	loadRomans()
	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			actual := selectNumeral(romans, testCase.Index)
			should.So(t, actual, should.Equal, testCase.ExpectedOutput)
		})
	}
}
