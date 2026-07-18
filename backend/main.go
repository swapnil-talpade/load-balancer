package main

import (
	"fmt"
	"net/http"
)


func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello World!\n")
	})

	fmt.Println("Backend running on port :9001")

	http.ListenAndServe(":9001",nil)

}