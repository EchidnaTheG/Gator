package handlers

import(
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/EchidnaTheG/Gator/internal/rss"
)


func HandlerAgg(s *commands.State,cmd commands.Command) error{
	if len(cmd.Arguments) <= 1 {
		return fmt.Errorf("not enough arguments")
	}
    RSSFeed, err :=rss.FetchFeed(context.Background(),cmd.Arguments[1])
	if err != nil{
		return nil
	}
	fmt.Printf("%v\n", *RSSFeed)
	return nil
}