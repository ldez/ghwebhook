package ghwebhook

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ldez/ghwebhook/eventtype"
)

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
		}, {
			name:         "invalid path",
			method:       http.MethodPost,
			path:         "/foo",
			expectedCode: http.StatusNotFound,
		}, {
			name:         "missing X-Github-Event header",
			method:       http.MethodPost,
			path:         defaultPath,
			expectedCode: http.StatusBadRequest,
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
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:   "secured: missing X-Hub-Signature header",
			method: http.MethodPost,
			path:   defaultPath,
			secret: "foo",
			body:   bytes.NewBufferString("{}"),
			headers: map[string]string{
				"X-Github-Event": eventtype.Push,
			},
			expectedCode: http.StatusForbidden,
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
			t.Parallel()

			webHook := NewWebHook(eventHandlers, WithSecret(test.secret))

			req := httptest.NewRequest(test.method, test.path, test.body)
			for key, value := range test.headers {
				req.Header.Add(key, value)
			}

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
		eventType     string
		body          []byte
		expectedError bool
	}{
		{
			name:          "empty body",
			eventType:     eventtype.Push,
			body:          []byte(""),
			expectedError: true,
		},
		{
			name:          "invalid JSON body",
			eventType:     eventtype.Push,
			body:          []byte("{"),
			expectedError: true,
		},
		{
			name:          "valid JSON body",
			eventType:     eventtype.Push,
			body:          []byte("{}"),
			expectedError: false,
		},
		{
			name:          "valid Raw body",
			eventType:     eventtype.Gist,
			body:          []byte("{}"),
			expectedError: false,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := webHook.handleEvents(test.eventType, test.body)

			if test.expectedError && err == nil {
				t.Errorf("Got no error, but want an error.")
			} else if !test.expectedError && err != nil {
				t.Errorf("Got %v, but want no error.", err)
			}
		})
	}
}
