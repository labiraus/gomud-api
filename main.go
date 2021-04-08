package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/labiraus/gomud-common/api"

	_ "github.com/denisenkom/go-mssqldb"
)

// This example demonstrates a trivial echo server.
func main() {
	fmt.Println("hi")
	ctx, ctxDone := context.WithCancel(context.Background())
	done := api.StartBasicApi(ctx, helloHandler)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	ctxDone()
	fmt.Println("Got signal: " + s.String() + " now closing")
	<-done
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r)
		}
	}()

	request := api.UserRequest{UserName: "someone"}

	resp, err := api.Post("http://user-service.gomud:8080", request)
	if err != nil {
		err = fmt.Errorf("user request:%v", err)
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	var response = api.UserResponse{}
	api.UnmarshalResponse(&response, resp)
	fmt.Printf("web handler got %#v\n", response)

	fmt.Fprintf(w, "Greetings %s", response.Greeting)
}
