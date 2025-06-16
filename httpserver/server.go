package httpserver

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(){
	//http.HandleFunc(path,function)
} 

func StartServer() error{
	RegisterRoutes()

	return nil
}