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
	"strconv"

	"github.com/google/go-github/github"
	"github.com/ldez/ghwebhook/eventtype"
)

const (
	defaultPort       = 80
	defaultPath       = "/postreceive"
	defaultEventTypes = eventtype.Push
)

type webHook struct {
	port          int      // default: 80
	path          string   // default: "/postreceive"
	secret        string   // optional
	eventTypes    []string // default: "push"
	eventHandlers eventHandlers
}

func NewWebHook(eventHandlers *eventHandlers, options ...serverOption) *webHook {
	server := &webHook{
		port:          defaultPort,
		path:          defaultPath,
		eventTypes:    []string{defaultEventTypes},
		eventHandlers: *eventHandlers,
	}

	for _, opt := range options {
		opt(server)
	}

	return server
}

func (s *webHook) ListenAndServe() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s)
}

func (s *webHook) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Printf("Invalid http method: %s", r.Method)
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != s.path {
		http.NotFound(w, r)
		return
	}

	eventType := r.Header.Get("X-Github-Event")
	if eventType == "" {
		log.Print("Missing X-Github-Event.")
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if eventType == eventtype.Ping {
		log.Print("ping.")
		fmt.Fprint(w, "pong !\n")
		return
	}

	if !isAcceptedEventType(s.eventTypes, eventType) {
		log.Printf("Unaccepted event. X-Github-Event: %s", eventType)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
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

		err := validateSignature(signature, s.secret, body)
		if err != nil {
			log.Print(err)
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
	}

	fmt.Println(string(body))

	err = s.handleEvents(eventType, body)
	if err != nil {
		log.Printf("Failed to unmashall request body: %v", err)
		log.Printf("Failed to unmashall request body: %s", string(body))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Event received.")
}

func isAcceptedEventType(authorizedEventTypes []string, eventType string) bool {
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

func (s *webHook) handleEvents(eventType string, body []byte) error {

	whPayload := &github.WebHookPayload{}
	err := json.Unmarshal(body, whPayload)
	if err != nil {
		return err
	}

	switch eventType {
	case eventtype.Ping:
		event := &github.PingEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPing != nil {
			s.eventHandlers.onPing(whPayload, event)
		}
	case eventtype.CommitComment:
		event := &github.CommitCommentEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onCommitComment != nil {
			s.eventHandlers.onCommitComment(whPayload, event)
		}
	case eventtype.Create:
		event := &github.CreateEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onCreate != nil {
			s.eventHandlers.onCreate(whPayload, event)
		}
	case eventtype.Delete:
		event := &github.DeleteEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onDelete != nil {
			s.eventHandlers.onDelete(whPayload, event)
		}
	case eventtype.Deployment:
		event := &github.DeploymentEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onDeployment != nil {
			s.eventHandlers.onDeployment(whPayload, event)
		}
	case eventtype.DeploymentStatus:
		event := &github.DeploymentStatusEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onDeploymentStatus != nil {
			s.eventHandlers.onDeploymentStatus(whPayload, event)
		}
	case eventtype.Download:
		if s.eventHandlers.onDownload != nil {
			s.eventHandlers.onDownload(whPayload, body)
		}
	case eventtype.Follow:
		if s.eventHandlers.onFollow != nil {
			s.eventHandlers.onFollow(whPayload, body)
		}
	case eventtype.Fork:
		event := &github.ForkEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onFork != nil {
			s.eventHandlers.onFork(whPayload, event)
		}
	case eventtype.Gist:
		if s.eventHandlers.onGist != nil {
			s.eventHandlers.onGist(whPayload, body)
		}
	case eventtype.Gollum:
		event := &github.GollumEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onGollum != nil {
			s.eventHandlers.onGollum(whPayload, event)
		}
	case eventtype.Installation:
		event := &github.InstallationEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onInstallation != nil {
			s.eventHandlers.onInstallation(whPayload, event)
		}
	case eventtype.InstallationRepositories:
		event := &github.InstallationRepositoriesEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onInstallationRepositories != nil {
			s.eventHandlers.onInstallationRepositories(whPayload, event)
		}
	case eventtype.IssueComment:
		event := &github.IssueCommentEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onIssueComment != nil {
			s.eventHandlers.onIssueComment(whPayload, event)
		}
	case eventtype.Issues:
		event := &github.IssuesEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onIssues != nil {
			s.eventHandlers.onIssues(whPayload, event)
		}
	case eventtype.Label:
		event := &github.LabelEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onLabel != nil {
			s.eventHandlers.onLabel(whPayload, event)
		}
	case eventtype.MarketplacePurchase:
		if s.eventHandlers.onMarketplacePurchase != nil {
			s.eventHandlers.onMarketplacePurchase(whPayload, body)
		}
	case eventtype.Member:
		event := &github.MemberEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onMember != nil {
			s.eventHandlers.onMember(whPayload, event)
		}
	case eventtype.Membership:
		event := &github.MembershipEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onMembership != nil {
			s.eventHandlers.onMembership(whPayload, event)
		}
	case eventtype.Milestone:
		event := &github.MilestoneEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onMilestone != nil {
			s.eventHandlers.onMilestone(whPayload, event)
		}
	case eventtype.Organization:
		event := &github.OrganizationEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onOrganization != nil {
			s.eventHandlers.onOrganization(whPayload, event)
		}
	case eventtype.OrgBlock:
		event := &github.OrgBlockEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onOrgBlock != nil {
			s.eventHandlers.onOrgBlock(whPayload, event)
		}
	case eventtype.PageBuild:
		event := &github.PageBuildEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPageBuild != nil {
			s.eventHandlers.onPageBuild(whPayload, event)
		}
	case eventtype.ProjectCard:
		event := &github.ProjectCardEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onProjectCard != nil {
			s.eventHandlers.onProjectCard(whPayload, event)
		}
	case eventtype.ProjectColumn:
		event := &github.ProjectColumnEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onProjectColumn != nil {
			s.eventHandlers.onProjectColumn(whPayload, event)
		}
	case eventtype.Project:
		event := &github.ProjectEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onProject != nil {
			s.eventHandlers.onProject(whPayload, event)
		}
	case eventtype.Public:
		event := &github.PublicEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPublic != nil {
			s.eventHandlers.onPublic(whPayload, event)
		}
	case eventtype.PullRequest:
		event := &github.PullRequestEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPullRequest != nil {
			s.eventHandlers.onPullRequest(whPayload, event)
		}
	case eventtype.PullRequestReview:
		event := &github.PullRequestReviewEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPullRequestReview != nil {
			s.eventHandlers.onPullRequestReview(whPayload, event)
		}
	case eventtype.PullRequestReviewComment:
		event := &github.PullRequestReviewCommentEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPullRequestReviewComment != nil {
			s.eventHandlers.onPullRequestReviewComment(whPayload, event)
		}
	case eventtype.Push:
		event := &github.PushEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onPush != nil {
			s.eventHandlers.onPush(whPayload, event)
		}
	case eventtype.Release:
		event := &github.ReleaseEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onRelease != nil {
			s.eventHandlers.onRelease(whPayload, event)
		}
	case eventtype.Repository:
		event := &github.RepositoryEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onRepository != nil {
			s.eventHandlers.onRepository(whPayload, event)
		}
	case eventtype.Status:
		event := &github.StatusEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onStatus != nil {
			s.eventHandlers.onStatus(whPayload, event)
		}
	case eventtype.Team:
		event := &github.TeamEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onTeam != nil {
			s.eventHandlers.onTeam(whPayload, event)
		}
	case eventtype.TeamAdd:
		event := &github.TeamAddEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onTeamAdd != nil {
			s.eventHandlers.onTeamAdd(whPayload, event)
		}
	case eventtype.Watch:
		event := &github.WatchEvent{}
		err := json.Unmarshal(body, event)
		if err != nil {
			return err
		}
		if s.eventHandlers.onWatch != nil {
			s.eventHandlers.onWatch(whPayload, event)
		}
	default:
		log.Printf("Unknow event type: %s", eventType)
	}

	return nil
}
