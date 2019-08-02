package ghwebhook

import (
	"regexp"

	"github.com/ldez/ghwebhook/v2/eventtype"
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
		eventtype.Ping,
		eventtype.CheckRun,
		eventtype.CheckSuite,
		eventtype.CommitComment,
		eventtype.ContentReference,
		eventtype.Create,
		eventtype.Delete,
		eventtype.Deployment,
		eventtype.DeploymentStatus,
		eventtype.Fork,
		eventtype.GitHubAppAuthorization,
		eventtype.Gollum,
		eventtype.Installation,
		eventtype.InstallationRepositories,
		eventtype.IssueComment,
		eventtype.Issues,
		eventtype.Label,
		eventtype.MarketplacePurchase,
		eventtype.Member,
		eventtype.Membership,
		eventtype.Milestone,
		eventtype.Organization,
		eventtype.OrgBlock,
		eventtype.PageBuild,
		eventtype.ProjectCard,
		eventtype.ProjectColumn,
		eventtype.Project,
		eventtype.Public,
		eventtype.PullRequest,
		eventtype.PullRequestReview,
		eventtype.PullRequestReviewComment,
		eventtype.Push,
		eventtype.Release,
		eventtype.Repository,
		eventtype.RepositoryImport,
		eventtype.RepositoryVulnerabilityAlert,
		eventtype.SecurityAdvisory,
		eventtype.Status,
		eventtype.Team,
		eventtype.TeamAdd,
		eventtype.Watch,
	}
}
