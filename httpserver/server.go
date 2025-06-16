package httpserver

import (

)

func RegisterRoutes(){
	//http.HandleFunc(path,function)
} 

func StartServer() error{
	RegisterRoutes()

	return nil
}