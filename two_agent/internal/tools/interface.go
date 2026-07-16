package tools

type Toolmaker interface {
	CallCommand() string
	Do(conditions []string) error
}

type Creater interface {
}
