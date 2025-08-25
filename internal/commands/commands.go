package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/EchidnaTheG/Gator/internal/config"
	"github.com/EchidnaTheG/Gator/internal/database"
	"github.com/EchidnaTheG/Gator/internal/rss"
	"github.com/google/uuid"
)



type State struct{
	Ptoconfig *config.Config
	Db *database.Queries
}

type Command struct{
	Name string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error{
	if len(cmd.Arguments) == 0 || len(cmd.Arguments) == 1{
		return fmt.Errorf("error, no arguments given")
	}
	
	_, err := s.Db.GetUser(context.Background(), cmd.Arguments[1])
	if err != nil{
		return fmt.Errorf("user not found")
	}

	s.Ptoconfig.Current_user_name = cmd.Arguments[1]
	
	err = s.Ptoconfig.SetUser(s.Ptoconfig.Current_user_name)
	if err != nil{
		fmt.Printf("SYSTEM: %v\n",err)
		return err
	}
	fmt.Printf("User %v has been set!\n", s.Ptoconfig.Current_user_name)
	return nil
}

func HandlerRegister (s *State,cmd Command) error{
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


func HandlerReset(s *State,cmd Command) error{
	err := s.Db.Reset(context.Background())
	if err != nil{
		return err
	}
	return nil
}

func HandlerUsers(s *State,cmd Command) error{
	users, err := s.Db.GetUsers(context.Background())
	if err != nil{
		return err
	}
	if len(users) == 0{
		fmt.Println("There are no users")
		return nil
	}
	for _, user := range users{
		if user.Name == s.Ptoconfig.Current_user_name{
			fmt.Printf("* %v (current)\n",user.Name)
			continue
		}
		fmt.Printf("* %v\n",user.Name)
	}
	return nil
}

func HandlerAgg(s *State,cmd Command) error{
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

func HandlerAddFeed(s *State,cmd Command) error{
	if len(cmd.Arguments) <= 2 {
		return fmt.Errorf("not enough arguments")
	}
	current_user := s.Ptoconfig.Current_user_name
	User, err := s.Db.GetUser(context.Background(),current_user)
	if err != nil{
		return err
	}
	userID := User.ID
	DBFeed,err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: cmd.Arguments[1], Valid: true},
		Url:       sql.NullString{String: cmd.Arguments[2], Valid: true},
		Userid:    userID,
	})
	if err != nil {
		return err
	}
	s.Db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Userid:    userID,
		Feedid: DBFeed.ID,
	})
	fmt.Printf("%v\n",DBFeed)
	return nil
}

func helperFunctionForFeeds(s *State, feedUserID  uuid.UUID ) (string, error){
	User, err := s.Db.GetUserByID(context.Background(),feedUserID)
	if err != nil{
		return "", err
	}
	return User.Name, nil
}

func HandlerFeeds(s *State,cmd Command) error{
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

func HandlerFollow(s *State,cmd Command) error{
	if len(cmd.Arguments) < 2{
		return fmt.Errorf("not enough arguments")
	}

	Feed, err := s.Db.LookUpFeedByURL(context.Background(), sql.NullString{String: cmd.Arguments[1], Valid: true})
	if err != nil{
		return err
	}
	feedid := Feed.ID
	User, err := s.Db.GetUser(context.Background(), s.Ptoconfig.Current_user_name)
	if err != nil{
		return err
	}
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

func HandlerFollowing(s *State,cmd Command) error{
	User, err := s.Db.GetUser(context.Background(), s.Ptoconfig.Current_user_name)
	if err != nil{
		return err
	}
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


type Commands struct{
	TypeOf map[string]func(s *State, cmd Command) error
}

func (c *Commands) Run (s *State, cmd Command) error{
	if s != nil{
		return c.TypeOf[cmd.Name](s,cmd)
		
	}
	return fmt.Errorf("state is nil")
}

func (c *Commands) Register (name string, f func(*State, Command) error) error{
	c.TypeOf[name]=f
	return nil // gotta add stricter error handling
}