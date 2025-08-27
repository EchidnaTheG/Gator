package handlers
import(
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/commands"
)

func HandlerLogin(s *commands.State, cmd commands.Command) error{
	if len(cmd.Arguments) == 0 || len(cmd.Arguments) == 1{
		return fmt.Errorf("error, no arguments given")
	}
	
	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[1])
	if err != nil{
		return fmt.Errorf("user not found")
	}

	s.Ptoconfig.Current_user_name = cmd.Arguments[1]
	
	err = s.Ptoconfig.SetUser(s.Ptoconfig.Current_user_name)
	if err != nil{
		fmt.Printf("SYSTEM: %v\n",err)
		return err
	}
	fmt.Printf("User %v has been set!\n", s.Ptoconfig.Current_user_name)
	return nil
}
