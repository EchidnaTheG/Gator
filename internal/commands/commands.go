package commands

import (
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/config"
	"github.com/EchidnaTheG/Gator/internal/database"
)



type State struct{
	Ptoconfig *config.Config
	Db *database.Queries
}

type Command struct{
	Name string
	Arguments []string
}



func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) (func(*State, Command) error){
	return func(s *State,cmd Command) error{
		current_user := s.Ptoconfig.Current_user_name
		User, err := s.Db.GetUser(context.Background(),current_user)
		if err != nil{
			return err
		}
		return handler(s,cmd,User)
	}
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