package cli

import (
	"encoding/json"
	"log"
	"os"
)

const (
	cmdServer = "server"
)

func printError(err error) {
	if GlobalOpts.Json {
		j, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{Error: err.Error()})
		log.Println(string(j))
	} else {
		log.Fatalln(err.Error())
	}
}

func exitWithError(err error) {
	printError(err)
	os.Exit(1)
}
