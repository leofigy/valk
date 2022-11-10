package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/state", echo)
	router.HandleFunc("/topology/{region}", echo)

	log.Fatal(http.ListenAndServe("127.0.0.1:7001", router))

}

func echo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Mux Vars %+v\n", vars)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	log.Println(r.Body)

	defer r.Body.Close()
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(payload))
}
