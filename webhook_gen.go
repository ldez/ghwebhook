package ghwebhook

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"
	"net/url"

	"github.com/google/go-github/v45/github"
)

func (s *WebHook) handleEvents(uri *url.URL, deliveryID string, rawEvent interface{}) error {
	switch event := rawEvent.(type) {

	case *github.BranchProtectionRuleEvent:
		if s.eventHandlers.onBranchProtectionRuleEvent != nil {
			s.eventHandlers.onBranchProtectionRuleEvent(uri, deliveryID, event)
		}
	case *github.CheckRunEvent:
		if s.eventHandlers.onCheckRunEvent != nil {
			s.eventHandlers.onCheckRunEvent(uri, deliveryID, event)
		}
	case *github.CheckSuiteEvent:
		if s.eventHandlers.onCheckSuiteEvent != nil {
			s.eventHandlers.onCheckSuiteEvent(uri, deliveryID, event)
		}
	case *github.CommitCommentEvent:
		if s.eventHandlers.onCommitCommentEvent != nil {
			s.eventHandlers.onCommitCommentEvent(uri, deliveryID, event)
		}
	case *github.ContentReferenceEvent:
		if s.eventHandlers.onContentReferenceEvent != nil {
			s.eventHandlers.onContentReferenceEvent(uri, deliveryID, event)
		}
	case *github.CreateEvent:
		if s.eventHandlers.onCreateEvent != nil {
			s.eventHandlers.onCreateEvent(uri, deliveryID, event)
		}
	case *github.DeleteEvent:
		if s.eventHandlers.onDeleteEvent != nil {
			s.eventHandlers.onDeleteEvent(uri, deliveryID, event)
		}
	case *github.DeployKeyEvent:
		if s.eventHandlers.onDeployKeyEvent != nil {
			s.eventHandlers.onDeployKeyEvent(uri, deliveryID, event)
		}
	case *github.DeploymentEvent:
		if s.eventHandlers.onDeploymentEvent != nil {
			s.eventHandlers.onDeploymentEvent(uri, deliveryID, event)
		}
	case *github.DeploymentStatusEvent:
		if s.eventHandlers.onDeploymentStatusEvent != nil {
			s.eventHandlers.onDeploymentStatusEvent(uri, deliveryID, event)
		}
	case *github.DiscussionEvent:
		if s.eventHandlers.onDiscussionEvent != nil {
			s.eventHandlers.onDiscussionEvent(uri, deliveryID, event)
		}
	case *github.ForkEvent:
		if s.eventHandlers.onForkEvent != nil {
			s.eventHandlers.onForkEvent(uri, deliveryID, event)
		}
	case *github.GitHubAppAuthorizationEvent:
		if s.eventHandlers.onGitHubAppAuthorizationEvent != nil {
			s.eventHandlers.onGitHubAppAuthorizationEvent(uri, deliveryID, event)
		}
	case *github.GollumEvent:
		if s.eventHandlers.onGollumEvent != nil {
			s.eventHandlers.onGollumEvent(uri, deliveryID, event)
		}
	case *github.InstallationEvent:
		if s.eventHandlers.onInstallationEvent != nil {
			s.eventHandlers.onInstallationEvent(uri, deliveryID, event)
		}
	case *github.InstallationRepositoriesEvent:
		if s.eventHandlers.onInstallationRepositoriesEvent != nil {
			s.eventHandlers.onInstallationRepositoriesEvent(uri, deliveryID, event)
		}
	case *github.IssueCommentEvent:
		if s.eventHandlers.onIssueCommentEvent != nil {
			s.eventHandlers.onIssueCommentEvent(uri, deliveryID, event)
		}
	case *github.IssuesEvent:
		if s.eventHandlers.onIssuesEvent != nil {
			s.eventHandlers.onIssuesEvent(uri, deliveryID, event)
		}
	case *github.LabelEvent:
		if s.eventHandlers.onLabelEvent != nil {
			s.eventHandlers.onLabelEvent(uri, deliveryID, event)
		}
	case *github.MarketplacePurchaseEvent:
		if s.eventHandlers.onMarketplacePurchaseEvent != nil {
			s.eventHandlers.onMarketplacePurchaseEvent(uri, deliveryID, event)
		}
	case *github.MemberEvent:
		if s.eventHandlers.onMemberEvent != nil {
			s.eventHandlers.onMemberEvent(uri, deliveryID, event)
		}
	case *github.MembershipEvent:
		if s.eventHandlers.onMembershipEvent != nil {
			s.eventHandlers.onMembershipEvent(uri, deliveryID, event)
		}
	case *github.MetaEvent:
		if s.eventHandlers.onMetaEvent != nil {
			s.eventHandlers.onMetaEvent(uri, deliveryID, event)
		}
	case *github.MilestoneEvent:
		if s.eventHandlers.onMilestoneEvent != nil {
			s.eventHandlers.onMilestoneEvent(uri, deliveryID, event)
		}
	case *github.OrganizationEvent:
		if s.eventHandlers.onOrganizationEvent != nil {
			s.eventHandlers.onOrganizationEvent(uri, deliveryID, event)
		}
	case *github.OrgBlockEvent:
		if s.eventHandlers.onOrgBlockEvent != nil {
			s.eventHandlers.onOrgBlockEvent(uri, deliveryID, event)
		}
	case *github.PackageEvent:
		if s.eventHandlers.onPackageEvent != nil {
			s.eventHandlers.onPackageEvent(uri, deliveryID, event)
		}
	case *github.PageBuildEvent:
		if s.eventHandlers.onPageBuildEvent != nil {
			s.eventHandlers.onPageBuildEvent(uri, deliveryID, event)
		}
	case *github.PingEvent:
		if s.eventHandlers.onPingEvent != nil {
			s.eventHandlers.onPingEvent(uri, deliveryID, event)
		}
	case *github.ProjectEvent:
		if s.eventHandlers.onProjectEvent != nil {
			s.eventHandlers.onProjectEvent(uri, deliveryID, event)
		}
	case *github.ProjectCardEvent:
		if s.eventHandlers.onProjectCardEvent != nil {
			s.eventHandlers.onProjectCardEvent(uri, deliveryID, event)
		}
	case *github.ProjectColumnEvent:
		if s.eventHandlers.onProjectColumnEvent != nil {
			s.eventHandlers.onProjectColumnEvent(uri, deliveryID, event)
		}
	case *github.PublicEvent:
		if s.eventHandlers.onPublicEvent != nil {
			s.eventHandlers.onPublicEvent(uri, deliveryID, event)
		}
	case *github.PullRequestEvent:
		if s.eventHandlers.onPullRequestEvent != nil {
			s.eventHandlers.onPullRequestEvent(uri, deliveryID, event)
		}
	case *github.PullRequestReviewEvent:
		if s.eventHandlers.onPullRequestReviewEvent != nil {
			s.eventHandlers.onPullRequestReviewEvent(uri, deliveryID, event)
		}
	case *github.PullRequestReviewCommentEvent:
		if s.eventHandlers.onPullRequestReviewCommentEvent != nil {
			s.eventHandlers.onPullRequestReviewCommentEvent(uri, deliveryID, event)
		}
	case *github.PullRequestReviewThreadEvent:
		if s.eventHandlers.onPullRequestReviewThreadEvent != nil {
			s.eventHandlers.onPullRequestReviewThreadEvent(uri, deliveryID, event)
		}
	case *github.PullRequestTargetEvent:
		if s.eventHandlers.onPullRequestTargetEvent != nil {
			s.eventHandlers.onPullRequestTargetEvent(uri, deliveryID, event)
		}
	case *github.PushEvent:
		if s.eventHandlers.onPushEvent != nil {
			s.eventHandlers.onPushEvent(uri, deliveryID, event)
		}
	case *github.RepositoryEvent:
		if s.eventHandlers.onRepositoryEvent != nil {
			s.eventHandlers.onRepositoryEvent(uri, deliveryID, event)
		}
	case *github.RepositoryDispatchEvent:
		if s.eventHandlers.onRepositoryDispatchEvent != nil {
			s.eventHandlers.onRepositoryDispatchEvent(uri, deliveryID, event)
		}
	case *github.RepositoryImportEvent:
		if s.eventHandlers.onRepositoryImportEvent != nil {
			s.eventHandlers.onRepositoryImportEvent(uri, deliveryID, event)
		}
	case *github.RepositoryVulnerabilityAlertEvent:
		if s.eventHandlers.onRepositoryVulnerabilityAlertEvent != nil {
			s.eventHandlers.onRepositoryVulnerabilityAlertEvent(uri, deliveryID, event)
		}
	case *github.ReleaseEvent:
		if s.eventHandlers.onReleaseEvent != nil {
			s.eventHandlers.onReleaseEvent(uri, deliveryID, event)
		}
	case *github.SecretScanningAlertEvent:
		if s.eventHandlers.onSecretScanningAlertEvent != nil {
			s.eventHandlers.onSecretScanningAlertEvent(uri, deliveryID, event)
		}
	case *github.StarEvent:
		if s.eventHandlers.onStarEvent != nil {
			s.eventHandlers.onStarEvent(uri, deliveryID, event)
		}
	case *github.StatusEvent:
		if s.eventHandlers.onStatusEvent != nil {
			s.eventHandlers.onStatusEvent(uri, deliveryID, event)
		}
	case *github.TeamEvent:
		if s.eventHandlers.onTeamEvent != nil {
			s.eventHandlers.onTeamEvent(uri, deliveryID, event)
		}
	case *github.TeamAddEvent:
		if s.eventHandlers.onTeamAddEvent != nil {
			s.eventHandlers.onTeamAddEvent(uri, deliveryID, event)
		}
	case *github.UserEvent:
		if s.eventHandlers.onUserEvent != nil {
			s.eventHandlers.onUserEvent(uri, deliveryID, event)
		}
	case *github.WatchEvent:
		if s.eventHandlers.onWatchEvent != nil {
			s.eventHandlers.onWatchEvent(uri, deliveryID, event)
		}
	case *github.WorkflowDispatchEvent:
		if s.eventHandlers.onWorkflowDispatchEvent != nil {
			s.eventHandlers.onWorkflowDispatchEvent(uri, deliveryID, event)
		}
	case *github.WorkflowJobEvent:
		if s.eventHandlers.onWorkflowJobEvent != nil {
			s.eventHandlers.onWorkflowJobEvent(uri, deliveryID, event)
		}
	case *github.WorkflowRunEvent:
		if s.eventHandlers.onWorkflowRunEvent != nil {
			s.eventHandlers.onWorkflowRunEvent(uri, deliveryID, event)
		}
	default:
		return fmt.Errorf("unknown event type: %T", rawEvent)
	}

	return nil
}
