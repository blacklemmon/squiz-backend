package routes

import (
	"net/http"
)

//func(http.ResponseWriter, *http.Request)
func Ping(w http.ResponseWriter, req *http.Request){
	if (req.Method == "GET") {
		w.Write([]byte("HHello!"))
	}
}