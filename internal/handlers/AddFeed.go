package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/EchidnaTheG/Gator/internal/database"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *commands.State,cmd commands.Command, User database.User) error{
	if len(cmd.Arguments) <= 2 {
		return fmt.Errorf("not enough arguments")
	}
	DBFeed,err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: cmd.Arguments[1], Valid: true},
		Url:       sql.NullString{String: cmd.Arguments[2], Valid: true},
		Userid:    User.ID,
	})
	if err != nil {
		return err
	}
	s.Db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Userid:    User.ID,
		Feedid: DBFeed.ID,
	})
	fmt.Printf("%v\n",DBFeed)
	return nil
}