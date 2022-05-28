package eventtype

// https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
const (
	BranchProtectionRule         = "branch_protection_rule"
	CheckRun                     = "check_run"
	CheckSuite                   = "check_suite"
	CodeScanningAlert            = "code_scanning_alert"
	CommitComment                = "commit_comment"
	ContentReference             = "content_reference"
	Create                       = "create"
	Delete                       = "delete"
	DeployKey                    = "deploy_key"
	Deployment                   = "deployment"
	DeploymentStatus             = "deployment_status"
	Discussion                   = "discussion"
	DiscussionComment            = "discussion_comment"
	Fork                         = "fork"
	GithubAppAuthorization       = "github_app_authorization"
	Gollum                       = "gollum"
	Installation                 = "installation"
	InstallationRepositories     = "installation_repositories"
	IssueComment                 = "issue_comment"
	Issues                       = "issues"
	Label                        = "label"
	MarketplacePurchase          = "marketplace_purchase"
	Member                       = "member"
	Membership                   = "membership"
	Meta                         = "meta"
	Milestone                    = "milestone"
	Organization                 = "organization"
	OrgBlock                     = "org_block"
	Package                      = "package"
	PageBuild                    = "page_build"
	Ping                         = "ping"
	Project                      = "project"
	ProjectCard                  = "project_card"
	ProjectColumn                = "project_column"
	Public                       = "public"
	PullRequest                  = "pull_request"
	PullRequestReview            = "pull_request_review"
	PullRequestReviewComment     = "pull_request_review_comment"
	PullRequestReviewThread      = "pull_request_review_thread"
	Push                         = "push"
	Release                      = "release"
	Repository                   = "repository"
	RepositoryDispatch           = "repository_dispatch"
	RepositoryImport             = "repository_import"
	RepositoryVulnerabilityAlert = "repository_vulnerability_alert"
	SecurityAdvisory             = "security_advisory"
	Sponsorship                  = "sponsorship"
	Star                         = "star"
	Status                       = "status"
	Team                         = "team"
	TeamAdd                      = "team_add"
	Watch                        = "watch"
	WorkflowDispatch             = "workflow_dispatch"
	WorkflowJob                  = "workflow_job"
	WorkflowRun                  = "workflow_run"
)
