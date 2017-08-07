package ghwebhook

import "github.com/ldez/ghwebhook/eventtype"

type serverOption func(*webHook)

func WithPort(port int) func(*webHook) {
	return func(server *webHook) {
		server.port = port
	}
}

func WithPath(path string) func(*webHook) {
	return func(server *webHook) {
		server.path = path
	}
}

func WithSecret(secret string) func(*webHook) {
	return func(server *webHook) {
		server.secret = secret
	}
}

func WithEventTypes(eventTypes ...string) func(*webHook) {
	return func(server *webHook) {
		server.eventTypes = eventTypes
	}
}

func Debug(server *webHook) {
	server.debug = true
}

func WithAllEventTypes(server *webHook) {
	server.eventTypes = []string{
		eventtype.CommitComment,
		eventtype.Create,
		eventtype.Delete,
		eventtype.Deployment,
		eventtype.DeploymentStatus,
		eventtype.Download,
		eventtype.Follow,
		eventtype.Fork,
		eventtype.Gist,
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
		eventtype.Status,
		eventtype.Team,
		eventtype.TeamAdd,
		eventtype.Watch,
	}
}
