package main

import (

    "fmt"
    "io"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "net/http"	
)

type input struct{
	Name string
}

type Output struct{
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

}

func greeting(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	decoder := json.NewDecoder(req.Body)
	
    	var t input   
    	err := decoder.Decode(&t)
   	 if err != nil && err == io.EOF {
		http.Error(rw, "Request body is missing", 400)
		return
    	} else if err != nil {
		http.Error(rw, "Internal Server Error Occured", 500)
		return
	}
	var outputStr = "Hello, "+t.Name+"!"	
	outputData := &Output{Greeting:outputStr}
	b, err := json.Marshal(outputData)
    	if err != nil {
        	http.Error(rw, "Internal Server Error Occured", 500)
		return
    	} else {
    		fmt.Fprintf(rw, string(b[:]))
	}

}


func main() {

    mux := httprouter.New() 
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", greeting)
    server := http.Server{

            Addr:        "0.0.0.0:8880",

            Handler: mux,

    }

    server.ListenAndServe()

}
