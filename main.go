package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hello")
	}
func main() {
	http.HandleFunc("/",handler )
	err:= http.ListenAndServe(":8080",nil )
	if err !=nil {
		log.Fatal("ListenAdServer", err)
	}
	}
