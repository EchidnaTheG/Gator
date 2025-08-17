package main 
import(
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/config"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"os"
	
)


func main(){
	// s is a struct with a pointer to the value that represents the state of the db
	var s commands.State
	value, err := config.Read()
	pToValue := &value
	if err != nil{
		fmt.Printf("SYSTEM: %v\n", err)
	}
	s.Ptoconfig=pToValue

	//registering all the commands with their handlers
	var Commands commands.Commands
	Commands.TypeOf = make(map[string]func(s *commands.State, cmd commands.Command) error)
	Commands.Register("login",commands.HandlerLogin)

	// collecting args
	args := os.Args
	if len(args) < 2 {
		fmt.Println("SYSTEM: Not enough arguments given.")
		return
	}
	 _,ok := Commands.TypeOf[args[1]] ; if !ok{
		fmt.Println("SYSTEM: Command Not Detected.")
		return
	}
	// excluding program name
	args = args[1:]
	cmd := commands.Command{Name: args[0],Arguments: args}
	err = Commands.Run(&s,cmd)
	if err != nil{
		fmt.Printf("SYSTEM: %v\n", err)
	}
	


}