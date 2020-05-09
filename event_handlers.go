package ghwebhook

import (
	"net/url"

	"github.com/google/go-github/v31/github"
)

// EventHandlers all event handlers
type EventHandlers struct {
	onPing func(*url.URL, *github.WebHookPayload, *github.PingEvent)

	onCheckRun                     func(*url.URL, *github.WebHookPayload, *github.CheckRunEvent)
	onCheckSuite                   func(*url.URL, *github.WebHookPayload, *github.CheckSuiteEvent)
	onCommitComment                func(*url.URL, *github.WebHookPayload, *github.CommitCommentEvent)
	onContentReference             func(*url.URL, *github.WebHookPayload, *ContentReferenceEvent)
	onCreate                       func(*url.URL, *github.WebHookPayload, *github.CreateEvent)
	onDelete                       func(*url.URL, *github.WebHookPayload, *github.DeleteEvent)
	onDeployment                   func(*url.URL, *github.WebHookPayload, *github.DeploymentEvent)
	onDeploymentStatus             func(*url.URL, *github.WebHookPayload, *github.DeploymentStatusEvent)
	onFork                         func(*url.URL, *github.WebHookPayload, *github.ForkEvent)
	onGitHubAppAuthorization       func(*url.URL, *github.WebHookPayload, *github.GitHubAppAuthorizationEvent)
	onGollum                       func(*url.URL, *github.WebHookPayload, *github.GollumEvent)
	onInstallation                 func(*url.URL, *github.WebHookPayload, *github.InstallationEvent)
	onInstallationRepositories     func(*url.URL, *github.WebHookPayload, *github.InstallationRepositoriesEvent)
	onIssueComment                 func(*url.URL, *github.WebHookPayload, *github.IssueCommentEvent)
	onIssues                       func(*url.URL, *github.WebHookPayload, *github.IssuesEvent)
	onLabel                        func(*url.URL, *github.WebHookPayload, *github.LabelEvent)
	onMarketplacePurchase          func(*url.URL, *github.WebHookPayload, *github.MarketplacePurchaseEvent)
	onMember                       func(*url.URL, *github.WebHookPayload, *github.MemberEvent)
	onMembership                   func(*url.URL, *github.WebHookPayload, *github.MembershipEvent)
	onMilestone                    func(*url.URL, *github.WebHookPayload, *github.MilestoneEvent)
	onOrganization                 func(*url.URL, *github.WebHookPayload, *github.OrganizationEvent)
	onOrgBlock                     func(*url.URL, *github.WebHookPayload, *github.OrgBlockEvent)
	onPageBuild                    func(*url.URL, *github.WebHookPayload, *github.PageBuildEvent)
	onProjectCard                  func(*url.URL, *github.WebHookPayload, *github.ProjectCardEvent)
	onProjectColumn                func(*url.URL, *github.WebHookPayload, *github.ProjectColumnEvent)
	onProject                      func(*url.URL, *github.WebHookPayload, *github.ProjectEvent)
	onPublic                       func(*url.URL, *github.WebHookPayload, *github.PublicEvent)
	onPullRequest                  func(*url.URL, *github.WebHookPayload, *github.PullRequestEvent)
	onPullRequestReview            func(*url.URL, *github.WebHookPayload, *github.PullRequestReviewEvent)
	onPullRequestReviewComment     func(*url.URL, *github.WebHookPayload, *github.PullRequestReviewCommentEvent)
	onPush                         func(*url.URL, *github.WebHookPayload, *github.PushEvent)
	onRelease                      func(*url.URL, *github.WebHookPayload, *github.ReleaseEvent)
	onRepository                   func(*url.URL, *github.WebHookPayload, *github.RepositoryEvent)
	onRepositoryImport             func(*url.URL, *github.WebHookPayload, *RepositoryImportEvent)
	onRepositoryVulnerabilityAlert func(*url.URL, *github.WebHookPayload, *github.RepositoryVulnerabilityAlertEvent)
	onSecurityAdvisory             func(*url.URL, *github.WebHookPayload, *SecurityAdvisoryEvent)
	onStatus                       func(*url.URL, *github.WebHookPayload, *github.StatusEvent)
	onTeam                         func(*url.URL, *github.WebHookPayload, *github.TeamEvent)
	onTeamAdd                      func(*url.URL, *github.WebHookPayload, *github.TeamAddEvent)
	onWatch                        func(*url.URL, *github.WebHookPayload, *github.WatchEvent)
}

// NewEventHandlers create a new event handlers.
func NewEventHandlers() *EventHandlers {
	return &EventHandlers{}
}

// OnPing Ping handler.
func (c *EventHandlers) OnPing(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PingEvent)) *EventHandlers {
	c.onPing = eventHandler
	return c
}

// OnCheckRun CheckRun handler.
func (c *EventHandlers) OnCheckRun(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.CheckRunEvent)) *EventHandlers {
	c.onCheckRun = eventHandler
	return c
}

// OnCheckSuite CheckSuite handler.
func (c *EventHandlers) OnCheckSuite(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.CheckSuiteEvent)) *EventHandlers {
	c.onCheckSuite = eventHandler
	return c
}

// OnCommitComment CommitComment handler.
func (c *EventHandlers) OnCommitComment(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.CommitCommentEvent)) *EventHandlers {
	c.onCommitComment = eventHandler
	return c
}

// OnContentReference ContentReference handler.
func (c *EventHandlers) OnContentReference(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *ContentReferenceEvent)) *EventHandlers {
	c.onContentReference = eventHandler
	return c
}

// OnCreate Create handler.
func (c *EventHandlers) OnCreate(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.CreateEvent)) *EventHandlers {
	c.onCreate = eventHandler
	return c
}

// OnDelete Delete handler.
func (c *EventHandlers) OnDelete(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.DeleteEvent)) *EventHandlers {
	c.onDelete = eventHandler
	return c
}

// OnDeployment Deployment handler.
func (c *EventHandlers) OnDeployment(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.DeploymentEvent)) *EventHandlers {
	c.onDeployment = eventHandler
	return c
}

// OnDeploymentStatus DeploymentStatus handler.
func (c *EventHandlers) OnDeploymentStatus(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.DeploymentStatusEvent)) *EventHandlers {
	c.onDeploymentStatus = eventHandler
	return c
}

// OnFork Fork handler.
func (c *EventHandlers) OnFork(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.ForkEvent)) *EventHandlers {
	c.onFork = eventHandler
	return c
}

// OnGollum Gollum handler.
func (c *EventHandlers) OnGollum(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.GollumEvent)) *EventHandlers {
	c.onGollum = eventHandler
	return c
}

// OnInstallation Installation handler.
func (c *EventHandlers) OnInstallation(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.InstallationEvent)) *EventHandlers {
	c.onInstallation = eventHandler
	return c
}

// OnInstallationRepositories InstallationRepositories handler.
func (c *EventHandlers) OnInstallationRepositories(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.InstallationRepositoriesEvent)) *EventHandlers {
	c.onInstallationRepositories = eventHandler
	return c
}

// OnIssueComment IssueComment handler.
func (c *EventHandlers) OnIssueComment(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.IssueCommentEvent)) *EventHandlers {
	c.onIssueComment = eventHandler
	return c
}

// OnIssues Issues handler.
func (c *EventHandlers) OnIssues(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.IssuesEvent)) *EventHandlers {
	c.onIssues = eventHandler
	return c
}

// OnLabel Label handler.
func (c *EventHandlers) OnLabel(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.LabelEvent)) *EventHandlers {
	c.onLabel = eventHandler
	return c
}

// OnMarketplacePurchase MarketplacePurchase handler.
func (c *EventHandlers) OnMarketplacePurchase(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.MarketplacePurchaseEvent)) *EventHandlers {
	c.onMarketplacePurchase = eventHandler
	return c
}

// OnMember Member handler.
func (c *EventHandlers) OnMember(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.MemberEvent)) *EventHandlers {
	c.onMember = eventHandler
	return c
}

// OnMembership Membership handler.
func (c *EventHandlers) OnMembership(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.MembershipEvent)) *EventHandlers {
	c.onMembership = eventHandler
	return c
}

// OnMilestone Milestone handler.
func (c *EventHandlers) OnMilestone(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.MilestoneEvent)) *EventHandlers {
	c.onMilestone = eventHandler
	return c
}

// OnOrganization Organization handler.
func (c *EventHandlers) OnOrganization(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.OrganizationEvent)) *EventHandlers {
	c.onOrganization = eventHandler
	return c
}

// OnOrgBlock OrgBlock handler.
func (c *EventHandlers) OnOrgBlock(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.OrgBlockEvent)) *EventHandlers {
	c.onOrgBlock = eventHandler
	return c
}

// OnPageBuild PageBuild handler.
func (c *EventHandlers) OnPageBuild(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PageBuildEvent)) *EventHandlers {
	c.onPageBuild = eventHandler
	return c
}

// OnProjectCard ProjectCard handler.
func (c *EventHandlers) OnProjectCard(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.ProjectCardEvent)) *EventHandlers {
	c.onProjectCard = eventHandler
	return c
}

// OnProjectColumn ProjectColumn handler.
func (c *EventHandlers) OnProjectColumn(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.ProjectColumnEvent)) *EventHandlers {
	c.onProjectColumn = eventHandler
	return c
}

// OnProject Project handler.
func (c *EventHandlers) OnProject(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.ProjectEvent)) *EventHandlers {
	c.onProject = eventHandler
	return c
}

// OnPublic Public handler.
func (c *EventHandlers) OnPublic(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PublicEvent)) *EventHandlers {
	c.onPublic = eventHandler
	return c
}

// OnPullRequest PullRequest handler.
func (c *EventHandlers) OnPullRequest(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PullRequestEvent)) *EventHandlers {
	c.onPullRequest = eventHandler
	return c
}

// OnPullRequestReview PullRequestReview handler.
func (c *EventHandlers) OnPullRequestReview(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PullRequestReviewEvent)) *EventHandlers {
	c.onPullRequestReview = eventHandler
	return c
}

// OnPullRequestReviewComment PullRequestReviewComment handler.
func (c *EventHandlers) OnPullRequestReviewComment(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PullRequestReviewCommentEvent)) *EventHandlers {
	c.onPullRequestReviewComment = eventHandler
	return c
}

// OnPush Push handler.
func (c *EventHandlers) OnPush(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.PushEvent)) *EventHandlers {
	c.onPush = eventHandler
	return c
}

// OnRelease Release handler.
func (c *EventHandlers) OnRelease(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.ReleaseEvent)) *EventHandlers {
	c.onRelease = eventHandler
	return c
}

// OnRepository Repository handler.
func (c *EventHandlers) OnRepository(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.RepositoryEvent)) *EventHandlers {
	c.onRepository = eventHandler
	return c
}

// OnRepositoryImport RepositoryImport handler.
func (c *EventHandlers) OnRepositoryImport(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *RepositoryImportEvent)) *EventHandlers {
	c.onRepositoryImport = eventHandler
	return c
}

// OnRepositoryVulnerabilityAlert RepositoryVulnerabilityAlert handler.
func (c *EventHandlers) OnRepositoryVulnerabilityAlert(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.RepositoryVulnerabilityAlertEvent)) *EventHandlers {
	c.onRepositoryVulnerabilityAlert = eventHandler
	return c
}

// OnSecurityAdvisory SecurityAdvisory handler.
func (c *EventHandlers) OnSecurityAdvisory(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *SecurityAdvisoryEvent)) *EventHandlers {
	c.onSecurityAdvisory = eventHandler
	return c
}

// OnStatus Status handler.
func (c *EventHandlers) OnStatus(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.StatusEvent)) *EventHandlers {
	c.onStatus = eventHandler
	return c
}

// OnTeam Team handler.
func (c *EventHandlers) OnTeam(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.TeamEvent)) *EventHandlers {
	c.onTeam = eventHandler
	return c
}

// OnTeamAdd TeamAdd handler.
func (c *EventHandlers) OnTeamAdd(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.TeamAddEvent)) *EventHandlers {
	c.onTeamAdd = eventHandler
	return c
}

// OnWatch Watch handler.
func (c *EventHandlers) OnWatch(eventHandler func(uri *url.URL, payload *github.WebHookPayload, event *github.WatchEvent)) *EventHandlers {
	c.onWatch = eventHandler
	return c
}
