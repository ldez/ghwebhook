package ghwebhook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/google/go-github/v32/github"
	"github.com/ldez/ghwebhook/v2/eventtype"
)

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

	if !s.path.MatchString(r.URL.Path) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if s.secret != "" {
		signature := r.Header.Get("X-Hub-Signature")
		if signature == "" {
			log.Printf("Missing X-Hub-Signature")
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}

		errSign := validateSignature(signature, s.secret, body)
		if errSign != nil {
			log.Print(errSign)
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
	}

	if s.debug {
		log.Println(string(body))
	}

	err = s.handleEvents(r.URL, eventType, body)
	if err != nil {
		log.Printf("Failed to unmashall request body: %v", err)
		log.Printf("Failed to unmashall request body: %s", string(body))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Event received.")
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

func validateSignature(signature string, secret string, body []byte) error {
	mac := hmac.New(sha1.New, []byte(secret))

	_, err := mac.Write(body)
	if err != nil {
		return err
	}

	expectedMAC := mac.Sum(nil)
	expectedSig := "sha1=" + hex.EncodeToString(expectedMAC)

	if !hmac.Equal([]byte(expectedSig), []byte(signature)) {
		return errors.New("signature verification failed")
	}

	return nil
}

func (s *WebHook) handleEvents(uri *url.URL, eventType string, body []byte) error {
	whPayload, err := parseWebHookPayload(body)
	if err != nil {
		return err
	}

	switch eventType {
	case eventtype.Ping:
		if s.eventHandlers.onPing != nil {
			event := &github.PingEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPing(uri, whPayload, event)
		}
	case eventtype.CheckRun:
		if s.eventHandlers.onCheckRun != nil {
			event := &github.CheckRunEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onCheckRun(uri, whPayload, event)
		}
	case eventtype.CheckSuite:
		if s.eventHandlers.onCheckSuite != nil {
			event := &github.CheckSuiteEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onCheckSuite(uri, whPayload, event)
		}
	case eventtype.CommitComment:
		if s.eventHandlers.onCommitComment != nil {
			event := &github.CommitCommentEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onCommitComment(uri, whPayload, event)
		}
	case eventtype.ContentReference:
		if s.eventHandlers.onContentReference != nil {
			event := &ContentReferenceEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onContentReference(uri, whPayload, event)
		}
	case eventtype.Create:
		if s.eventHandlers.onCreate != nil {
			event := &github.CreateEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onCreate(uri, whPayload, event)
		}
	case eventtype.Delete:
		if s.eventHandlers.onDelete != nil {
			event := &github.DeleteEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onDelete(uri, whPayload, event)
		}
	case eventtype.Deployment:
		if s.eventHandlers.onDeployment != nil {
			event := &github.DeploymentEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onDeployment(uri, whPayload, event)
		}
	case eventtype.DeploymentStatus:
		if s.eventHandlers.onDeploymentStatus != nil {
			event := &github.DeploymentStatusEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onDeploymentStatus(uri, whPayload, event)
		}
	case eventtype.Fork:
		if s.eventHandlers.onFork != nil {
			event := &github.ForkEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onFork(uri, whPayload, event)
		}
	case eventtype.GitHubAppAuthorization:
		if s.eventHandlers.onGitHubAppAuthorization != nil {
			event := &github.GitHubAppAuthorizationEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onGitHubAppAuthorization(uri, whPayload, event)
		}
	case eventtype.Gollum:
		if s.eventHandlers.onGollum != nil {
			event := &github.GollumEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onGollum(uri, whPayload, event)
		}
	case eventtype.Installation:
		if s.eventHandlers.onInstallation != nil {
			event := &github.InstallationEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onInstallation(uri, whPayload, event)
		}
	case eventtype.InstallationRepositories:
		if s.eventHandlers.onInstallationRepositories != nil {
			event := &github.InstallationRepositoriesEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onInstallationRepositories(uri, whPayload, event)
		}
	case eventtype.IssueComment:
		if s.eventHandlers.onIssueComment != nil {
			event := &github.IssueCommentEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onIssueComment(uri, whPayload, event)
		}
	case eventtype.Issues:
		if s.eventHandlers.onIssues != nil {
			event := &github.IssuesEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onIssues(uri, whPayload, event)
		}
	case eventtype.Label:
		if s.eventHandlers.onLabel != nil {
			event := &github.LabelEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onLabel(uri, whPayload, event)
		}
	case eventtype.MarketplacePurchase:
		if s.eventHandlers.onMarketplacePurchase != nil {
			event := &github.MarketplacePurchaseEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onMarketplacePurchase(uri, whPayload, event)
		}
	case eventtype.Member:
		if s.eventHandlers.onMember != nil {
			event := &github.MemberEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onMember(uri, whPayload, event)
		}
	case eventtype.Membership:
		if s.eventHandlers.onMembership != nil {
			event := &github.MembershipEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onMembership(uri, whPayload, event)
		}
	case eventtype.Milestone:
		if s.eventHandlers.onMilestone != nil {
			event := &github.MilestoneEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onMilestone(uri, whPayload, event)
		}
	case eventtype.Organization:
		if s.eventHandlers.onOrganization != nil {
			event := &github.OrganizationEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onOrganization(uri, whPayload, event)
		}
	case eventtype.OrgBlock:
		if s.eventHandlers.onOrgBlock != nil {
			event := &github.OrgBlockEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onOrgBlock(uri, whPayload, event)
		}
	case eventtype.PageBuild:
		if s.eventHandlers.onPageBuild != nil {
			event := &github.PageBuildEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPageBuild(uri, whPayload, event)
		}
	case eventtype.ProjectCard:
		if s.eventHandlers.onProjectCard != nil {
			event := &github.ProjectCardEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onProjectCard(uri, whPayload, event)
		}
	case eventtype.ProjectColumn:
		if s.eventHandlers.onProjectColumn != nil {
			event := &github.ProjectColumnEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onProjectColumn(uri, whPayload, event)
		}
	case eventtype.Project:
		if s.eventHandlers.onProject != nil {
			event := &github.ProjectEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onProject(uri, whPayload, event)
		}
	case eventtype.Public:
		if s.eventHandlers.onPublic != nil {
			event := &github.PublicEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPublic(uri, whPayload, event)
		}
	case eventtype.PullRequest:
		if s.eventHandlers.onPullRequest != nil {
			event := &github.PullRequestEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPullRequest(uri, whPayload, event)
		}
	case eventtype.PullRequestReview:
		if s.eventHandlers.onPullRequestReview != nil {
			event := &github.PullRequestReviewEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPullRequestReview(uri, whPayload, event)
		}
	case eventtype.PullRequestReviewComment:
		if s.eventHandlers.onPullRequestReviewComment != nil {
			event := &github.PullRequestReviewCommentEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPullRequestReviewComment(uri, whPayload, event)
		}
	case eventtype.Push:
		if s.eventHandlers.onPush != nil {
			event := &github.PushEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onPush(uri, whPayload, event)
		}
	case eventtype.Release:
		if s.eventHandlers.onRelease != nil {
			event := &github.ReleaseEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onRelease(uri, whPayload, event)
		}
	case eventtype.Repository:
		if s.eventHandlers.onRepository != nil {
			event := &github.RepositoryEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onRepository(uri, whPayload, event)
		}
	case eventtype.RepositoryImport:
		if s.eventHandlers.onRepositoryImport != nil {
			event := &RepositoryImportEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onRepositoryImport(uri, whPayload, event)
		}
	case eventtype.RepositoryVulnerabilityAlert:
		if s.eventHandlers.onRepositoryVulnerabilityAlert != nil {
			event := &github.RepositoryVulnerabilityAlertEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onRepositoryVulnerabilityAlert(uri, whPayload, event)
		}
	case eventtype.SecurityAdvisory:
		if s.eventHandlers.onSecurityAdvisory != nil {
			event := &SecurityAdvisoryEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onSecurityAdvisory(uri, whPayload, event)
		}
	case eventtype.Status:
		if s.eventHandlers.onStatus != nil {
			event := &github.StatusEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onStatus(uri, whPayload, event)
		}
	case eventtype.Team:
		if s.eventHandlers.onTeam != nil {
			event := &github.TeamEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onTeam(uri, whPayload, event)
		}
	case eventtype.TeamAdd:
		if s.eventHandlers.onTeamAdd != nil {
			event := &github.TeamAddEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onTeamAdd(uri, whPayload, event)
		}
	case eventtype.Watch:
		if s.eventHandlers.onWatch != nil {
			event := &github.WatchEvent{}
			err := json.Unmarshal(body, event)
			if err != nil {
				return err
			}
			s.eventHandlers.onWatch(uri, whPayload, event)
		}
	default:
		log.Printf("Unknow event type: %s", eventType)
	}

	return nil
}

func parseWebHookPayload(body []byte) (*github.WebHookPayload, error) {
	whPayload := &github.WebHookPayload{}

	err := json.Unmarshal(body, whPayload)
	if err != nil {
		return nil, err
	}

	return whPayload, nil
}
