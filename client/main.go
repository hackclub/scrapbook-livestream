package main

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	queue := []string{}

	initFile()
	// c, _, err := websocket.DefaultDialer.Dial("ws://hackclub-scrapbook-livestream.herokuapp.com", nil)
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go func() {
		for {
			c.WriteMessage(websocket.TextMessage, []byte("ping"))

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
		} else {
			queue = append(queue, string(msg))
			log.Printf("Adding to queue: %s", string(msg))
		}
	}
}

func initFile() {
	_, err := os.Create("./scrapbook_updates.txt")
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
}
