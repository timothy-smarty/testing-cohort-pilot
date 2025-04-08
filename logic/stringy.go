package logic

type Stringy interface {
	Romans() string
}

type MyStringy struct {
}

func (this MyStringy) Romans() string {
	return "IVXLCDM"
}
