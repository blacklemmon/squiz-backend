package httpserver

import (
	"fmt"
	"squizz/tools"
)

func ConnectToDB() error{
	fmt.Println("Connecting to database")

	return nil
}

func BootUp() error{
	fmt.Println("Server booting...")

	config,err := tools.LoadConfig()
	//_,err := tools.LoadConfig()
	if (err!=nil){
		fmt.Println(err)
		return err
	}

	err = ConnectToDB()
	if (err!=nil){
		fmt.Println(err)
		return err
	}

	err = StartServer(config.Server.Port)
	if (err!=nil){
		fmt.Println(err)
		return err
	}

	fmt.Println("Server boot finished")
	return nil
}