package commands

import (
	"fmt"

	"github.com/EchidnaTheG/Gator/internal/config"
)



type State struct{
	Ptoconfig *config.Config
}

type Command struct{
	Name string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error{
	if len(cmd.Arguments) == 0 || len(cmd.Arguments) == 1{
		return fmt.Errorf("error, no arguments given")
	}
	
	s.Ptoconfig.Current_user_name = cmd.Arguments[1]
	
	err := s.Ptoconfig.SetUser(s.Ptoconfig.Current_user_name)
	if err != nil{
		fmt.Printf("SYSTEM: %v\n",err)
		return err
	}
	fmt.Printf("User %v has been set!\n", s.Ptoconfig.Current_user_name)
	return nil
}
type Commands struct{
	TypeOf map[string]func(s *State, cmd Command) error
}

func (c *Commands) Run (s *State, cmd Command) error{
	if s != nil{
		return c.TypeOf[cmd.Name](s,cmd)
		
	}
	return fmt.Errorf("state is nil")
}

func (c *Commands) Register (name string, f func(*State, Command) error) error{
	c.TypeOf[name]=f
	return nil // gotta add stricter error handling
}