package ghwebhook

import (
	"regexp"

	"github.com/ldez/ghwebhook/v4/eventtype"
)

type serverOption func(*WebHook)

// WithPort define the server port.
func WithPort(port int) func(*WebHook) {
	return func(server *WebHook) {
		server.port = port
	}
}

// WithPath define the HTTP handler path.
func WithPath(path string) func(*WebHook) {
	return func(server *WebHook) {
		server.path = regexp.MustCompile(path)
	}
}

// WithSecret define the GitHub secret.
func WithSecret(secret string) func(*WebHook) {
	return func(server *WebHook) {
		server.secret = secret
	}
}

// WithEventTypes define accepted event types.
func WithEventTypes(eventTypes ...string) func(*WebHook) {
	return func(server *WebHook) {
		server.eventTypes = eventTypes
	}
}

// Debug activate debug mode.
func Debug(server *WebHook) {
	server.debug = true
}

// WithAllEventTypes accept all possible event types.
func WithAllEventTypes(server *WebHook) {
	server.eventTypes = []string{
		eventtype.BranchProtectionConfiguration,
		eventtype.BranchProtectionRule,
		eventtype.CheckRun,
		eventtype.CheckSuite,
		eventtype.CodeScanningAlert,
		eventtype.CommitComment,
		eventtype.ContentReference,
		eventtype.Create,
		eventtype.CustomProperty,
		eventtype.CustomPropertyValues,
		eventtype.Delete,
		eventtype.DependabotAlert,
		eventtype.DeployKey,
		eventtype.Deployment,
		eventtype.DeploymentReview,
		eventtype.DeploymentStatus,
		eventtype.DeploymentProtectionRule,
		eventtype.Discussion,
		eventtype.DiscussionComment,
		eventtype.Fork,
		eventtype.GitHubAppAuthorization,
		eventtype.Gollum,
		eventtype.Installation,
		eventtype.InstallationRepositories,
		eventtype.InstallationTarget,
		eventtype.IssueComment,
		eventtype.Issues,
		eventtype.Label,
		eventtype.MarketplacePurchase,
		eventtype.Member,
		eventtype.Membership,
		eventtype.MergeGroup,
		eventtype.Meta,
		eventtype.Milestone,
		eventtype.Organization,
		eventtype.OrgBlock,
		eventtype.Package,
		eventtype.PageBuild,
		eventtype.PersonalAccessTokenRequest,
		eventtype.Ping,
		eventtype.ProjectV2,
		eventtype.ProjectV2Item,
		eventtype.Public,
		eventtype.PullRequest,
		eventtype.PullRequestReview,
		eventtype.PullRequestReviewComment,
		eventtype.PullRequestReviewThread,
		eventtype.PullRequestTarget,
		eventtype.Push,
		eventtype.Repository,
		eventtype.RepositoryDispatch,
		eventtype.RepositoryImport,
		eventtype.RepositoryRuleset,
		eventtype.RepositoryVulnerabilityAlert,
		eventtype.Release,
		eventtype.SecretScanningAlert,
		eventtype.SecretScanningAlertLocation,
		eventtype.SecurityAdvisory,
		eventtype.SecurityAndAnalysis,
		eventtype.Sponsorship,
		eventtype.Star,
		eventtype.Status,
		eventtype.Team,
		eventtype.TeamAdd,
		eventtype.User,
		eventtype.Watch,
		eventtype.WorkflowDispatch,
		eventtype.WorkflowJob,
		eventtype.WorkflowRun,
	}
}
