package logic

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestConverterFixture(t *testing.T) {
	gunit.Run(new(ConverterFixture), t)
}

type ConverterFixture struct {
	*gunit.Fixture
}

func (this *ConverterFixture) TestConvert() {
	actual := Convert("28")
	expected := "XXVIII"
	this.So(actual, should.Equal, expected)
}
