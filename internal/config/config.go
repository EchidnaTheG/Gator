package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}
const jsonName = "/.gatorconfig.json"

func (c *Config) SetUser(username string) error{
	userHome, err := os.UserHomeDir()
	if err != nil{
		return err
	}
	userHome += jsonName
	file, err :=os.OpenFile(userHome,os.O_RDWR,0644)
	if err != nil{
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil{
		return err
	}

	file.Close()
	var payload Config
	json.Unmarshal(data,&payload)
	payload.Current_user_name = username
	c.Current_user_name= username
	value, err := json.Marshal(&payload)
	if err != nil{
		return err
	}
	err = os.WriteFile(userHome,value,0644)
	if err != nil{
		return err
	}
	return nil
}

func Read() (Data Config, error error){
	userHome, err := os.UserHomeDir()
	if err != nil{
		return Config{}, err
	}
	userHome += jsonName
	data, err :=os.ReadFile(userHome)

	if err != nil{
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil{
		return Config{}, err
	}
	return config,nil
}
