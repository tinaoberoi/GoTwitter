package main

import (
	"encoding/json"
	"os"
	"proj1/server"
	"strconv"
)

func create_config(num_consumers int, mode string) server.Config {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	return server.Config{Encoder: enc, Decoder: dec, Mode: mode, ConsumersCount: num_consumers}
}

func main() {
	var numOfConsumers int
	var Mode string
	if len(os.Args) < 2 {
		numOfConsumers = 0
		Mode = "s"
	} else {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		numOfConsumers = n
		Mode = "p"
	}
	config := create_config(numOfConsumers, Mode)
	server.Run(config)
}
