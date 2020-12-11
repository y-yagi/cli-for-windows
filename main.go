package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/getlantern/systray"
)

var server *http.Server

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	server = &http.Server{Addr: ":3000"}
	go func(server *http.Server) {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}(server)
	systray.Run(onReady, onExit)

	return
}

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("My CLI")
	systray.SetTooltip("My CLI")
	mQuit := systray.AddMenuItem("Quit", "Quit")
	<-mQuit.ClickedCh
	systray.Quit()
}

func onExit() {
	server.Close()
}
