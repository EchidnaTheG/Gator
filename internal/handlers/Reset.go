package handlers
import(
	"context"
	"github.com/EchidnaTheG/Gator/internal/commands"
)

func HandlerReset(s *commands.State,cmd commands.Command) error{
	err := s.Db.Reset(context.Background())
	if err != nil{
		return err
	}
	return nil
}
