package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func ListenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func main() {
	handler := newHandler()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go ListenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)
}
