package main

import (
	"flag"
	"fmt"
	"net/http"
)

var flagConfig string

func main() {
	flag.StringVar(&flagConfig, "config", "", "-config=[config path]")
	flag.Parse()

	ReadConfig(flagConfig)
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/status", StatusHandler)
	fmt.Println("Start Server")
	http.ListenAndServe(":8080", nil)

}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, config.ServerInfo.EnvName+", Hello, World\n")
	fmt.Fprintf(w, "Version:"+Version+"\n")
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
