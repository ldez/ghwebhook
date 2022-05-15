package ghwebhook

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-github/v44/github"
	"github.com/ldez/ghwebhook/v3/eventtype"
)

const FixturesDir = "./test-fixtures"

func TestServeHTTP(t *testing.T) {
	eventHandlers := NewEventHandlers()

	testCases := []struct {
		name         string
		method       string
		body         io.Reader
		path         string
		headers      map[string]string
		secret       string
		expectedCode int
	}{
		{
			name:         "invalid method",
			method:       http.MethodGet,
			path:         defaultPath,
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "invalid path",
			method:       http.MethodPost,
			path:         "/foo",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "missing X-Github-Event header",
			method:       http.MethodPost,
			path:         defaultPath,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:   "ping path prefix",
			method: http.MethodPost,
			path:   defaultPath + "/(.+)",
			headers: map[string]string{
				"X-Github-Event": eventtype.Ping,
			},
			body:         bytes.NewBufferString("{}"),
			expectedCode: http.StatusOK,
		},
		{
			name:   "ping",
			method: http.MethodPost,
			path:   defaultPath,
			headers: map[string]string{
				"X-Github-Event": eventtype.Ping,
			},
			body:         bytes.NewBufferString("{}"),
			expectedCode: http.StatusOK,
		},
		{
			name:   "unaccepted event type",
			method: http.MethodPost,
			path:   defaultPath,
			headers: map[string]string{
				"X-Github-Event": eventtype.Milestone,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:   "empty body",
			method: http.MethodPost,
			path:   defaultPath,
			headers: map[string]string{
				"X-Github-Event": eventtype.Push,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:   "secured: invalid signature",
			method: http.MethodPost,
			path:   defaultPath,
			secret: "foo",
			body:   bytes.NewBufferString("{}"),
			headers: map[string]string{
				"X-Github-Event":  eventtype.Push,
				"X-Hub-Signature": "foo",
			},
			expectedCode: http.StatusForbidden,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			// t.Parallel()

			webHook := NewWebHook(eventHandlers, WithSecret(test.secret))

			req := httptest.NewRequest(test.method, test.path, test.body)
			for key, value := range test.headers {
				req.Header.Add(key, value)
			}

			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			webHook.ServeHTTP(resp, req)

			if resp.Code != test.expectedCode {
				t.Errorf("Got %d, want %d. %q", resp.Code, test.expectedCode, resp.Body.String())
			}
		})
	}
}

func Test_handleEvents(t *testing.T) {
	eventHandlers := NewEventHandlers()
	webHook := NewWebHook(eventHandlers)

	testCases := []struct {
		name          string
		event         interface{}
		expectedError bool
	}{
		{
			name:          "valid event type",
			event:         &github.PushEvent{},
			expectedError: false,
		},
		{
			name:          "invalid event type",
			event:         &struct{}{},
			expectedError: true,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := webHook.handleEvents(nil, "", test.event)

			if test.expectedError && err == nil {
				t.Errorf("Got no error, but want an error.")
			} else if !test.expectedError && err != nil {
				t.Errorf("Got %v, but want no error.", err)
			}
		})
	}
}

func Test_handleEvents_payload(t *testing.T) {
	testCases := []struct {
		name            string
		eventType       string
		fixtureFile     string
		eventHandlers   *EventHandlers
		expectedHandler func(*testing.T, *EventHandlers)
	}{
		{
			name:        "issue event: opened",
			eventType:   eventtype.Issues,
			fixtureFile: "gh-issue_opened.json",
			eventHandlers: NewEventHandlers().
				OnIssuesEvent(func(uri *url.URL, deliveryID string, event *github.IssuesEvent) {
					assertEventAction(t, event, "opened")
				}),
			expectedHandler: func(t *testing.T, eh *EventHandlers) {
				t.Helper()

				if eh.onIssuesEvent == nil {
					t.Error("Got nil, want onIssues function")
				}
			},
		},
		{
			name:        "pull request event: opened",
			eventType:   eventtype.PullRequest,
			fixtureFile: "gh-pr_opened.json",
			eventHandlers: NewEventHandlers().
				OnPullRequestEvent(func(uri *url.URL, deliveryID string, event *github.PullRequestEvent) {
					assertEventAction(t, event, "opened")
				}),
			expectedHandler: func(t *testing.T, eh *EventHandlers) {
				t.Helper()

				if eh.onPullRequestEvent == nil {
					t.Error("Got nil, want onPullRequest function")
				}
			},
		},
		{
			name:        "ping event",
			eventType:   eventtype.Ping,
			fixtureFile: "gh-ping.json",
			eventHandlers: NewEventHandlers().
				OnPingEvent(func(uri *url.URL, deliveryID string, event *github.PingEvent) {
					if event == nil {
						t.Fatal("Got nil, want an event.")
					}

					expectedZen := "Mind your words, they are important."
					if event.GetZen() != expectedZen {
						t.Errorf("Got %s, expected %s", event.GetZen(), expectedZen)
					}
				}),
			expectedHandler: func(t *testing.T, eh *EventHandlers) {
				t.Helper()

				if eh.onPingEvent == nil {
					t.Error("Got nil, want onPing function")
				}
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			webHook := NewWebHook(test.eventHandlers)
			event, err := github.ParseWebHook(test.eventType, mustReadFixtureFile(test.fixtureFile))
			if err != nil {
				t.Errorf("Got %v, but want no error.", err)
			}

			err = webHook.handleEvents(nil, "", event)
			if err != nil {
				t.Errorf("Got %v, but want no error.", err)
			}

			test.expectedHandler(t, test.eventHandlers)
		})
	}
}

func assertEventAction(t *testing.T, event event, action string) {
	t.Helper()

	if event == nil {
		t.Fatal("Got nil, want an event.")
	}

	if event.GetAction() != action {
		t.Errorf("Got %s, expected %s", event.GetAction(), action)
	}
}

type event interface {
	GetAction() string
}

func fixturePath(filename string) string {
	return filepath.Join(FixturesDir, filename)
}

func mustReadFixtureFile(filename string) []byte {
	content, err := os.ReadFile(fixturePath(filename))
	if err != nil {
		panic(err)
	}
	return content
}
