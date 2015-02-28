package hello

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/helloworld", handlerHelloWorld)
	http.HandleFunc("/fetch", handlerFetchUrl)

}

func handlerFetchUrl(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://www.google.com/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<a href='helloworld'>helloworld</a><br/>")
	fmt.Fprint(w, "<a href='fetch'>fetch google url</a>")
}
