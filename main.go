package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

var (
	handlers = make(map[string]handler)
)

var (
	HOME_PATH string
)

func init() {
	log.SetFlags(log.Lshortfile)

	//register handlers
	handlers["GET:/"] = GetHomeHandler
	handlers["GET:/json"] = GetJsonHandler
}

func main() {
	http.Handle("/public/", http.FileServer(http.Dir("")))
	http.HandleFunc("/", IndexHandler)

	if err := http.ListenAndServe(":9898", nil); nil != err {
		panic(err)
	}
}

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	key := fmt.Sprintf("%s:%s", r.Method, r.URL.Path)
	log.Println(key)
	handler, ok := handlers[key]
	if !ok {
		http.NotFound(rw, r)
		return
	}

	handler(rw, r)
}

func GetHomeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("homepage"))
}

func GetJsonHandler(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		Name   string
		Gender string
		Age    int
	}{
		"elvindu",
		"male",
		30,
	}

	bin, err := json.Marshal(data)
	if nil != err {
		fmt.Fprintf(rw, err.Error())
		return
	}

	fmt.Fprintf(rw, "%s", string(bin))
}
