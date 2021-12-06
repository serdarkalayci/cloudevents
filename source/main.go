package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/kelseyhightower/envconfig"
)

type Sender struct {
	client cloudevents.Client

	// If the K_SINK environment variable is set, then events are sent there,
	// otherwise we simply reply to the inbound request.
	Target string `envconfig:"K_SINK"`
}

// Message is the structure of the event we send
type Message struct {
	Name    string `json:"name,omitempty"`
	Number1 int    `json:"number1"`
	Number2 int    `json:"number2"`
}

func (s Sender) saveHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	number1, _ := strconv.Atoi(r.FormValue("number1"))
	number2, _ := strconv.Atoi(r.FormValue("number2"))

	message := Message{Name: name, Number1: number1, Number2: number2}
	log.Printf("Sending event: %v", message)
	e := cloudevents.NewEvent(cloudevents.VersionV1)
	e.SetType("dev.knative.docs.sample")
	e.SetSource("https://github.com/knative/docs/docs/serving/samples/cloudevents/cloudevents-go")
	if err := e.SetData("application/json", message); err != nil {
		//return cloudevents.NewHTTPResult(500, "failed to set response data: %s", err)
		log.Printf("Error setting data %s", err)
	}

	ctx := cloudevents.ContextWithTarget(context.Background(), s.Target)
	s.client.Send(ctx, e)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {

	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	s := Sender{client: client}
	if err := envconfig.Process("", &s); err != nil {
		log.Fatal(err.Error())
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/save", s.saveHandler)
	http.ListenAndServe(":8080", nil)
}
