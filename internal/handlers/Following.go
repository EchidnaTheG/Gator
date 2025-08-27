package handlers

import (
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/database"
	"github.com/EchidnaTheG/Gator/internal/commands"
)


func HandlerFollowing(s *commands.State,cmd commands.Command, User database.User) error{
	userID := User.ID
	FeedFollows, err :=s.Db.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil{
		return err
	}

	for _, FeedFollow := range FeedFollows{
		Feed, err := s.Db.LookUpFeedByID(context.Background(),FeedFollow.Feedid)
		if err != nil{
			return err
		}
		fmt.Printf("Following: %v\n",Feed.Name.String )
	}
	return nil
}
