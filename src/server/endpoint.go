package main

import (
	"fmt"
	"net"
	"net/http"
)

type EndPoint struct {
	id       string
	addr     string
	port     int
	handler  http.Handler
	listener net.Listener
}

func NewEndPoint(id string, addr string, port int, handler http.Handler, listener net.Listener) *EndPoint {
	ep := new(EndPoint)
	ep.id = id
	ep.addr = addr
	ep.port = port
	ep.handler = handler
	ep.listener = listener
	return ep
}

func (this *EndPoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is instance of Endpoint")
	//w.Write([]byte("The counter  is: " + this.count))
}

type Step struct {
	url      string
	response string
}

type Scenario struct {
	steps []Step
}
