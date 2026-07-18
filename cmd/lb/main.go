package main

import (
	"io"
	"log"
	"net/http"
)



func main(){
	http.HandleFunc("/",proxyHandler)

	log.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe(":8080",nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request){
	resp,err:=http.Get("http://localhost:9001"+r.RequestURI)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadGateway)
	}

	defer resp.Body.Close()

	io.Copy(w,resp.Body)

}



