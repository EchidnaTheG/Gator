package main 
import(
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/config"
)


func main(){
	value, err := config.Read()
	if err != nil{
		fmt.Printf("Error: %v", err)
	}
	err = value.SetUser("Eli")
	if err != nil{
		fmt.Printf("Error: %v", err)
	}
	value, err = config.Read()
	if err != nil{
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(value)
}