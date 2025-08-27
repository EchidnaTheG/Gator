package handlers

import (
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/EchidnaTheG/Gator/internal/database"
	"database/sql"
)

func HandlerUnfollow(s *commands.State,cmd commands.Command, User database.User) error{
	if len(cmd.Arguments) < 2{
		return fmt.Errorf("not enough arguments")
	}

	Feed, err := s.Db.LookUpFeedByURL(context.Background(), sql.NullString{String: cmd.Arguments[1], Valid: true})
	if err != nil{
		return err
	}
	feedid := Feed.ID
	userID := User.ID
	err = s.Db.UnfollowFeedByUser(context.Background(),database.UnfollowFeedByUserParams{Userid: userID,Feedid: feedid})
	if err != nil{
		return err
	}
	fmt.Printf("Feed %v Has been Unfollowed By %v!\n",Feed.Name.String,User.Name)
	return nil
}