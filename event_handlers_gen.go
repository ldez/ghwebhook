package ghwebhook

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"net/url"

	"github.com/google/go-github/v45/github"
)

// EventHandlers all event handlers.
type EventHandlers struct {
	onBranchProtectionRuleEvent         func(uri *url.URL, deliveryID string, event *github.BranchProtectionRuleEvent)
	onCheckRunEvent                     func(uri *url.URL, deliveryID string, event *github.CheckRunEvent)
	onCheckSuiteEvent                   func(uri *url.URL, deliveryID string, event *github.CheckSuiteEvent)
	onCommitCommentEvent                func(uri *url.URL, deliveryID string, event *github.CommitCommentEvent)
	onContentReferenceEvent             func(uri *url.URL, deliveryID string, event *github.ContentReferenceEvent)
	onCreateEvent                       func(uri *url.URL, deliveryID string, event *github.CreateEvent)
	onDeleteEvent                       func(uri *url.URL, deliveryID string, event *github.DeleteEvent)
	onDeployKeyEvent                    func(uri *url.URL, deliveryID string, event *github.DeployKeyEvent)
	onDeploymentEvent                   func(uri *url.URL, deliveryID string, event *github.DeploymentEvent)
	onDeploymentStatusEvent             func(uri *url.URL, deliveryID string, event *github.DeploymentStatusEvent)
	onDiscussionEvent                   func(uri *url.URL, deliveryID string, event *github.DiscussionEvent)
	onForkEvent                         func(uri *url.URL, deliveryID string, event *github.ForkEvent)
	onGitHubAppAuthorizationEvent       func(uri *url.URL, deliveryID string, event *github.GitHubAppAuthorizationEvent)
	onGollumEvent                       func(uri *url.URL, deliveryID string, event *github.GollumEvent)
	onInstallationEvent                 func(uri *url.URL, deliveryID string, event *github.InstallationEvent)
	onInstallationRepositoriesEvent     func(uri *url.URL, deliveryID string, event *github.InstallationRepositoriesEvent)
	onIssueCommentEvent                 func(uri *url.URL, deliveryID string, event *github.IssueCommentEvent)
	onIssuesEvent                       func(uri *url.URL, deliveryID string, event *github.IssuesEvent)
	onLabelEvent                        func(uri *url.URL, deliveryID string, event *github.LabelEvent)
	onMarketplacePurchaseEvent          func(uri *url.URL, deliveryID string, event *github.MarketplacePurchaseEvent)
	onMemberEvent                       func(uri *url.URL, deliveryID string, event *github.MemberEvent)
	onMembershipEvent                   func(uri *url.URL, deliveryID string, event *github.MembershipEvent)
	onMetaEvent                         func(uri *url.URL, deliveryID string, event *github.MetaEvent)
	onMilestoneEvent                    func(uri *url.URL, deliveryID string, event *github.MilestoneEvent)
	onOrganizationEvent                 func(uri *url.URL, deliveryID string, event *github.OrganizationEvent)
	onOrgBlockEvent                     func(uri *url.URL, deliveryID string, event *github.OrgBlockEvent)
	onPackageEvent                      func(uri *url.URL, deliveryID string, event *github.PackageEvent)
	onPageBuildEvent                    func(uri *url.URL, deliveryID string, event *github.PageBuildEvent)
	onPingEvent                         func(uri *url.URL, deliveryID string, event *github.PingEvent)
	onProjectEvent                      func(uri *url.URL, deliveryID string, event *github.ProjectEvent)
	onProjectCardEvent                  func(uri *url.URL, deliveryID string, event *github.ProjectCardEvent)
	onProjectColumnEvent                func(uri *url.URL, deliveryID string, event *github.ProjectColumnEvent)
	onPublicEvent                       func(uri *url.URL, deliveryID string, event *github.PublicEvent)
	onPullRequestEvent                  func(uri *url.URL, deliveryID string, event *github.PullRequestEvent)
	onPullRequestReviewEvent            func(uri *url.URL, deliveryID string, event *github.PullRequestReviewEvent)
	onPullRequestReviewCommentEvent     func(uri *url.URL, deliveryID string, event *github.PullRequestReviewCommentEvent)
	onPullRequestReviewThreadEvent      func(uri *url.URL, deliveryID string, event *github.PullRequestReviewThreadEvent)
	onPullRequestTargetEvent            func(uri *url.URL, deliveryID string, event *github.PullRequestTargetEvent)
	onPushEvent                         func(uri *url.URL, deliveryID string, event *github.PushEvent)
	onRepositoryEvent                   func(uri *url.URL, deliveryID string, event *github.RepositoryEvent)
	onRepositoryDispatchEvent           func(uri *url.URL, deliveryID string, event *github.RepositoryDispatchEvent)
	onRepositoryVulnerabilityAlertEvent func(uri *url.URL, deliveryID string, event *github.RepositoryVulnerabilityAlertEvent)
	onReleaseEvent                      func(uri *url.URL, deliveryID string, event *github.ReleaseEvent)
	onSecretScanningAlertEvent          func(uri *url.URL, deliveryID string, event *github.SecretScanningAlertEvent)
	onStarEvent                         func(uri *url.URL, deliveryID string, event *github.StarEvent)
	onStatusEvent                       func(uri *url.URL, deliveryID string, event *github.StatusEvent)
	onTeamEvent                         func(uri *url.URL, deliveryID string, event *github.TeamEvent)
	onTeamAddEvent                      func(uri *url.URL, deliveryID string, event *github.TeamAddEvent)
	onUserEvent                         func(uri *url.URL, deliveryID string, event *github.UserEvent)
	onWatchEvent                        func(uri *url.URL, deliveryID string, event *github.WatchEvent)
	onWorkflowDispatchEvent             func(uri *url.URL, deliveryID string, event *github.WorkflowDispatchEvent)
	onWorkflowJobEvent                  func(uri *url.URL, deliveryID string, event *github.WorkflowJobEvent)
	onWorkflowRunEvent                  func(uri *url.URL, deliveryID string, event *github.WorkflowRunEvent)
}

// NewEventHandlers create a new event handlers.
func NewEventHandlers() *EventHandlers {
	return &EventHandlers{}
}

// OnBranchProtectionRuleEvent BranchProtectionRuleEvent handler.
func (c *EventHandlers) OnBranchProtectionRuleEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.BranchProtectionRuleEvent)) *EventHandlers {
	c.onBranchProtectionRuleEvent = eventHandler
	return c
}

// OnCheckRunEvent CheckRunEvent handler.
func (c *EventHandlers) OnCheckRunEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.CheckRunEvent)) *EventHandlers {
	c.onCheckRunEvent = eventHandler
	return c
}

// OnCheckSuiteEvent CheckSuiteEvent handler.
func (c *EventHandlers) OnCheckSuiteEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.CheckSuiteEvent)) *EventHandlers {
	c.onCheckSuiteEvent = eventHandler
	return c
}

// OnCommitCommentEvent CommitCommentEvent handler.
func (c *EventHandlers) OnCommitCommentEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.CommitCommentEvent)) *EventHandlers {
	c.onCommitCommentEvent = eventHandler
	return c
}

// OnContentReferenceEvent ContentReferenceEvent handler.
func (c *EventHandlers) OnContentReferenceEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.ContentReferenceEvent)) *EventHandlers {
	c.onContentReferenceEvent = eventHandler
	return c
}

// OnCreateEvent CreateEvent handler.
func (c *EventHandlers) OnCreateEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.CreateEvent)) *EventHandlers {
	c.onCreateEvent = eventHandler
	return c
}

// OnDeleteEvent DeleteEvent handler.
func (c *EventHandlers) OnDeleteEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.DeleteEvent)) *EventHandlers {
	c.onDeleteEvent = eventHandler
	return c
}

// OnDeployKeyEvent DeployKeyEvent handler.
func (c *EventHandlers) OnDeployKeyEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.DeployKeyEvent)) *EventHandlers {
	c.onDeployKeyEvent = eventHandler
	return c
}

// OnDeploymentEvent DeploymentEvent handler.
func (c *EventHandlers) OnDeploymentEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.DeploymentEvent)) *EventHandlers {
	c.onDeploymentEvent = eventHandler
	return c
}

// OnDeploymentStatusEvent DeploymentStatusEvent handler.
func (c *EventHandlers) OnDeploymentStatusEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.DeploymentStatusEvent)) *EventHandlers {
	c.onDeploymentStatusEvent = eventHandler
	return c
}

// OnDiscussionEvent DiscussionEvent handler.
func (c *EventHandlers) OnDiscussionEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.DiscussionEvent)) *EventHandlers {
	c.onDiscussionEvent = eventHandler
	return c
}

// OnForkEvent ForkEvent handler.
func (c *EventHandlers) OnForkEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.ForkEvent)) *EventHandlers {
	c.onForkEvent = eventHandler
	return c
}

// OnGitHubAppAuthorizationEvent GitHubAppAuthorizationEvent handler.
func (c *EventHandlers) OnGitHubAppAuthorizationEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.GitHubAppAuthorizationEvent)) *EventHandlers {
	c.onGitHubAppAuthorizationEvent = eventHandler
	return c
}

// OnGollumEvent GollumEvent handler.
func (c *EventHandlers) OnGollumEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.GollumEvent)) *EventHandlers {
	c.onGollumEvent = eventHandler
	return c
}

// OnInstallationEvent InstallationEvent handler.
func (c *EventHandlers) OnInstallationEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.InstallationEvent)) *EventHandlers {
	c.onInstallationEvent = eventHandler
	return c
}

// OnInstallationRepositoriesEvent InstallationRepositoriesEvent handler.
func (c *EventHandlers) OnInstallationRepositoriesEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.InstallationRepositoriesEvent)) *EventHandlers {
	c.onInstallationRepositoriesEvent = eventHandler
	return c
}

// OnIssueCommentEvent IssueCommentEvent handler.
func (c *EventHandlers) OnIssueCommentEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.IssueCommentEvent)) *EventHandlers {
	c.onIssueCommentEvent = eventHandler
	return c
}

// OnIssuesEvent IssuesEvent handler.
func (c *EventHandlers) OnIssuesEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.IssuesEvent)) *EventHandlers {
	c.onIssuesEvent = eventHandler
	return c
}

// OnLabelEvent LabelEvent handler.
func (c *EventHandlers) OnLabelEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.LabelEvent)) *EventHandlers {
	c.onLabelEvent = eventHandler
	return c
}

// OnMarketplacePurchaseEvent MarketplacePurchaseEvent handler.
func (c *EventHandlers) OnMarketplacePurchaseEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.MarketplacePurchaseEvent)) *EventHandlers {
	c.onMarketplacePurchaseEvent = eventHandler
	return c
}

// OnMemberEvent MemberEvent handler.
func (c *EventHandlers) OnMemberEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.MemberEvent)) *EventHandlers {
	c.onMemberEvent = eventHandler
	return c
}

// OnMembershipEvent MembershipEvent handler.
func (c *EventHandlers) OnMembershipEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.MembershipEvent)) *EventHandlers {
	c.onMembershipEvent = eventHandler
	return c
}

// OnMetaEvent MetaEvent handler.
func (c *EventHandlers) OnMetaEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.MetaEvent)) *EventHandlers {
	c.onMetaEvent = eventHandler
	return c
}

// OnMilestoneEvent MilestoneEvent handler.
func (c *EventHandlers) OnMilestoneEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.MilestoneEvent)) *EventHandlers {
	c.onMilestoneEvent = eventHandler
	return c
}

// OnOrganizationEvent OrganizationEvent handler.
func (c *EventHandlers) OnOrganizationEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.OrganizationEvent)) *EventHandlers {
	c.onOrganizationEvent = eventHandler
	return c
}

// OnOrgBlockEvent OrgBlockEvent handler.
func (c *EventHandlers) OnOrgBlockEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.OrgBlockEvent)) *EventHandlers {
	c.onOrgBlockEvent = eventHandler
	return c
}

// OnPackageEvent PackageEvent handler.
func (c *EventHandlers) OnPackageEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PackageEvent)) *EventHandlers {
	c.onPackageEvent = eventHandler
	return c
}

// OnPageBuildEvent PageBuildEvent handler.
func (c *EventHandlers) OnPageBuildEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PageBuildEvent)) *EventHandlers {
	c.onPageBuildEvent = eventHandler
	return c
}

// OnPingEvent PingEvent handler.
func (c *EventHandlers) OnPingEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PingEvent)) *EventHandlers {
	c.onPingEvent = eventHandler
	return c
}

// OnProjectEvent ProjectEvent handler.
func (c *EventHandlers) OnProjectEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.ProjectEvent)) *EventHandlers {
	c.onProjectEvent = eventHandler
	return c
}

// OnProjectCardEvent ProjectCardEvent handler.
func (c *EventHandlers) OnProjectCardEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.ProjectCardEvent)) *EventHandlers {
	c.onProjectCardEvent = eventHandler
	return c
}

// OnProjectColumnEvent ProjectColumnEvent handler.
func (c *EventHandlers) OnProjectColumnEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.ProjectColumnEvent)) *EventHandlers {
	c.onProjectColumnEvent = eventHandler
	return c
}

// OnPublicEvent PublicEvent handler.
func (c *EventHandlers) OnPublicEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PublicEvent)) *EventHandlers {
	c.onPublicEvent = eventHandler
	return c
}

// OnPullRequestEvent PullRequestEvent handler.
func (c *EventHandlers) OnPullRequestEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PullRequestEvent)) *EventHandlers {
	c.onPullRequestEvent = eventHandler
	return c
}

// OnPullRequestReviewEvent PullRequestReviewEvent handler.
func (c *EventHandlers) OnPullRequestReviewEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PullRequestReviewEvent)) *EventHandlers {
	c.onPullRequestReviewEvent = eventHandler
	return c
}

// OnPullRequestReviewCommentEvent PullRequestReviewCommentEvent handler.
func (c *EventHandlers) OnPullRequestReviewCommentEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PullRequestReviewCommentEvent)) *EventHandlers {
	c.onPullRequestReviewCommentEvent = eventHandler
	return c
}

// OnPullRequestReviewThreadEvent PullRequestReviewThreadEvent handler.
func (c *EventHandlers) OnPullRequestReviewThreadEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PullRequestReviewThreadEvent)) *EventHandlers {
	c.onPullRequestReviewThreadEvent = eventHandler
	return c
}

// OnPullRequestTargetEvent PullRequestTargetEvent handler.
func (c *EventHandlers) OnPullRequestTargetEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PullRequestTargetEvent)) *EventHandlers {
	c.onPullRequestTargetEvent = eventHandler
	return c
}

// OnPushEvent PushEvent handler.
func (c *EventHandlers) OnPushEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.PushEvent)) *EventHandlers {
	c.onPushEvent = eventHandler
	return c
}

// OnRepositoryEvent RepositoryEvent handler.
func (c *EventHandlers) OnRepositoryEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.RepositoryEvent)) *EventHandlers {
	c.onRepositoryEvent = eventHandler
	return c
}

// OnRepositoryDispatchEvent RepositoryDispatchEvent handler.
func (c *EventHandlers) OnRepositoryDispatchEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.RepositoryDispatchEvent)) *EventHandlers {
	c.onRepositoryDispatchEvent = eventHandler
	return c
}

// OnRepositoryVulnerabilityAlertEvent RepositoryVulnerabilityAlertEvent handler.
func (c *EventHandlers) OnRepositoryVulnerabilityAlertEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.RepositoryVulnerabilityAlertEvent)) *EventHandlers {
	c.onRepositoryVulnerabilityAlertEvent = eventHandler
	return c
}

// OnReleaseEvent ReleaseEvent handler.
func (c *EventHandlers) OnReleaseEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.ReleaseEvent)) *EventHandlers {
	c.onReleaseEvent = eventHandler
	return c
}

// OnSecretScanningAlertEvent SecretScanningAlertEvent handler.
func (c *EventHandlers) OnSecretScanningAlertEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.SecretScanningAlertEvent)) *EventHandlers {
	c.onSecretScanningAlertEvent = eventHandler
	return c
}

// OnStarEvent StarEvent handler.
func (c *EventHandlers) OnStarEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.StarEvent)) *EventHandlers {
	c.onStarEvent = eventHandler
	return c
}

// OnStatusEvent StatusEvent handler.
func (c *EventHandlers) OnStatusEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.StatusEvent)) *EventHandlers {
	c.onStatusEvent = eventHandler
	return c
}

// OnTeamEvent TeamEvent handler.
func (c *EventHandlers) OnTeamEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.TeamEvent)) *EventHandlers {
	c.onTeamEvent = eventHandler
	return c
}

// OnTeamAddEvent TeamAddEvent handler.
func (c *EventHandlers) OnTeamAddEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.TeamAddEvent)) *EventHandlers {
	c.onTeamAddEvent = eventHandler
	return c
}

// OnUserEvent UserEvent handler.
func (c *EventHandlers) OnUserEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.UserEvent)) *EventHandlers {
	c.onUserEvent = eventHandler
	return c
}

// OnWatchEvent WatchEvent handler.
func (c *EventHandlers) OnWatchEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.WatchEvent)) *EventHandlers {
	c.onWatchEvent = eventHandler
	return c
}

// OnWorkflowDispatchEvent WorkflowDispatchEvent handler.
func (c *EventHandlers) OnWorkflowDispatchEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.WorkflowDispatchEvent)) *EventHandlers {
	c.onWorkflowDispatchEvent = eventHandler
	return c
}

// OnWorkflowJobEvent WorkflowJobEvent handler.
func (c *EventHandlers) OnWorkflowJobEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.WorkflowJobEvent)) *EventHandlers {
	c.onWorkflowJobEvent = eventHandler
	return c
}

// OnWorkflowRunEvent WorkflowRunEvent handler.
func (c *EventHandlers) OnWorkflowRunEvent(eventHandler func(uri *url.URL, deliveryID string, event *github.WorkflowRunEvent)) *EventHandlers {
	c.onWorkflowRunEvent = eventHandler
	return c
}
