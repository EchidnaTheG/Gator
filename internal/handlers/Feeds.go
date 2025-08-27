package handlers

import (
	"context"
	"fmt"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/google/uuid"
)

func helperFunctionForFeeds(s *commands.State, feedUserID  uuid.UUID ) (string, error){
	User, err := s.Db.GetUserByID(context.Background(),feedUserID)
	if err != nil{
		return "", err
	}
	return User.Name, nil
}

func HandlerFeeds(s *commands.State,cmd commands.Command) error{
	Feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil{
		return err
	}
	for i, Feed := range Feeds{
		name, err := helperFunctionForFeeds(s,Feed.Userid)
		if err != nil{
			return err
		}
		fmt.Printf("%v.\n Name: %v\n URL: %v\n Created By: %v\n", i+1, Feed.Name.String, Feed.Url.String, name)

	}
	return nil

}