package handlers
import(
	"context"
	"fmt"
	"time"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/EchidnaTheG/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister (s *commands.State,cmd commands.Command) error{
	if len(cmd.Arguments) == 0 || len(cmd.Arguments) == 1{
		return fmt.Errorf("error, no name given")
	}
	
	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[1])
	if err == nil{
		return fmt.Errorf("user error, duplicate name detected and not saved")
	}
	
	
	user , err :=s.Db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now() ,UpdatedAt: time.Now(),Name: cmd.Arguments[1]})
	if err != nil{
		return err
	}
	s.Ptoconfig.Current_user_name= user.Name
	s.Ptoconfig.SetUser(user.Name)
	fmt.Printf("SYSTEM: USER %v WAS CREATED\n", user.Name)
	return nil
}