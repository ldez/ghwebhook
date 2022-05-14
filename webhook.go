package ghwebhook

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/google/go-github/v44/github"
	"github.com/ldez/ghwebhook/v3/eventtype"
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

	err = s.handleEvents(uri, event)
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

func (s *WebHook) handleEvents(uri *url.URL, rawEvent interface{}) error {
	switch event := rawEvent.(type) {
	case *github.BranchProtectionRuleEvent:
		if s.eventHandlers.onBranchProtectionRule != nil {
			s.eventHandlers.onBranchProtectionRule(uri, event)
		}
	case *github.CheckRunEvent:
		if s.eventHandlers.onCheckRun != nil {
			s.eventHandlers.onCheckRun(uri, event)
		}
	case *github.CheckSuiteEvent:
		if s.eventHandlers.onCheckSuite != nil {
			s.eventHandlers.onCheckSuite(uri, event)
		}
	case *github.CommitCommentEvent:
		if s.eventHandlers.onCommitComment != nil {
			s.eventHandlers.onCommitComment(uri, event)
		}
	case *github.ContentReferenceEvent:
		if s.eventHandlers.onContentReference != nil {
			s.eventHandlers.onContentReference(uri, event)
		}
	case *github.CreateEvent:
		if s.eventHandlers.onCreate != nil {
			s.eventHandlers.onCreate(uri, event)
		}
	case *github.DeleteEvent:
		if s.eventHandlers.onDelete != nil {
			s.eventHandlers.onDelete(uri, event)
		}
	case *github.DeployKeyEvent:
		if s.eventHandlers.onDeployKey != nil {
			s.eventHandlers.onDeployKey(uri, event)
		}
	case *github.DeploymentEvent:
		if s.eventHandlers.onDeployment != nil {
			s.eventHandlers.onDeployment(uri, event)
		}
	case *github.DeploymentStatusEvent:
		if s.eventHandlers.onDeploymentStatus != nil {
			s.eventHandlers.onDeploymentStatus(uri, event)
		}
	case *github.DiscussionEvent:
		if s.eventHandlers.onDiscussion != nil {
			s.eventHandlers.onDiscussion(uri, event)
		}
	case *github.ForkEvent:
		if s.eventHandlers.onFork != nil {
			s.eventHandlers.onFork(uri, event)
		}
	case *github.GitHubAppAuthorizationEvent:
		if s.eventHandlers.onGitHubAppAuthorization != nil {
			s.eventHandlers.onGitHubAppAuthorization(uri, event)
		}
	case *github.GollumEvent:
		if s.eventHandlers.onGollum != nil {
			s.eventHandlers.onGollum(uri, event)
		}
	case *github.InstallationEvent:
		if s.eventHandlers.onInstallation != nil {
			s.eventHandlers.onInstallation(uri, event)
		}
	case *github.InstallationRepositoriesEvent:
		if s.eventHandlers.onInstallationRepositories != nil {
			s.eventHandlers.onInstallationRepositories(uri, event)
		}
	case *github.IssueCommentEvent:
		if s.eventHandlers.onIssueComment != nil {
			s.eventHandlers.onIssueComment(uri, event)
		}
	case *github.IssuesEvent:
		if s.eventHandlers.onIssues != nil {
			s.eventHandlers.onIssues(uri, event)
		}
	case *github.LabelEvent:
		if s.eventHandlers.onLabel != nil {
			s.eventHandlers.onLabel(uri, event)
		}
	case *github.MarketplacePurchaseEvent:
		if s.eventHandlers.onMarketplacePurchase != nil {
			s.eventHandlers.onMarketplacePurchase(uri, event)
		}
	case *github.MemberEvent:
		if s.eventHandlers.onMember != nil {
			s.eventHandlers.onMember(uri, event)
		}
	case *github.MembershipEvent:
		if s.eventHandlers.onMembership != nil {
			s.eventHandlers.onMembership(uri, event)
		}
	case *github.MetaEvent:
		if s.eventHandlers.onMeta != nil {
			s.eventHandlers.onMeta(uri, event)
		}
	case *github.MilestoneEvent:
		if s.eventHandlers.onMilestone != nil {
			s.eventHandlers.onMilestone(uri, event)
		}
	case *github.OrganizationEvent:
		if s.eventHandlers.onOrganization != nil {
			s.eventHandlers.onOrganization(uri, event)
		}
	case *github.OrgBlockEvent:
		if s.eventHandlers.onOrgBlock != nil {
			s.eventHandlers.onOrgBlock(uri, event)
		}
	case *github.PackageEvent:
		if s.eventHandlers.onPackage != nil {
			s.eventHandlers.onPackage(uri, event)
		}
	case *github.PageBuildEvent:
		if s.eventHandlers.onPageBuild != nil {
			s.eventHandlers.onPageBuild(uri, event)
		}
	case *github.PingEvent:
		if s.eventHandlers.onPing != nil {
			s.eventHandlers.onPing(uri, event)
		}
	case *github.ProjectEvent:
		if s.eventHandlers.onProject != nil {
			s.eventHandlers.onProject(uri, event)
		}
	case *github.ProjectCardEvent:
		if s.eventHandlers.onProjectCard != nil {
			s.eventHandlers.onProjectCard(uri, event)
		}
	case *github.ProjectColumnEvent:
		if s.eventHandlers.onProjectColumn != nil {
			s.eventHandlers.onProjectColumn(uri, event)
		}
	case *github.PublicEvent:
		if s.eventHandlers.onPublic != nil {
			s.eventHandlers.onPublic(uri, event)
		}
	case *github.PullRequestEvent:
		if s.eventHandlers.onPullRequest != nil {
			s.eventHandlers.onPullRequest(uri, event)
		}
	case *github.PullRequestReviewEvent:
		if s.eventHandlers.onPullRequestReview != nil {
			s.eventHandlers.onPullRequestReview(uri, event)
		}
	case *github.PullRequestReviewCommentEvent:
		if s.eventHandlers.onPullRequestReviewComment != nil {
			s.eventHandlers.onPullRequestReviewComment(uri, event)
		}
	case *github.PullRequestTargetEvent:
		if s.eventHandlers.onPullRequestTarget != nil {
			s.eventHandlers.onPullRequestTarget(uri, event)
		}
	case *github.PushEvent:
		if s.eventHandlers.onPush != nil {
			s.eventHandlers.onPush(uri, event)
		}
	case *github.ReleaseEvent:
		if s.eventHandlers.onRelease != nil {
			s.eventHandlers.onRelease(uri, event)
		}
	case *github.RepositoryEvent:
		if s.eventHandlers.onRepository != nil {
			s.eventHandlers.onRepository(uri, event)
		}
	case *github.RepositoryDispatchEvent:
		if s.eventHandlers.onRepositoryDispatch != nil {
			s.eventHandlers.onRepositoryDispatch(uri, event)
		}
	case *github.RepositoryVulnerabilityAlertEvent:
		if s.eventHandlers.onRepositoryVulnerabilityAlert != nil {
			s.eventHandlers.onRepositoryVulnerabilityAlert(uri, event)
		}
	case *github.SecurityAdvisoryEvent:
		if s.eventHandlers.onSecurityAdvisory != nil {
			s.eventHandlers.onSecurityAdvisory(uri, event)
		}
	case *github.StarEvent:
		if s.eventHandlers.onStar != nil {
			s.eventHandlers.onStar(uri, event)
		}
	case *github.StatusEvent:
		if s.eventHandlers.onStatus != nil {
			s.eventHandlers.onStatus(uri, event)
		}
	case *github.TeamEvent:
		if s.eventHandlers.onTeam != nil {
			s.eventHandlers.onTeam(uri, event)
		}
	case *github.TeamAddEvent:
		if s.eventHandlers.onTeamAdd != nil {
			s.eventHandlers.onTeamAdd(uri, event)
		}
	case *github.UserEvent:
		if s.eventHandlers.onUser != nil {
			s.eventHandlers.onUser(uri, event)
		}
	case *github.WatchEvent:
		if s.eventHandlers.onWatch != nil {
			s.eventHandlers.onWatch(uri, event)
		}
	case *github.WorkflowDispatchEvent:
		if s.eventHandlers.onWorkflowDispatch != nil {
			s.eventHandlers.onWorkflowDispatch(uri, event)
		}
	case *github.WorkflowJobEvent:
		if s.eventHandlers.onWorkflowJob != nil {
			s.eventHandlers.onWorkflowJob(uri, event)
		}
	case *github.WorkflowRunEvent:
		if s.eventHandlers.onWorkflowRun != nil {
			s.eventHandlers.onWorkflowRun(uri, event)
		}
	default:
		return fmt.Errorf("unknow event type: %T", rawEvent)
	}

	return nil
}
