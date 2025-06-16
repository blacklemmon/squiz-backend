package main

import(
	"fmt"
	"squizz/httpserver"
)
func main(){
	bootOK := httpserver.BootUp()
	
	if (bootOK!=nil) {
		fmt.Println("Boot got an error")
	}
	fmt.Println("Squizz Server is up")
}