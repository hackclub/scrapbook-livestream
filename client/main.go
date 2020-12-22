package main

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	for {
		doTheThings()
	}
}

func doTheThings() {
	log.Println("initializing...")
	time.Sleep(2 * time.Second)
	log.Println("connecting...")

	dialer := &websocket.Dialer{}

	c, _, err := dialer.Dial("ws://hackclub-scrapbook-livestream.herokuapp.com", nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	queue := []string{}

	initFile()

	go func() {
		for {
			err = c.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				log.Println(err)
				break
			}

			if len(queue) > 0 {
				writeToFile(queue[0])
				log.Printf("Displaying: %s", queue[0])

				queue = queue[1:]
			} else {
				writeToFile("")
				log.Println("clearing file")
			}

			time.Sleep(5 * time.Second)
		}
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		} else {
			queue = append(queue, string(msg))
			log.Printf("Adding to queue: %s", string(msg))
		}
	}
}

func initFile() {
	file, err := os.Create("./scrapbook_updates.txt")
	if err != nil {
		log.Println(err)
	}

	err = file.Close()
	if err != nil {
		log.Println(err)
	}
}

func writeToFile(stuff string) {
	file, err := os.Create("./scrapbook_updates.txt")
	if err != nil {
		log.Println(err)
	} else {
		i, err := file.Write([]byte(stuff))
		if err != nil {
			log.Println(err)
		} else {
			log.Println(i)
		}
	}

	err = file.Close()
	if err != nil {
		log.Println(err)
	}
}
