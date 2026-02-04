package models

import "fmt"

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Registered map[string]func(s *State, cmd Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	fun, ok := c.Registered[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	err := fun(s, cmd)
	return err
}

func (c *Commands) Register(name string, f func(s *State, cmd Command) error) {
	c.Registered[name] = f
}
