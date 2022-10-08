package ghwebhook

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/google/go-github/v47/github"
	"github.com/ldez/ghwebhook/v4/eventtype"
)

//go:generate go run internal/generator.go

const (
	defaultPort       = 80
	defaultPath       = "/postreceive"
	defaultEventTypes = eventtype.Push
)

// WebHook server.
type WebHook struct {
	port          int            // default: 80
	path          *regexp.Regexp // default: "/postreceive"
	secret        string         // optional
	eventTypes    []string       // default: "push"
	eventHandlers EventHandlers
	debug         bool
}

// NewWebHook create a new server as a GitHub WebHook.
func NewWebHook(eventHandlers *EventHandlers, options ...serverOption) *WebHook {
	server := &WebHook{
		port:          defaultPort,
		path:          regexp.MustCompile(defaultPath),
		eventTypes:    []string{defaultEventTypes},
		eventHandlers: *eventHandlers,
	}

	for _, opt := range options {
		opt(server)
	}

	return server
}

// ListenAndServe run GitHub WebHook server.
func (s *WebHook) ListenAndServe() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s)
}

// ServeHTTP HTTP server handler.
func (s *WebHook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Invalid http method: %s", r.Method)
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uri := r.URL
	if !s.path.MatchString(uri.Path) {
		http.NotFound(w, r)
		return
	}

	eventType := r.Header.Get("X-Github-Event")
	if eventType == "" {
		log.Print("Missing X-Github-Event.")
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if !isAcceptedEventType(s.eventTypes, eventType) {
		log.Printf("Unaccepted event. X-Github-Event: %s", eventType)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer func() { _ = r.Body.Close() }()
	}

	payload, err := github.ValidatePayload(r, []byte(s.secret))
	if err != nil {
		log.Printf("invalid payload: %v: %s", err, string(payload))
		http.Error(w, "invalid payload", http.StatusForbidden)
		return
	}

	webHookType := github.WebHookType(r)
	event, err := github.ParseWebHook(webHookType, payload)
	if err != nil {
		log.Printf("%s: payload parsing: %v: %s", webHookType, err, string(payload))
		http.Error(w, "payload parsing", http.StatusBadRequest)
		return
	}

	deliveryID := github.DeliveryID(r)

	err = s.handleEvents(uri, deliveryID, event)
	if err != nil {
		log.Printf("Failed to unmashall request body: %v", err)
		log.Printf("Failed to unmashall request body: %s", string(payload))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprintf(w, `{"message": "Event %s received"}`, webHookType)
}

func isAcceptedEventType(authorizedEventTypes []string, eventType string) bool {
	if eventType == eventtype.Ping {
		return true
	}
	for _, aet := range authorizedEventTypes {
		if eventType == aet {
			return true
		}
	}
	return false
}
