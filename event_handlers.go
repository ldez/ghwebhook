package ghwebhook

import "github.com/google/go-github/github"

type eventHandlers struct {
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
func NewEventHandlers() *eventHandlers {
	return &eventHandlers{}
}

// OnPing Ping handler.
func (c *eventHandlers) OnPing(eventHandler func(payload *github.WebHookPayload, event *github.PingEvent)) *eventHandlers {
	c.onPing = eventHandler
	return c
}

// OnCommitComment CommitComment handler.
func (c *eventHandlers) OnCommitComment(eventHandler func(payload *github.WebHookPayload, event *github.CommitCommentEvent)) *eventHandlers {
	c.onCommitComment = eventHandler
	return c
}

// OnCreate Create handler.
func (c *eventHandlers) OnCreate(eventHandler func(payload *github.WebHookPayload, event *github.CreateEvent)) *eventHandlers {
	c.onCreate = eventHandler
	return c
}

// OnDelete Delete handler.
func (c *eventHandlers) OnDelete(eventHandler func(payload *github.WebHookPayload, event *github.DeleteEvent)) *eventHandlers {
	c.onDelete = eventHandler
	return c
}

// OnDeployment Deployment handler.
func (c *eventHandlers) OnDeployment(eventHandler func(payload *github.WebHookPayload, event *github.DeploymentEvent)) *eventHandlers {
	c.onDeployment = eventHandler
	return c
}

// OnDeploymentStatus DeploymentStatus handler.
func (c *eventHandlers) OnDeploymentStatus(eventHandler func(payload *github.WebHookPayload, event *github.DeploymentStatusEvent)) *eventHandlers {
	c.onDeploymentStatus = eventHandler
	return c
}

// OnDownload Download handler.
func (c *eventHandlers) OnDownload(eventHandler func(payload *github.WebHookPayload, event []byte)) *eventHandlers {
	c.onDownload = eventHandler
	return c
}

// OnFollow Follow handler.
func (c *eventHandlers) OnFollow(eventHandler func(payload *github.WebHookPayload, event []byte)) *eventHandlers {
	c.onFollow = eventHandler
	return c
}

// OnFork Fork handler.
func (c *eventHandlers) OnFork(eventHandler func(payload *github.WebHookPayload, event *github.ForkEvent)) *eventHandlers {
	c.onFork = eventHandler
	return c
}

// OnGist Gist handler.
func (c *eventHandlers) OnGist(eventHandler func(payload *github.WebHookPayload, event []byte)) *eventHandlers {
	c.onGist = eventHandler
	return c
}

// OnGollum Gollum handler.
func (c *eventHandlers) OnGollum(eventHandler func(payload *github.WebHookPayload, event *github.GollumEvent)) *eventHandlers {
	c.onGollum = eventHandler
	return c
}

// OnInstallation Installation handler.
func (c *eventHandlers) OnInstallation(eventHandler func(payload *github.WebHookPayload, event *github.InstallationEvent)) *eventHandlers {
	c.onInstallation = eventHandler
	return c
}

// OnInstallationRepositories InstallationRepositories handler.
func (c *eventHandlers) OnInstallationRepositories(eventHandler func(payload *github.WebHookPayload, event *github.InstallationRepositoriesEvent)) *eventHandlers {
	c.onInstallationRepositories = eventHandler
	return c
}

// OnIssueComment IssueComment handler.
func (c *eventHandlers) OnIssueComment(eventHandler func(payload *github.WebHookPayload, event *github.IssueCommentEvent)) *eventHandlers {
	c.onIssueComment = eventHandler
	return c
}

// OnIssues Issues handler.
func (c *eventHandlers) OnIssues(eventHandler func(payload *github.WebHookPayload, event *github.IssuesEvent)) *eventHandlers {
	c.onIssues = eventHandler
	return c
}

// OnLabel Label handler.
func (c *eventHandlers) OnLabel(eventHandler func(payload *github.WebHookPayload, event *github.LabelEvent)) *eventHandlers {
	c.onLabel = eventHandler
	return c
}

// OnMarketplacePurchase MarketplacePurchase handler.
func (c *eventHandlers) OnMarketplacePurchase(eventHandler func(payload *github.WebHookPayload, event []byte)) *eventHandlers {
	c.onMarketplacePurchase = eventHandler
	return c
}

// OnMember Member handler.
func (c *eventHandlers) OnMember(eventHandler func(payload *github.WebHookPayload, event *github.MemberEvent)) *eventHandlers {
	c.onMember = eventHandler
	return c
}

// OnMembership Membership handler.
func (c *eventHandlers) OnMembership(eventHandler func(payload *github.WebHookPayload, event *github.MembershipEvent)) *eventHandlers {
	c.onMembership = eventHandler
	return c
}

// OnMilestone Milestone handler.
func (c *eventHandlers) OnMilestone(eventHandler func(payload *github.WebHookPayload, event *github.MilestoneEvent)) *eventHandlers {
	c.onMilestone = eventHandler
	return c
}

// OnOrganization Organization handler.
func (c *eventHandlers) OnOrganization(eventHandler func(payload *github.WebHookPayload, event *github.OrganizationEvent)) *eventHandlers {
	c.onOrganization = eventHandler
	return c
}

// OnOrgBlock OrgBlock handler.
func (c *eventHandlers) OnOrgBlock(eventHandler func(payload *github.WebHookPayload, event *github.OrgBlockEvent)) *eventHandlers {
	c.onOrgBlock = eventHandler
	return c
}

// OnPageBuild PageBuild handler.
func (c *eventHandlers) OnPageBuild(eventHandler func(payload *github.WebHookPayload, event *github.PageBuildEvent)) *eventHandlers {
	c.onPageBuild = eventHandler
	return c
}

// OnProjectCard ProjectCard handler.
func (c *eventHandlers) OnProjectCard(eventHandler func(payload *github.WebHookPayload, event *github.ProjectCardEvent)) *eventHandlers {
	c.onProjectCard = eventHandler
	return c
}

// OnProjectColumn ProjectColumn handler.
func (c *eventHandlers) OnProjectColumn(eventHandler func(payload *github.WebHookPayload, event *github.ProjectColumnEvent)) *eventHandlers {
	c.onProjectColumn = eventHandler
	return c
}

// OnProject Project handler.
func (c *eventHandlers) OnProject(eventHandler func(payload *github.WebHookPayload, event *github.ProjectEvent)) *eventHandlers {
	c.onProject = eventHandler
	return c
}

// OnPublic Public handler.
func (c *eventHandlers) OnPublic(eventHandler func(payload *github.WebHookPayload, event *github.PublicEvent)) *eventHandlers {
	c.onPublic = eventHandler
	return c
}

// OnPullRequest PullRequest handler.
func (c *eventHandlers) OnPullRequest(eventHandler func(payload *github.WebHookPayload, event *github.PullRequestEvent)) *eventHandlers {
	c.onPullRequest = eventHandler
	return c
}

// OnPullRequestReview PullRequestReview handler.
func (c *eventHandlers) OnPullRequestReview(eventHandler func(payload *github.WebHookPayload, event *github.PullRequestReviewEvent)) *eventHandlers {
	c.onPullRequestReview = eventHandler
	return c
}

// OnPullRequestReviewComment PullRequestReviewComment handler.
func (c *eventHandlers) OnPullRequestReviewComment(eventHandler func(payload *github.WebHookPayload, event *github.PullRequestReviewCommentEvent)) *eventHandlers {
	c.onPullRequestReviewComment = eventHandler
	return c
}

// OnPush Push handler.
func (c *eventHandlers) OnPush(eventHandler func(payload *github.WebHookPayload, event *github.PushEvent)) *eventHandlers {
	c.onPush = eventHandler
	return c
}

// OnRelease Release handler.
func (c *eventHandlers) OnRelease(eventHandler func(payload *github.WebHookPayload, event *github.ReleaseEvent)) *eventHandlers {
	c.onRelease = eventHandler
	return c
}

// OnRepository Repository handler.
func (c *eventHandlers) OnRepository(eventHandler func(payload *github.WebHookPayload, event *github.RepositoryEvent)) *eventHandlers {
	c.onRepository = eventHandler
	return c
}

// OnStatus Status handler.
func (c *eventHandlers) OnStatus(eventHandler func(payload *github.WebHookPayload, event *github.StatusEvent)) *eventHandlers {
	c.onStatus = eventHandler
	return c
}

// OnTeam Team handler.
func (c *eventHandlers) OnTeam(eventHandler func(payload *github.WebHookPayload, event *github.TeamEvent)) *eventHandlers {
	c.onTeam = eventHandler
	return c
}

// OnTeamAdd TeamAdd handler.
func (c *eventHandlers) OnTeamAdd(eventHandler func(payload *github.WebHookPayload, event *github.TeamAddEvent)) *eventHandlers {
	c.onTeamAdd = eventHandler
	return c
}

// OnWatch Watch handler.
func (c *eventHandlers) OnWatch(eventHandler func(payload *github.WebHookPayload, event *github.WatchEvent)) *eventHandlers {
	c.onWatch = eventHandler
	return c
}
