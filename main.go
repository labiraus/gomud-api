package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/labiraus/gomud-common/api"
)

// This example demonstrates a trivial echo server.
func main() {
	log.Println("api starting up")
	ctx, ctxDone := context.WithCancel(context.Background())
	done := api.StartBasicApi(ctx, helloHandler)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	ctxDone()
	log.Println("Got signal: " + s.String() + " now closing")
	<-done
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()

	request := userRequest{UserName: "someone"}

	resp, err := api.Post("http://service-user.gomud:8080", request)
	if err != nil {
		err = fmt.Errorf("user request:%v", err)
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	var response = userResponse{}
	api.UnmarshalResponse(&response, resp)
	log.Printf("web handler got %#v\n", response)

	fmt.Fprintf(w, "Greetings %s", response.Greeting)
}

type userRequest struct {
	UserName string
}

type userResponse struct {
	Greeting string
}

// Validate checks if it is a valid request
func (r userRequest) Validate() error {
	return nil
}

// Validate checks if it is a valid request
func (r userResponse) Validate() error {
	return nil
}
