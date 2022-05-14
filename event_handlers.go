package ghwebhook

import (
	"net/url"

	"github.com/google/go-github/v44/github"
)

// EventHandlers all event handlers.
type EventHandlers struct {
	onBranchProtectionRule         func(*url.URL, *github.BranchProtectionRuleEvent)
	onCheckRun                     func(*url.URL, *github.CheckRunEvent)
	onCheckSuite                   func(*url.URL, *github.CheckSuiteEvent)
	onCommitComment                func(*url.URL, *github.CommitCommentEvent)
	onContentReference             func(*url.URL, *github.ContentReferenceEvent)
	onCreate                       func(*url.URL, *github.CreateEvent)
	onDelete                       func(*url.URL, *github.DeleteEvent)
	onDeployKey                    func(*url.URL, *github.DeployKeyEvent)
	onDeployment                   func(*url.URL, *github.DeploymentEvent)
	onDeploymentStatus             func(*url.URL, *github.DeploymentStatusEvent)
	onDiscussion                   func(*url.URL, *github.DiscussionEvent)
	onFork                         func(*url.URL, *github.ForkEvent)
	onGitHubAppAuthorization       func(*url.URL, *github.GitHubAppAuthorizationEvent)
	onGollum                       func(*url.URL, *github.GollumEvent)
	onInstallation                 func(*url.URL, *github.InstallationEvent)
	onInstallationRepositories     func(*url.URL, *github.InstallationRepositoriesEvent)
	onIssueComment                 func(*url.URL, *github.IssueCommentEvent)
	onIssues                       func(*url.URL, *github.IssuesEvent)
	onLabel                        func(*url.URL, *github.LabelEvent)
	onMarketplacePurchase          func(*url.URL, *github.MarketplacePurchaseEvent)
	onMember                       func(*url.URL, *github.MemberEvent)
	onMembership                   func(*url.URL, *github.MembershipEvent)
	onMeta                         func(*url.URL, *github.MetaEvent)
	onMilestone                    func(*url.URL, *github.MilestoneEvent)
	onOrganization                 func(*url.URL, *github.OrganizationEvent)
	onOrgBlock                     func(*url.URL, *github.OrgBlockEvent)
	onPackage                      func(*url.URL, *github.PackageEvent)
	onPageBuild                    func(*url.URL, *github.PageBuildEvent)
	onPing                         func(*url.URL, *github.PingEvent)
	onProjectCard                  func(*url.URL, *github.ProjectCardEvent)
	onProjectColumn                func(*url.URL, *github.ProjectColumnEvent)
	onProject                      func(*url.URL, *github.ProjectEvent)
	onPublic                       func(*url.URL, *github.PublicEvent)
	onPullRequest                  func(*url.URL, *github.PullRequestEvent)
	onPullRequestReview            func(*url.URL, *github.PullRequestReviewEvent)
	onPullRequestReviewComment     func(*url.URL, *github.PullRequestReviewCommentEvent)
	onPullRequestTarget            func(*url.URL, *github.PullRequestTargetEvent)
	onPush                         func(*url.URL, *github.PushEvent)
	onRelease                      func(*url.URL, *github.ReleaseEvent)
	onRepository                   func(*url.URL, *github.RepositoryEvent)
	onRepositoryDispatch           func(*url.URL, *github.RepositoryDispatchEvent)
	onRepositoryVulnerabilityAlert func(*url.URL, *github.RepositoryVulnerabilityAlertEvent)
	onStar                         func(*url.URL, *github.StarEvent)
	onSecurityAdvisory             func(*url.URL, *github.SecurityAdvisoryEvent)
	onStatus                       func(*url.URL, *github.StatusEvent)
	onTeam                         func(*url.URL, *github.TeamEvent)
	onTeamAdd                      func(*url.URL, *github.TeamAddEvent)
	onUser                         func(*url.URL, *github.UserEvent)
	onWatch                        func(*url.URL, *github.WatchEvent)
	onWorkflowDispatch             func(*url.URL, *github.WorkflowDispatchEvent)
	onWorkflowJob                  func(*url.URL, *github.WorkflowJobEvent)
	onWorkflowRun                  func(*url.URL, *github.WorkflowRunEvent)
}

// NewEventHandlers create a new event handlers.
func NewEventHandlers() *EventHandlers {
	return &EventHandlers{}
}

// OnPing Ping handler.
func (c *EventHandlers) OnPing(eventHandler func(uri *url.URL, event *github.PingEvent)) *EventHandlers {
	c.onPing = eventHandler
	return c
}

// OnCheckRun CheckRun handler.
func (c *EventHandlers) OnCheckRun(eventHandler func(uri *url.URL, event *github.CheckRunEvent)) *EventHandlers {
	c.onCheckRun = eventHandler
	return c
}

// OnCheckSuite CheckSuite handler.
func (c *EventHandlers) OnCheckSuite(eventHandler func(uri *url.URL, event *github.CheckSuiteEvent)) *EventHandlers {
	c.onCheckSuite = eventHandler
	return c
}

// OnCommitComment CommitComment handler.
func (c *EventHandlers) OnCommitComment(eventHandler func(uri *url.URL, event *github.CommitCommentEvent)) *EventHandlers {
	c.onCommitComment = eventHandler
	return c
}

// OnContentReference ContentReference handler.
func (c *EventHandlers) OnContentReference(eventHandler func(uri *url.URL, event *github.ContentReferenceEvent)) *EventHandlers {
	c.onContentReference = eventHandler
	return c
}

// OnCreate Create handler.
func (c *EventHandlers) OnCreate(eventHandler func(uri *url.URL, event *github.CreateEvent)) *EventHandlers {
	c.onCreate = eventHandler
	return c
}

// OnDelete Delete handler.
func (c *EventHandlers) OnDelete(eventHandler func(uri *url.URL, event *github.DeleteEvent)) *EventHandlers {
	c.onDelete = eventHandler
	return c
}

// OnDeployment Deployment handler.
func (c *EventHandlers) OnDeployment(eventHandler func(uri *url.URL, event *github.DeploymentEvent)) *EventHandlers {
	c.onDeployment = eventHandler
	return c
}

// OnDeploymentStatus DeploymentStatus handler.
func (c *EventHandlers) OnDeploymentStatus(eventHandler func(uri *url.URL, event *github.DeploymentStatusEvent)) *EventHandlers {
	c.onDeploymentStatus = eventHandler
	return c
}

// OnFork Fork handler.
func (c *EventHandlers) OnFork(eventHandler func(uri *url.URL, event *github.ForkEvent)) *EventHandlers {
	c.onFork = eventHandler
	return c
}

// OnGollum Gollum handler.
func (c *EventHandlers) OnGollum(eventHandler func(uri *url.URL, event *github.GollumEvent)) *EventHandlers {
	c.onGollum = eventHandler
	return c
}

// OnInstallation Installation handler.
func (c *EventHandlers) OnInstallation(eventHandler func(uri *url.URL, event *github.InstallationEvent)) *EventHandlers {
	c.onInstallation = eventHandler
	return c
}

// OnInstallationRepositories InstallationRepositories handler.
func (c *EventHandlers) OnInstallationRepositories(eventHandler func(uri *url.URL, event *github.InstallationRepositoriesEvent)) *EventHandlers {
	c.onInstallationRepositories = eventHandler
	return c
}

// OnIssueComment IssueComment handler.
func (c *EventHandlers) OnIssueComment(eventHandler func(uri *url.URL, event *github.IssueCommentEvent)) *EventHandlers {
	c.onIssueComment = eventHandler
	return c
}

// OnIssues Issues handler.
func (c *EventHandlers) OnIssues(eventHandler func(uri *url.URL, event *github.IssuesEvent)) *EventHandlers {
	c.onIssues = eventHandler
	return c
}

// OnLabel Label handler.
func (c *EventHandlers) OnLabel(eventHandler func(uri *url.URL, event *github.LabelEvent)) *EventHandlers {
	c.onLabel = eventHandler
	return c
}

// OnMarketplacePurchase MarketplacePurchase handler.
func (c *EventHandlers) OnMarketplacePurchase(eventHandler func(uri *url.URL, event *github.MarketplacePurchaseEvent)) *EventHandlers {
	c.onMarketplacePurchase = eventHandler
	return c
}

// OnMember Member handler.
func (c *EventHandlers) OnMember(eventHandler func(uri *url.URL, event *github.MemberEvent)) *EventHandlers {
	c.onMember = eventHandler
	return c
}

// OnMembership Membership handler.
func (c *EventHandlers) OnMembership(eventHandler func(uri *url.URL, event *github.MembershipEvent)) *EventHandlers {
	c.onMembership = eventHandler
	return c
}

// OnMilestone Milestone handler.
func (c *EventHandlers) OnMilestone(eventHandler func(uri *url.URL, event *github.MilestoneEvent)) *EventHandlers {
	c.onMilestone = eventHandler
	return c
}

// OnOrganization Organization handler.
func (c *EventHandlers) OnOrganization(eventHandler func(uri *url.URL, event *github.OrganizationEvent)) *EventHandlers {
	c.onOrganization = eventHandler
	return c
}

// OnOrgBlock OrgBlock handler.
func (c *EventHandlers) OnOrgBlock(eventHandler func(uri *url.URL, event *github.OrgBlockEvent)) *EventHandlers {
	c.onOrgBlock = eventHandler
	return c
}

// OnPageBuild PageBuild handler.
func (c *EventHandlers) OnPageBuild(eventHandler func(uri *url.URL, event *github.PageBuildEvent)) *EventHandlers {
	c.onPageBuild = eventHandler
	return c
}

// OnProjectCard ProjectCard handler.
func (c *EventHandlers) OnProjectCard(eventHandler func(uri *url.URL, event *github.ProjectCardEvent)) *EventHandlers {
	c.onProjectCard = eventHandler
	return c
}

// OnProjectColumn ProjectColumn handler.
func (c *EventHandlers) OnProjectColumn(eventHandler func(uri *url.URL, event *github.ProjectColumnEvent)) *EventHandlers {
	c.onProjectColumn = eventHandler
	return c
}

// OnProject Project handler.
func (c *EventHandlers) OnProject(eventHandler func(uri *url.URL, event *github.ProjectEvent)) *EventHandlers {
	c.onProject = eventHandler
	return c
}

// OnPublic Public handler.
func (c *EventHandlers) OnPublic(eventHandler func(uri *url.URL, event *github.PublicEvent)) *EventHandlers {
	c.onPublic = eventHandler
	return c
}

// OnPullRequest PullRequest handler.
func (c *EventHandlers) OnPullRequest(eventHandler func(uri *url.URL, event *github.PullRequestEvent)) *EventHandlers {
	c.onPullRequest = eventHandler
	return c
}

// OnPullRequestReview PullRequestReview handler.
func (c *EventHandlers) OnPullRequestReview(eventHandler func(uri *url.URL, event *github.PullRequestReviewEvent)) *EventHandlers {
	c.onPullRequestReview = eventHandler
	return c
}

// OnPullRequestReviewComment PullRequestReviewComment handler.
func (c *EventHandlers) OnPullRequestReviewComment(eventHandler func(uri *url.URL, event *github.PullRequestReviewCommentEvent)) *EventHandlers {
	c.onPullRequestReviewComment = eventHandler
	return c
}

// OnPush Push handler.
func (c *EventHandlers) OnPush(eventHandler func(uri *url.URL, event *github.PushEvent)) *EventHandlers {
	c.onPush = eventHandler
	return c
}

// OnRelease Release handler.
func (c *EventHandlers) OnRelease(eventHandler func(uri *url.URL, event *github.ReleaseEvent)) *EventHandlers {
	c.onRelease = eventHandler
	return c
}

// OnRepository Repository handler.
func (c *EventHandlers) OnRepository(eventHandler func(uri *url.URL, event *github.RepositoryEvent)) *EventHandlers {
	c.onRepository = eventHandler
	return c
}

// OnRepositoryVulnerabilityAlert RepositoryVulnerabilityAlert handler.
func (c *EventHandlers) OnRepositoryVulnerabilityAlert(eventHandler func(uri *url.URL, event *github.RepositoryVulnerabilityAlertEvent)) *EventHandlers {
	c.onRepositoryVulnerabilityAlert = eventHandler
	return c
}

// OnSecurityAdvisory SecurityAdvisory handler.
func (c *EventHandlers) OnSecurityAdvisory(eventHandler func(uri *url.URL, event *github.SecurityAdvisoryEvent)) *EventHandlers {
	c.onSecurityAdvisory = eventHandler
	return c
}

// OnStatus Status handler.
func (c *EventHandlers) OnStatus(eventHandler func(uri *url.URL, event *github.StatusEvent)) *EventHandlers {
	c.onStatus = eventHandler
	return c
}

// OnTeam Team handler.
func (c *EventHandlers) OnTeam(eventHandler func(uri *url.URL, event *github.TeamEvent)) *EventHandlers {
	c.onTeam = eventHandler
	return c
}

// OnTeamAdd TeamAdd handler.
func (c *EventHandlers) OnTeamAdd(eventHandler func(uri *url.URL, event *github.TeamAddEvent)) *EventHandlers {
	c.onTeamAdd = eventHandler
	return c
}

// OnWatch Watch handler.
func (c *EventHandlers) OnWatch(eventHandler func(uri *url.URL, event *github.WatchEvent)) *EventHandlers {
	c.onWatch = eventHandler
	return c
}
