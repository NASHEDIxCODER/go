package main

import (
	"fmt"
	"net/http"
)

//writting basic HTTP server go using net/http



func hello (w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w, "hello\n")
	//a fundamental concept in net/http sever is handlers.
	// ahandler is an object implementing the http.Handler interface.
	// a common way to writte a handler is by using the http.HandlerFunc adapter on the functions with the appropriate signature.
	//func serviung as handler take a http.REsponseWritter and a http.Request as arguments.
	// the response writter is used to fill in the HTTP response. Here our simple response is just "hello\n".

}

//this handler does something a little more shophisticated by reading all the http request headers and echoing them into the response body.
 func headers(w http.ResponseWriter, req *http.Request){
	for name, headers:= range req.Header{
		for _, h := range headers {
			fmt.Fprintln(w, "%v: %v\n", name, h)
		}
	}
 }
 func main(){
	//we register our handler on server routes using the http.HandeFunc convenience funtion.
	//it sets up the default router in the net/http package and takes a function as an argument.

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// finally we call the ListenAndServe with the port and a handler. nil tells it to use the default router we'have just set up.

	http.ListenAndServe(":8090",nil)
 }