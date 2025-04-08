package logic

import (
	"bytes"
	"testing"

	"bitbucket.org/smartystreets/injector"
	"github.com/smarty/assertions/should"
)

func TestConverter(t *testing.T) {
	inj := injector.New()
	injector.RegisterSingleton[Stringy](inj, func() Stringy { return &MyStringy{} })
	injector.RegisterTransient[*Converter](inj, NewConverter)
	injector.Verify(inj)

	converter := injector.Get[*Converter](inj)

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

			buffer := new(bytes.Buffer)
			actual, err := converter.Convert(testCase.Input, NewMockAdaptor(buffer))
			should.So(t, actual, should.Equal, testCase.ExpectedOutput)
			if testCase.ExpectedError == nil {
				should.So(t, err, should.BeNil)
			} else {
				should.So(t, err, should.Wrap, testCase.ExpectedError)
			}
		})
	}
}

func TestConverterPackage(t *testing.T) {
	converter := NewConverter(&MyStringy{})
	testTable := map[string]struct {
		Input          string
		ExpectedOutput string
		ExpectedError  error
	}{
		"Convert perfect number":   {Input: "28", ExpectedOutput: "XXVIII\n", ExpectedError: nil},
		"Convert non-numeric":      {Input: "hello", ExpectedOutput: "", ExpectedError: ErrorNotANumber},
		"Convert max-value":        {Input: "3999", ExpectedOutput: "MMMCMXCIX\n", ExpectedError: nil},
		"Convert beyond-max-value": {Input: "4000", ExpectedOutput: "", ExpectedError: ErrorNumberTooBig},
		"Convert min-value":        {Input: "1", ExpectedOutput: "I\n", ExpectedError: nil},
		"Convert too-small":        {Input: "0", ExpectedOutput: "", ExpectedError: ErrorNumberTooSmall},
	}

	for name, testCase := range testTable {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("friggin panicked dude: %s", err.(error).Error())
				}
			}()

			buffer := new(bytes.Buffer)
			_, err := converter.Convert(testCase.Input, NewMockAdaptor(buffer))
			should.So(t, buffer.String(), should.Equal, testCase.ExpectedOutput)
			if testCase.ExpectedError == nil {
				should.So(t, err, should.BeNil)
			} else {
				should.So(t, err, should.Wrap, testCase.ExpectedError)
			}
		})
	}
}
