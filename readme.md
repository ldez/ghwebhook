# GHWebHook

[![release](https://img.shields.io/github/tag/ldez/ghwebhook.svg)](https://github.com/ldez/ghwebhook/releases)
[![Build Status](https://github.com/ldez/ghwebhook/workflows/Main/badge.svg?branch=master)](https://github.com/ldez/ghwebhook/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ldez/ghwebhook)](https://pkg.go.dev/github.com/ldez/ghwebhook/v4)

[![Sponsor](https://img.shields.io/badge/Sponsor%20me-%E2%9D%A4%EF%B8%8F-pink)](https://github.com/sponsors/ldez)

Create a GitHub WebHook in 5 seconds!

## Description

- Default port: `80`
- Default path: `/postreceive`
- Default event type: `push`

## Examples

Basic:
```go
package main

import (
	"log"
	"net/url"

	"github.com/google/go-github/v62/github"
	ghw "github.com/ldez/ghwebhook/v4"
)

func main() {
	eventHandlers := ghw.NewEventHandlers().
		OnIssuesEvent(func(uri *url.URL, deliveryID string, event *github.IssuesEvent) {
			go func() {
				log.Println(uri, deliveryID, event.GetAction(), event.Issue)
			}()
		}).
		OnPullRequestEvent(func(uri *url.URL, deliveryID string, event *github.PullRequestEvent) {
			log.Println(uri, deliveryID, event.GetAction(), event.GetNumber(), event.PullRequest)
		})

	webHook := ghw.NewWebHook(eventHandlers, ghw.WithAllEventTypes)

	err := webHook.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
```

Secured WebHook with custom port and path:

```go
package main

import (
	"log"
	"net/url"

	"github.com/google/go-github/v62/github"
	ghw "github.com/ldez/ghwebhook/v4"
	"github.com/ldez/ghwebhook/v4/eventtype"
)

func main() {
	eventHandlers := ghw.NewEventHandlers().
		OnIssuesEvent(func(uri *url.URL, deliveryID string, event *github.IssuesEvent) {
			go func() {
				log.Println(uri, deliveryID, event.GetAction(), event.Issue)
			}()
		}).
		OnPullRequestEvent(func(uri *url.URL, deliveryID string, event *github.PullRequestEvent) {
			log.Println(uri, deliveryID, event.GetAction(), event.GetNumber(), event.PullRequest)
		})

	webhook := ghw.NewWebHook(
		eventHandlers,
		ghw.WithPort(5000),
		ghw.WithPath("/github"),
		ghw.WithSecret("SECRET"),
		ghw.Debug,
		ghw.WithEventTypes(eventtype.Issues, eventtype.PullRequest))

	err := webhook.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
```

## References

- https://docs.github.com/en/developers/webhooks-and-events/webhooks/about-webhooks
- https://docs.github.com/en/developers/webhooks-and-events/webhooks/securing-your-webhooks
- https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
