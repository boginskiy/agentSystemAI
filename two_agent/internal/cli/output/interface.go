package output

type Printer interface {
	Print(mess string) (int, error)
}

type Formater interface {
	LineMess() string
	WelcomeMess() string
}
