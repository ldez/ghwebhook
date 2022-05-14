package ghwebhook

import (
	"regexp"

	"github.com/ldez/ghwebhook/v3/eventtype"
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
		eventtype.BranchProtectionRule,
		eventtype.CheckRun,
		eventtype.CheckSuite,
		eventtype.CodeScanningAlert,
		eventtype.CommitComment,
		eventtype.ContentReference,
		eventtype.Create,
		eventtype.Delete,
		eventtype.DeployKey,
		eventtype.Deployment,
		eventtype.DeploymentStatus,
		eventtype.Discussion,
		eventtype.DiscussionComment,
		eventtype.Fork,
		eventtype.GithubAppAuthorization,
		eventtype.Gollum,
		eventtype.Installation,
		eventtype.InstallationRepositories,
		eventtype.IssueComment,
		eventtype.Issues,
		eventtype.Label,
		eventtype.MarketplacePurchase,
		eventtype.Member,
		eventtype.Membership,
		eventtype.Meta,
		eventtype.Milestone,
		eventtype.Organization,
		eventtype.OrgBlock,
		eventtype.Package,
		eventtype.PageBuild,
		eventtype.Ping,
		eventtype.Project,
		eventtype.ProjectCard,
		eventtype.ProjectColumn,
		eventtype.Public,
		eventtype.PullRequest,
		eventtype.PullRequestReview,
		eventtype.PullRequestReviewComment,
		eventtype.PullRequestReviewThread,
		eventtype.Push,
		eventtype.Release,
		eventtype.Repository,
		eventtype.RepositoryDispatch,
		eventtype.RepositoryImport,
		eventtype.RepositoryVulnerabilityAlert,
		eventtype.SecurityAdvisory,
		eventtype.Sponsorship,
		eventtype.Star,
		eventtype.Status,
		eventtype.Team,
		eventtype.TeamAdd,
		eventtype.Watch,
		eventtype.WorkflowDispatch,
		eventtype.WorkflowJob,
		eventtype.WorkflowRun,
	}
}
