package main

import "fmt"
import "net/http"
import "html"
import "io/ioutil"
import "log"
import "encoding/json"


type Message struct {
    Name string
    Body string
    Time int64
}

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
			case "POST":
				fmt.Fprintf(w, "Toma teu post então, %q", html.EscapeString(r.URL.Path))
			default:
				fmt.Fprintf(w, "Só suportamos POST e GET aqui")
			}
	})
	http.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://www.mocky.io/v2/5dbb6b30310000b5084c0b8a")
		if err != nil {
			log.Panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, string(body[:]))
	})
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		m := Message{"Alice", "Hello", 1294706395881547000}
		b, err := json.Marshal(m)
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprintf(w, string(b[:]))
	})
	http.ListenAndServe(":8080", nil)
}
