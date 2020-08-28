package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var rspID = 0
var rnd *rand.Rand

const DefaultFailRate = 10    // in percent
const DefaultFailTimeout = 10 // in milliseconds
const DefaultServerPort = "8080"

var ss serverSettings // hate this!

type serverSettings struct {
	failureRate    int
	failureTimeout int
	serverPort     string
	serverName     string
}

func handler(w http.ResponseWriter, r *http.Request) {
	rspID++
	if rnd.Intn(100) > ss.failureRate {
		fmt.Fprintf(w, "Hello! You've requested %s!\n", r.URL.Path)
		fmt.Fprintf(w, "Server: %s, Port: %s, Respose Id: %d!\n", ss.serverName, ss.serverPort[1:], rspID)
	} else {
		time.Sleep(time.Duration(ss.failureTimeout) * time.Millisecond)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Simulated a server failure!\n"))
		fmt.Fprintf(w, "Server: %s, Port: %s, Respose Id: %d!\n", ss.serverName, ss.serverPort[1:], rspID)
	}
}

func parseArgs(ss *serverSettings) {
	fR := flag.Int("failure-rate", DefaultFailRate, "the rate of failures to simulate (in percent)")
	fTO := flag.Int("failure-timeout", DefaultFailTimeout, "the timeout in case of failure (in msecs)")
	sP := flag.String("server-port", DefaultServerPort, "the server port to listen for HTTP requests")
	flag.Parse()

	ss.failureRate = *fR
	ss.failureTimeout = *fTO
	ss.serverPort = ":" + *sP
	ss.serverName, _ = os.Hostname()

}

func main() {
	// parse the args
	parseArgs(&ss)

	// for random generation
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	// start the server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(ss.serverPort, nil))
}
