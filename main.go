package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func wasteCPU() {
	for i := 0; i < 100000; i++ {
		_ = i * i // for testing cpu profiling
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	wasteCPU()
	log.Println("ready for profiling")
}

func main() {
	http.HandleFunc("/v1/test", handler)
	go func() {
		log.Println("pprof listening at port 6001")
		err := http.ListenAndServe(":6001", nil) // this is where pprof will be accessed via /debug/pprof/profile
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println("listening at port 8001")
	http.ListenAndServe(":8001", nil) // we'll use this to access the handler
}

// access pprof at http://localhost:6001/debug/pprof/profile
// access handler at http://localhost:8001/v1/test

// go tool pprof http://localhost:6001/debug/pprof/profile\?seconds\=10 can be used to profile for CPU usage
// go tool pprof http://localhost:6001/debug/pprof/heap can be used to profile for memory usage

// here if you hit the 8001 port, you will be able to use pprof via port 6001 to look at the profiling metrics.
