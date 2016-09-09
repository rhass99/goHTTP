package main

import (
	"net/http"
	"time"
)

//Custom Verbose Handler
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

//Functions as handlers
func timeHandler2(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("Function as handler time is: " + tm))
}

//Functions as handlers closure
func timeHandler3(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("Closure time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

//Functions as handlers closure2
func timeHandler4(format string) http.Handler {
	return http.handlerfunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("Closure time is: " + tm))
	})
}

func main() {
	mux := http.NewServeMux()

	//Built in Handler
	rh1 := http.RedirectHandler("http://www.google.com", 307)
	mux.Handle("/built", rh1)

	//Custom Handler
	rh2 := &timeHandler{format: time.RFC1123}
	mux.Handle("/custom", rh2)

	//Function as handler
	rh3 := http.HandlerFunc(timeHandler2)
	mux.Handle("/func", rh3)

	//Shortcut Function as handler
	mux.HandleFunc("/shortfunc", timeHandler2)

	//Closure Function as handler
	mux.Handle("/closure", timeHandler3("RFC1123"))

	//Server Run
	http.ListenAndServe(":3000", mux)

}
