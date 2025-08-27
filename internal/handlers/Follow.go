package handlers

import (
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/google/uuid"
	"github.com/EchidnaTheG/Gator/internal/database"
	"time"
	"database/sql"
)

func HandlerFollow(s *commands.State,cmd commands.Command, User database.User) error{
	if len(cmd.Arguments) < 2{
		return fmt.Errorf("not enough arguments")
	}

	Feed, err := s.Db.LookUpFeedByURL(context.Background(), sql.NullString{String: cmd.Arguments[1], Valid: true})
	if err != nil{
		return err
	}
	feedid := Feed.ID
	userID := User.ID
	s.Db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Userid:    userID,
		Feedid: feedid,
	})
	fmt.Printf("Name of Feed Followed: %v\nBy Current User: %v\n",Feed.Name.String,User.Name)
	return nil
}