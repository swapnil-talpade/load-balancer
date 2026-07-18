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
	req,err:=http.NewRequest(r.Method, 
		"http://localhost:9001"+r.URL.RequestURI(),
		r.Body)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	// forward request headers
	for key,values:=range r.Header{
		for _,value:=range values{
			req.Header.Add(key,value)
		}
	}

	// send requests to backend server
	resp,err:=http.DefaultClient.Do(req)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	// forward response headers
	for key,values:=range resp.Header{
		for _,value:=range values{
			w.Header().Add(key,value)
		}
	}

	// forward status code
	w.WriteHeader(resp.StatusCode)

	// forward response body
	io.Copy(w, resp.Body)

}



