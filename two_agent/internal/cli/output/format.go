package output

import (
	"fmt"
	"os/user"
)

const DefaultUser = "user"

type Format struct {
}

func NewFormat() *Format {
	return &Format{}
}

func (f *Format) LineMess(mess string) string {
	return fmt.Sprintf("%s\n", mess)
}

func (f *Format) LineMessWithErr(mess string, err error) string {
	return fmt.Sprintf("%s Error: %s\n", mess, err.Error())
}

func (f *Format) WelcomeMess() string {
	user, err := user.Current()
	if err != nil {
		return fmt.Sprintf("Hello, %s\n", DefaultUser)
	}
	return fmt.Sprintf("Hello, %s\n", user.Username)
}

func (f *Format) EnterCommand(mess string) string {
	return fmt.Sprintf("%s:\n", mess)
}
