package handlers

import(	
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/commands"
)

func HandlerUsers(s *commands.State,cmd commands.Command) error{
	users, err := s.Db.GetUsers(context.Background())
	if err != nil{
		return err
	}
	if len(users) == 0{
		fmt.Println("There are no users")
		return nil
	}
	for _, user := range users{
		if user.Name == s.Ptoconfig.Current_user_name{
			fmt.Printf("* %v (current)\n",user.Name)
			continue
		}
		fmt.Printf("* %v\n",user.Name)
	}
	return nil
}
