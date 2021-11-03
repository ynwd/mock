package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fastrodev/fastrex"
)

func main() {
	app := fastrex.New()
	app.Get("/", func(req fastrex.Request, res fastrex.Response) {
		res.Send("root")
	})
	app.Post("/", func(req fastrex.Request, res fastrex.Response) {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		res.Send(bodyString)
	})
	app.Listen(8080)
}
