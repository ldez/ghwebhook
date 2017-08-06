# GHWebHook

Create a Github WebHook in 5 seconds!

## Description

- Default port: `80`
- Default path: `/ghwebhook`
- Default event type: `push`

## Examples

Basic:
```go
package main

import (
	"log"

	"github.com/google/go-github/github"
	ghw "github.com/ldez/ghwebhook"
)

func main() {

	eventHandlers := ghw.NewEventHandlers().
		OnIssues(func(payload *github.WebHookPayload, event *github.IssuesEvent) {
			go func() {
				log.Println(event.Action, event.Issue)
			}()
		}).
		OnPullRequest(func(payload *github.WebHookPayload, event *github.PullRequestEvent) {
			log.Println(event.Action, event.Number, event.PullRequest)
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

	"github.com/google/go-github/github"
	ghw "github.com/ldez/ghwebhook"
	"github.com/ldez/ghwebhook/eventtype"
)

func main() {

	eventHandlers := ghw.NewEventHandlers().
		OnIssues(func(payload *github.WebHookPayload, event *github.IssuesEvent) {
			go func() {
				log.Println(event.Action, event.Issue)
			}()
		}).
		OnPullRequest(func(payload *github.WebHookPayload, event *github.PullRequestEvent) {
			log.Println(event.Action, event.Number, event.PullRequest)
		})

	webhook := ghw.NewWebHook(
		eventHandlers,
		ghw.WithPort(5000),
		ghw.WithPath("/github"),
		ghw.WithSecret("SECRET"),
		ghw.WithEventTypes(eventtype.Issues, eventtype.PullRequest))

	err := webhook.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
```

## References

- https://developer.github.com/webhooks/
- https://developer.github.com/webhooks/securing/
- https://developer.github.com/v3/activity/events/types/
