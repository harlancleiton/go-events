package main

import "github.com/harlancleiton/go-events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()
	rabbitmq.Publish("ex-minhafila", "", ch, []byte("Hello, World!"))
}
