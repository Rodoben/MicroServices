package main

import(
	
	"net/http"
	"fmt"
	"log"
	"io"
)

/*

ListenAndServe:

ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

The handler is typically nil, in which case the DefaultServeMux is used.

ListenAndServe always returns a non-nil error. 

ListenAndServe is used to initialize or setup a  server

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
func ListenAndServe(addr string, handler Handler) error
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

ioutil/os
func ReadAll(r io.Reader) ([]byte, error)

func Error(w ResponseWriter, error string, code int)

*/


func main() {
	fmt.Println("Hi")


http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	log.Println("Hello prinitng log")
	rs,err:=io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", rs)

	
})

http.HandleFunc("/hi", func(w http.ResponseWriter,r *http.Request){
	log.Println("Hello prinitng hi log")
	rs,err:=io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", rs,nil)

})

http.HandleFunc("/bye", func(http.ResponseWriter,*http.Request){
	log.Println("Hello prinitng bye log")

})
        http.ListenAndServe(":8087",nil)
	
}