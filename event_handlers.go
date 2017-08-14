package ghwebhook

import "github.com/google/go-github/github"

// EventHandlers all event handlers
type EventHandlers struct {
	onPing                     func(*github.WebHookPayload, *github.PingEvent)
	onCommitComment            func(*github.WebHookPayload, *github.CommitCommentEvent)
	onCreate                   func(*github.WebHookPayload, *github.CreateEvent)
	onDelete                   func(*github.WebHookPayload, *github.DeleteEvent)
	onDeployment               func(*github.WebHookPayload, *github.DeploymentEvent)
	onDeploymentStatus         func(*github.WebHookPayload, *github.DeploymentStatusEvent)
	onDownload                 func(*github.WebHookPayload, []byte)
	onFollow                   func(*github.WebHookPayload, []byte)
	onFork                     func(*github.WebHookPayload, *github.ForkEvent)
	onGist                     func(*github.WebHookPayload, []byte)
	onGollum                   func(*github.WebHookPayload, *github.GollumEvent)
	onInstallation             func(*github.WebHookPayload, *github.InstallationEvent)
	onInstallationRepositories func(*github.WebHookPayload, *github.InstallationRepositoriesEvent)
	onIssueComment             func(*github.WebHookPayload, *github.IssueCommentEvent)
	onIssues                   func(*github.WebHookPayload, *github.IssuesEvent)
	onLabel                    func(*github.WebHookPayload, *github.LabelEvent)
	onMarketplacePurchase      func(*github.WebHookPayload, []byte)
	onMember                   func(*github.WebHookPayload, *github.MemberEvent)
	onMembership               func(*github.WebHookPayload, *github.MembershipEvent)
	onMilestone                func(*github.WebHookPayload, *github.MilestoneEvent)
	onOrganization             func(*github.WebHookPayload, *github.OrganizationEvent)
	onOrgBlock                 func(*github.WebHookPayload, *github.OrgBlockEvent)
	onPageBuild                func(*github.WebHookPayload, *github.PageBuildEvent)
	onProjectCard              func(*github.WebHookPayload, *github.ProjectCardEvent)
	onProjectColumn            func(*github.WebHookPayload, *github.ProjectColumnEvent)
	onProject                  func(*github.WebHookPayload, *github.ProjectEvent)
	onPublic                   func(*github.WebHookPayload, *github.PublicEvent)
	onPullRequest              func(*github.WebHookPayload, *github.PullRequestEvent)
	onPullRequestReview        func(*github.WebHookPayload, *github.PullRequestReviewEvent)
	onPullRequestReviewComment func(*github.WebHookPayload, *github.PullRequestReviewCommentEvent)
	onPush                     func(*github.WebHookPayload, *github.PushEvent)
	onRelease                  func(*github.WebHookPayload, *github.ReleaseEvent)
	onRepository               func(*github.WebHookPayload, *github.RepositoryEvent)
	onStatus                   func(*github.WebHookPayload, *github.StatusEvent)
	onTeam                     func(*github.WebHookPayload, *github.TeamEvent)
	onTeamAdd                  func(*github.WebHookPayload, *github.TeamAddEvent)
	onWatch                    func(*github.WebHookPayload, *github.WatchEvent)
}

// NewEventHandlers create a new event handlers.
func NewEventHandlers() *EventHandlers {
	return &EventHandlers{}
}

// OnPing Ping handler.
func (c *EventHandlers) OnPing(eventHandler func(payload *github.WebHookPayload, event *github.PingEvent)) *EventHandlers {
	c.onPing = eventHandler
	return c
}

// OnCommitComment CommitComment handler.
func (c *EventHandlers) OnCommitComment(eventHandler func(payload *github.WebHookPayload, event *github.CommitCommentEvent)) *EventHandlers {
	c.onCommitComment = eventHandler
	return c
}

// OnCreate Create handler.
func (c *EventHandlers) OnCreate(eventHandler func(payload *github.WebHookPayload, event *github.CreateEvent)) *EventHandlers {
	c.onCreate = eventHandler
	return c
}

// OnDelete Delete handler.
func (c *EventHandlers) OnDelete(eventHandler func(payload *github.WebHookPayload, event *github.DeleteEvent)) *EventHandlers {
	c.onDelete = eventHandler
	return c
}

// OnDeployment Deployment handler.
func (c *EventHandlers) OnDeployment(eventHandler func(payload *github.WebHookPayload, event *github.DeploymentEvent)) *EventHandlers {
	c.onDeployment = eventHandler
	return c
}

// OnDeploymentStatus DeploymentStatus handler.
func (c *EventHandlers) OnDeploymentStatus(eventHandler func(payload *github.WebHookPayload, event *github.DeploymentStatusEvent)) *EventHandlers {
	c.onDeploymentStatus = eventHandler
	return c
}

// OnDownload Download handler.
func (c *EventHandlers) OnDownload(eventHandler func(payload *github.WebHookPayload, event []byte)) *EventHandlers {
	c.onDownload = eventHandler
	return c
}

// OnFollow Follow handler.
func (c *EventHandlers) OnFollow(eventHandler func(payload *github.WebHookPayload, event []byte)) *EventHandlers {
	c.onFollow = eventHandler
	return c
}

// OnFork Fork handler.
func (c *EventHandlers) OnFork(eventHandler func(payload *github.WebHookPayload, event *github.ForkEvent)) *EventHandlers {
	c.onFork = eventHandler
	return c
}

// OnGist Gist handler.
func (c *EventHandlers) OnGist(eventHandler func(payload *github.WebHookPayload, event []byte)) *EventHandlers {
	c.onGist = eventHandler
	return c
}

// OnGollum Gollum handler.
func (c *EventHandlers) OnGollum(eventHandler func(payload *github.WebHookPayload, event *github.GollumEvent)) *EventHandlers {
	c.onGollum = eventHandler
	return c
}

// OnInstallation Installation handler.
func (c *EventHandlers) OnInstallation(eventHandler func(payload *github.WebHookPayload, event *github.InstallationEvent)) *EventHandlers {
	c.onInstallation = eventHandler
	return c
}

// OnInstallationRepositories InstallationRepositories handler.
func (c *EventHandlers) OnInstallationRepositories(eventHandler func(payload *github.WebHookPayload, event *github.InstallationRepositoriesEvent)) *EventHandlers {
	c.onInstallationRepositories = eventHandler
	return c
}

// OnIssueComment IssueComment handler.
func (c *EventHandlers) OnIssueComment(eventHandler func(payload *github.WebHookPayload, event *github.IssueCommentEvent)) *EventHandlers {
	c.onIssueComment = eventHandler
	return c
}

// OnIssues Issues handler.
func (c *EventHandlers) OnIssues(eventHandler func(payload *github.WebHookPayload, event *github.IssuesEvent)) *EventHandlers {
	c.onIssues = eventHandler
	return c
}

// OnLabel Label handler.
func (c *EventHandlers) OnLabel(eventHandler func(payload *github.WebHookPayload, event *github.LabelEvent)) *EventHandlers {
	c.onLabel = eventHandler
	return c
}

// OnMarketplacePurchase MarketplacePurchase handler.
func (c *EventHandlers) OnMarketplacePurchase(eventHandler func(payload *github.WebHookPayload, event []byte)) *EventHandlers {
	c.onMarketplacePurchase = eventHandler
	return c
}

// OnMember Member handler.
func (c *EventHandlers) OnMember(eventHandler func(payload *github.WebHookPayload, event *github.MemberEvent)) *EventHandlers {
	c.onMember = eventHandler
	return c
}

// OnMembership Membership handler.
func (c *EventHandlers) OnMembership(eventHandler func(payload *github.WebHookPayload, event *github.MembershipEvent)) *EventHandlers {
	c.onMembership = eventHandler
	return c
}

// OnMilestone Milestone handler.
func (c *EventHandlers) OnMilestone(eventHandler func(payload *github.WebHookPayload, event *github.MilestoneEvent)) *EventHandlers {
	c.onMilestone = eventHandler
	return c
}

// OnOrganization Organization handler.
func (c *EventHandlers) OnOrganization(eventHandler func(payload *github.WebHookPayload, event *github.OrganizationEvent)) *EventHandlers {
	c.onOrganization = eventHandler
	return c
}

// OnOrgBlock OrgBlock handler.
func (c *EventHandlers) OnOrgBlock(eventHandler func(payload *github.WebHookPayload, event *github.OrgBlockEvent)) *EventHandlers {
	c.onOrgBlock = eventHandler
	return c
}

// OnPageBuild PageBuild handler.
func (c *EventHandlers) OnPageBuild(eventHandler func(payload *github.WebHookPayload, event *github.PageBuildEvent)) *EventHandlers {
	c.onPageBuild = eventHandler
	return c
}

// OnProjectCard ProjectCard handler.
func (c *EventHandlers) OnProjectCard(eventHandler func(payload *github.WebHookPayload, event *github.ProjectCardEvent)) *EventHandlers {
	c.onProjectCard = eventHandler
	return c
}

// OnProjectColumn ProjectColumn handler.
func (c *EventHandlers) OnProjectColumn(eventHandler func(payload *github.WebHookPayload, event *github.ProjectColumnEvent)) *EventHandlers {
	c.onProjectColumn = eventHandler
	return c
}

// OnProject Project handler.
func (c *EventHandlers) OnProject(eventHandler func(payload *github.WebHookPayload, event *github.ProjectEvent)) *EventHandlers {
	c.onProject = eventHandler
	return c
}

// OnPublic Public handler.
func (c *EventHandlers) OnPublic(eventHandler func(payload *github.WebHookPayload, event *github.PublicEvent)) *EventHandlers {
	c.onPublic = eventHandler
	return c
}

// OnPullRequest PullRequest handler.
func (c *EventHandlers) OnPullRequest(eventHandler func(payload *github.WebHookPayload, event *github.PullRequestEvent)) *EventHandlers {
	c.onPullRequest = eventHandler
	return c
}

// OnPullRequestReview PullRequestReview handler.
func (c *EventHandlers) OnPullRequestReview(eventHandler func(payload *github.WebHookPayload, event *github.PullRequestReviewEvent)) *EventHandlers {
	c.onPullRequestReview = eventHandler
	return c
}

// OnPullRequestReviewComment PullRequestReviewComment handler.
func (c *EventHandlers) OnPullRequestReviewComment(eventHandler func(payload *github.WebHookPayload, event *github.PullRequestReviewCommentEvent)) *EventHandlers {
	c.onPullRequestReviewComment = eventHandler
	return c
}

// OnPush Push handler.
func (c *EventHandlers) OnPush(eventHandler func(payload *github.WebHookPayload, event *github.PushEvent)) *EventHandlers {
	c.onPush = eventHandler
	return c
}

// OnRelease Release handler.
func (c *EventHandlers) OnRelease(eventHandler func(payload *github.WebHookPayload, event *github.ReleaseEvent)) *EventHandlers {
	c.onRelease = eventHandler
	return c
}

// OnRepository Repository handler.
func (c *EventHandlers) OnRepository(eventHandler func(payload *github.WebHookPayload, event *github.RepositoryEvent)) *EventHandlers {
	c.onRepository = eventHandler
	return c
}

// OnStatus Status handler.
func (c *EventHandlers) OnStatus(eventHandler func(payload *github.WebHookPayload, event *github.StatusEvent)) *EventHandlers {
	c.onStatus = eventHandler
	return c
}

// OnTeam Team handler.
func (c *EventHandlers) OnTeam(eventHandler func(payload *github.WebHookPayload, event *github.TeamEvent)) *EventHandlers {
	c.onTeam = eventHandler
	return c
}

// OnTeamAdd TeamAdd handler.
func (c *EventHandlers) OnTeamAdd(eventHandler func(payload *github.WebHookPayload, event *github.TeamAddEvent)) *EventHandlers {
	c.onTeamAdd = eventHandler
	return c
}

// OnWatch Watch handler.
func (c *EventHandlers) OnWatch(eventHandler func(payload *github.WebHookPayload, event *github.WatchEvent)) *EventHandlers {
	c.onWatch = eventHandler
	return c
}
