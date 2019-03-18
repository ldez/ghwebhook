package eventtype

// https://developer.github.com/v3/activity/events/types/
const (
	Ping                         = "ping"
	CheckRun                     = "check_run"
	CheckSuite                   = "check_suite"
	CommitComment                = "commit_comment"
	ContentReference             = "content_reference" // TODO not yet supported by go-github
	Create                       = "create"
	Delete                       = "delete"
	Deployment                   = "deployment"
	DeploymentStatus             = "deployment_status"
	Download                     = "download" // Deprecated
	Follow                       = "follow"   // Deprecated
	Fork                         = "fork"
	ForkApply                    = "fork_apply" // Deprecated
	GitHubAppAuthorization       = "github_app_authorization"
	Gist                         = "gist" // Deprecated
	Gollum                       = "gollum"
	Installation                 = "installation"
	InstallationRepositories     = "installation_repositories"
	IssueComment                 = "issue_comment"
	Issues                       = "issues"
	Label                        = "label"
	MarketplacePurchase          = "marketplace_purchase"
	Member                       = "member"
	Membership                   = "membership"
	Milestone                    = "milestone"
	Organization                 = "organization"
	OrgBlock                     = "org_block"
	PageBuild                    = "page_build"
	ProjectCard                  = "project_card"
	ProjectColumn                = "project_column"
	Project                      = "project"
	Public                       = "public"
	PullRequest                  = "pull_request"
	PullRequestReview            = "pull_request_review"
	PullRequestReviewComment     = "pull_request_review_comment"
	Push                         = "push"
	Release                      = "release"
	Repository                   = "repository"
	RepositoryImport             = "repository_import" // TODO not yet supported by go-github
	RepositoryVulnerabilityAlert = "repository_vulnerability_alert"
	SecurityAdvisory             = "security_advisory" // TODO not yet supported by go-github
	Status                       = "status"
	Team                         = "team"
	TeamAdd                      = "team_add"
	Watch                        = "watch"
)
