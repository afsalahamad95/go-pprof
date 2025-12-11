package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

var global = make([][]byte, 0)

func wasteHeap() {
	for i := 0; i < 500000; i++ {
		global = append(global, make([]byte, 1024))
	}
}

func handleWasteHeap(w http.ResponseWriter, r *http.Request) {
	wasteHeap()
}

func main() {
	http.HandleFunc("/waste-heap", handleWasteHeap)
	go func() {
		http.ListenAndServe(":9001", nil) // for pprof
	}()
	log.Println("listening on port 6001")
	http.ListenAndServe(":6001", nil)
}

// access heap pprof at http://localhost:9001/pprof/debug/heap?seconds=10
