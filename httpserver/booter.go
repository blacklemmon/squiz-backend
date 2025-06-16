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

	_,err := tools.LoadConfig()
	//config,err := tools.LoadConfig()
	if (err!=nil){
		fmt.Println(err)
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