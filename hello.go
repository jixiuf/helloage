package hello

import (
	"appengine"
	"appengine/urlfetch"
	"appengine/user"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/helloworld", handlerHelloWorld)
	http.HandleFunc("/fetch", handlerFetchUrl)
	http.HandleFunc("/user", handlerUser)

}
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<a href='helloworld'>helloworld</a><br/>")
	fmt.Fprint(w, "<a href='fetch'>fetch google url</a><br/>")
	fmt.Fprint(w, "<a href='user'>say hello to appengine.user</a>")
}

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		fmt.Fprint(w, "user not login,try login... ")
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			fmt.Fprint(w, "user not login,try login error ")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "user not login,not found ")
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello, User:%v!", u)
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
