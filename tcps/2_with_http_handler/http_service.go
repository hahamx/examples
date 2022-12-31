package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	httpPort = ":3010"
	muxPort  = ":3020"
	wg       = sync.WaitGroup{}
)

// ////////////动物园一  http server
type zoomKeeper struct{}

func (h zoomKeeper) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/":
		io.WriteString(res, `welcome to visit owl, elephants, tiger!`)
	case "/owls":
		io.WriteString(res, `miao,miao...`)
	case "/elephants":
		io.WriteString(res, `ang,ang...`)
	case "/tiger":
		io.WriteString(res, `aoo,aoo...`)
	}
}

func httpServer() {
	var zk zoomKeeper
	fmt.Printf("zoomkeeper 动物园一:%v\n", httpPort)
	http.ListenAndServe(httpPort, zk)
	wg.Done()
}

// ////////////动物园二  Mux server
type zooms struct{}

func (h zooms) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `welcome Big Zoom Here have  owl, elephants, tiger!`)
}

type zoomOwl struct{}

func (h zoomOwl) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `wu,wu,wu,wu,...`)
}

type zoomElephants struct{}

func (h zoomElephants) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `ang,ang,ang,ang,...`)
}

type zoomFood struct{}

func (h zoomFood) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `corn,apple,rice,grass...`)
}

type zoomTiger struct{}

func (h zoomTiger) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `aoo,aoo,aoo,aoo,...`)
}

func muxHttpServer() {
	var manager zooms
	var owl zoomOwl
	var ele zoomElephants
	var eleFood zoomFood
	var tiger zoomTiger

	router := http.NewServeMux()
	router.Handle("/", manager)
	router.Handle("/owl", owl)
	router.Handle("/ele/", ele)
	router.Handle("/ele/food", eleFood)
	router.Handle("/tiger", tiger)

	// router.Handle("/codesearch", search)
	// router.Handle("codesearch.google.com",)
	fmt.Printf("zoomkeeper 动物园二:%v\n", muxPort)

	http.ListenAndServe(muxPort, router)
	wg.Done()
}

func main() {

	wg.Add(2)
	go httpServer()

	go muxHttpServer()
	wg.Wait()
}
