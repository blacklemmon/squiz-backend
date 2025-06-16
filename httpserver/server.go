package httpserver

import (
	"fmt"
	"net/http"
	"squizz/routes"
)

func RegisterRoutes(){
	http.HandleFunc("/", routes.Ping)
} 

func StartServer(port int) error{
	RegisterRoutes()

	address := fmt.Sprintf(":%d",port)

	httpErr := http.ListenAndServe(address,nil)
	return httpErr
}