package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/getlantern/systray"
	"github.com/kardianos/service"
)

var server *http.Server

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

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	server = &http.Server{Addr: ":3000"}
	go func(server *http.Server) {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}(server)
	// systray.Run(onReady, onExit)
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "CLIforWindows",
		DisplayName: "CLI for Windows",
		Description: "This is an example of Go CLI for Windows.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 2 {
		if os.Args[1] == "-uninstall" {
			err = s.Uninstall()
			if err != nil {
				log.Fatal(err)
			}
		} else if os.Args[1] == "-install" {
			err = s.Install()
			if err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
