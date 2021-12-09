package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/kelseyhightower/envconfig"
)

type Receiver struct {
	client cloudevents.Client
}

func main() {
	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := Receiver{client: client}
	if err := envconfig.Process("", &r); err != nil {
		log.Fatal(err.Error())
	}

	// Depending on whether targeting data has been supplied,
	// we will either reply with our response or send it on to
	// an event sink.
	var receiver interface{} // the SDK reflects on the signature.
	receiver = r.Receive

	if err := client.StartReceiver(context.Background(), receiver); err != nil {
		log.Fatal(err)
	}
}

// Request is the structure of the event we expect to receive.
type Request struct {
	Name    string `json:"name"`
	Number1 int    `json:"number1"`
	Number2 int    `json:"number2"`
	Sleep   int    `json:"sleep"`
	Prime   int    `json:"prime"`
	Bloat   int    `json:"bloat"`
}

// Response is the structure of the event we send in response to requests.
type Response struct {
	Message string `json:"message,omitempty"`
}

// handle shared the logic for producing the Response event from the Request.
func handle(req Request) Response {
	// Consume time, cpu and memory in parallel.
	message := fmt.Sprintf("Hello, %s, the sum of your numbers is %d", req.Name, req.Number1+req.Number2)
	var wg sync.WaitGroup
	if req.Sleep != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			message = fmt.Sprintf("%s - %s", message, sleep(req.Sleep))
		}()
	}
	if req.Prime != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			message = fmt.Sprintf("%s - %s", message, prime(req.Prime))
		}()
	}
	if req.Bloat != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			message = fmt.Sprintf("%s - %s", message, bloat(req.Bloat))
		}()
	}
	wg.Wait()
	return Response{Message: message}
}

// Receive is invoked whenever we receive an event.
func (recv *Receiver) Receive(ctx context.Context, event cloudevents.Event) cloudevents.Result {
	req := Request{}
	if err := event.DataAs(&req); err != nil {
		return cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
	}
	log.Printf("Got an event: %+v", req)

	resp := handle(req)
	log.Printf("Replying with event: %q", resp.Message)

	return cloudevents.NewHTTPResult(200, "OK")
}
