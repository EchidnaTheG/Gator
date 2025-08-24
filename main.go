package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/EchidnaTheG/Gator/internal/commands"
	"github.com/EchidnaTheG/Gator/internal/config"
	"github.com/EchidnaTheG/Gator/internal/database"
	_ "github.com/lib/pq"
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
	dbURL := s.Ptoconfig.Db_url
	db, err := sql.Open("postgres", dbURL)
	if err != nil{
		fmt.Printf("SYSTEM: Error Opening Up Postgres, %v\n",err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s.Db = dbQueries


	//registering all the commands with their handlers
	var Commands commands.Commands
	Commands.TypeOf = make(map[string]func(s *commands.State, cmd commands.Command) error)
	Commands.Register("login",commands.HandlerLogin)
	Commands.Register("register",commands.HandlerRegister)
	Commands.Register("reset",commands.HandlerReset)
	Commands.Register("users", commands.HandlerUsers)
	// collecting args
	args := os.Args
	if len(args) < 2 {
		fmt.Println("SYSTEM: Not enough arguments given.")
		os.Exit(1)
	}
	 _,ok := Commands.TypeOf[args[1]] ; if !ok{
		fmt.Println("SYSTEM: Command Not Detected.")
		os.Exit(1)
	}
	// excluding program name
	args = args[1:]
	cmd := commands.Command{Name: args[0],Arguments: args}
	err = Commands.Run(&s,cmd)
	if err != nil{
		fmt.Printf("SYSTEM: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
	


}