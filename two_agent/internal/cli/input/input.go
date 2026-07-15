package input

import (
	"bufio"
	"io"
)

type Input struct {
	buff *bufio.Scanner
}

func NewScanner(reader io.Reader) *Input {
	return &Input{
		buff: bufio.NewScanner(reader),
	}
}

func (s *Input) scan() error {
	s.buff.Scan()
	if s.buff.Err() != nil {
		return s.buff.Err()
	}
	return nil
}

func (s *Input) takeText() (string, error) {
	text := s.buff.Text()
	if s.buff.Err() != nil {
		return "", s.buff.Err()
	}
	return text, nil
}

func (s *Input) Scan() (string, error) {
	err := s.scan()
	if err != nil {
		return "", err
	}
	return s.takeText()
}
