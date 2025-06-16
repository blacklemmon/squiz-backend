package httpserver

import (
	"fmt"
)

func LoadConfig() error{
	fmt.Println("Loading configuration")
	/*This function should read a yaml file and store the
		server settings in a type struct
	*/
	return nil
}

func ConnectToDB() error{
	fmt.Println("Connecting to database")

	return nil
}

func BootUp() error{
	fmt.Println("Server booting...")

	err := LoadConfig()
	if (err!=nil){
		return err
	}

	err = ConnectToDB()
	if (err!=nil){
		return err
	}

	err = StartServer()
	if (err!=nil){
		return err
	}

	fmt.Println("Server boot finished")
	return nil
}