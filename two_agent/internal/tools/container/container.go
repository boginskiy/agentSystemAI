package container

type Container struct {
	Command string
}

func NewContainer(command string) *Container {
	return &Container{
		Command: command,
	}
}

func (c *Container) CallCommand() string {
	return c.Command
}
