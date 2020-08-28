// +build ignore

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var rsp_id = 0
var rnd *rand.Rand
var DEFAULT_FAIL_RATE = 10    // percent
var DEFAULT_FAIL_TIMEOUT = 10 // milliseconds
var DEFAULT_SERVER_PORT = 8080
var ss server_settings // hate this!

type server_settings struct {
	failure_rate    int
	failure_timeout int
	server_port     string
}

func handler(w http.ResponseWriter, r *http.Request) {
	rsp_id++
	if rnd.Intn(100) > ss.failure_rate {
		fmt.Fprintf(w, "Hello! You've requested %s!\n", r.URL.Path)
		fmt.Fprintf(w, "Respose Id: %d!\n", rsp_id)
	} else {
		time.Sleep(time.Duration(ss.failure_timeout) * time.Millisecond)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Simulated a server failure!\n"))
		fmt.Fprintf(w, "Respose Id: %d!\n", rsp_id)
	}
}

func parseArgs(ss *server_settings) {
	f_r := flag.Int("failure-rate", 10, "the rate of failures to simulate (in percent)")
	f_to := flag.Int("failure-timeout", 10, "the timeout in case of failure (in msecs)")
	s_p := flag.String("server-port", "8080", "the server port to listen for HTTP requests")
	flag.Parse()

	ss.failure_rate = *f_r
	ss.failure_timeout = *f_to
	ss.server_port = ":" + *s_p

	fmt.Printf("%+v\n", ss)
}

func main() {
	// parse the args
	parseArgs(&ss)

	// for random generation
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	// start the server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(ss.server_port, nil))
}
