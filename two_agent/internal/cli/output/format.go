package output

import (
	"fmt"
	"os/user"
)

const DefaultUser = "user"

type Format struct {
}

func (f *Format) LineMess() string {
	return "%s\n"
}

func (f *Format) WelcomeMess() string {
	user, err := user.Current()
	if err != nil {
		return fmt.Sprintf("Hello, %s\n", DefaultUser)
	}
	return fmt.Sprintf("Hello, %s\n", user.Username)
}
