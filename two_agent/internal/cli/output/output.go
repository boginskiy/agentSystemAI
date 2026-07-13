package output

import (
	"fmt"
	"os"
)

type Output struct {
	Formatter Format
}

func NewOutput(format Format) *Output {
	return &Output{
		Formatter: format,
	}
}

func (o *Output) Print(mess string) (int, error) {
	n, err := fmt.Fprint(os.Stdout, mess)
	if err != nil {
		return 0, err
	}
	return n, nil
}
