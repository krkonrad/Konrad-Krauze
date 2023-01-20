package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ConectionServer(URL string) error {
	const timeout = 1 * time.Minute
	fiinish_conection := time.Now().Add(timeout)
	for i := 0; time.Now().Before(fiinish_conection); i++ {
		_, error := http.Head(URL)
		if error == nil {
			return nil
		}
		log.Printf("no connection to the server (%s); reconnect ...", error)
		time.Sleep(time.Second << uint(i))
	}
	return fmt.Errorf("Server %s noconnection %s ", URL, timeout)
}
func main() {
	if error := ConectionServer("nameserver"); error != nil {
		log.Fatalf("noconnection %v\n", error)

	}
}
