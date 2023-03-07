package main

import (
	"net/http"
	"log"
	"os"
	"handler/handler"
	"fmt"
	
	
	
)

func main(){
fmt.Println("server started")

l := log.New(os.Stdout,"product-api",log.LstdFlags)

hh := handler.NewHello(l)


   sm := http.NewServeMux()
   sm.Handle("/",hh)

	http.ListenAndServe(":9093",nil)
}